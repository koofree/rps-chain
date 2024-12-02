package keeper

import (
	"context"
	"fmt"
	"strings"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/0xlb/rpschain/x/rps/rules"
	"github.com/0xlb/rpschain/x/rps/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) CreateGame(ctx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	params, err := ms.k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight() + int64(params.Ttl)

	newGame := types.Game{
		GameNumber:       ms.k.NextGameNumber(ctx),
		PlayerA:          msg.Creator,
		PlayerB:          msg.Opponent,
		Rounds:           msg.Rounds,
		Status:           rules.StatusWaiting,
		Score:            []uint64{0, 0},
		ExpirationHeight: uint64(blockHeight),
	}

	if err := newGame.Validate(); err != nil {
		return nil, err
	}

	if err := ms.k.Games.Set(ctx, newGame.GameNumber, newGame); err != nil {
		return nil, err
	}

	// add the game to the ActiveGamesQueue
	err = ms.k.ActiveGamesQueue.Set(ctx, collections.Join(newGame.ExpirationHeight, newGame.GameNumber))
	if err != nil {
		return nil, err
	}

	sdkCtx.EventManager().EmitTypedEvent(&types.EventCreateGame{
		GameNumber: newGame.GameNumber,
		PlayerA:    newGame.PlayerA,
		PlayerB:    newGame.PlayerB,
	})

	return &types.MsgCreateGameResponse{}, nil
}

func (ms msgServer) MakeMove(ctx context.Context, msg *types.MsgMakeMove) (*types.MsgMakeMoveResponse, error) {

	// Game exists
	game, err := ms.k.Games.Get(ctx, msg.GameIndex)
	if err != nil {
		return nil, err
	}

	// Game Status is InProgress on Waiting
	if game.Ended() {
		return nil, types.ErrGameEnded
	}

	if err := game.AddPlayerMove(msg.Player, msg.Move); err != nil {
		return nil, err
	}

	// game status is InProgress
	game.Status = rules.StatusInProgress

	if err := game.Validate(); err != nil {
		return nil, err
	}

	if err := ms.k.Games.Set(ctx, msg.GameIndex, game); err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitTypedEvent(&types.EventMakeMove{
		GameNumber: msg.GameIndex,
		Player:     msg.Player,
		Move:       msg.Move,
	})

	return &types.MsgMakeMoveResponse{}, nil
}

func (ms msgServer) RevealMove(ctx context.Context, msg *types.MsgRevealMove) (*types.MsgRevealMoveResponse, error) {
	// Game exists
	game, err := ms.k.Games.Get(ctx, msg.GameIndex)
	if err != nil {
		return nil, err
	}

	// Game Status is InProgress or Waiting
	if game.Ended() {
		return nil, types.ErrGameEnded
	}

	if err := game.RevealPlayerMove(msg.Player, msg.RevealedMove, msg.Salt); err != nil {
		return nil, err
	}

	if game.IsRoundRevealed() {
		playerAResult := rules.DetermineRoundWinner(
			rules.Choice(game.GetPlayerALastMove()),
			rules.Choice(game.GetPlayerBLastMove()),
		)
		// game.Score stores the playerA and playerB wins in an array
		if playerAResult == rules.Win {
			game.AddWinToPlayerA()
		}
		if playerAResult == rules.Loss {
			game.AddWinToPlayerB()
		}
	}

	game.Status = rules.GameResultByMajority(game.GetPlayerAScore(), game.GetPlayerBScore(), game.Rounds)

	if err := ms.k.Games.Set(ctx, msg.GameIndex, game); err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitTypedEvent(&types.EventRevealMove{
		GameNumber:   msg.GameIndex,
		Player:       msg.Player,
		RevealedMove: msg.RevealedMove,
	})

	if game.Ended() {
		// remove the game from the ActiveGamesQueue
		if err := ms.k.ActiveGamesQueue.Remove(ctx, collections.Join(game.ExpirationHeight, game.GameNumber)); err != nil {
			return nil, err
		}
		// game has ended. Emit the game ended event
		sdkCtx.EventManager().EmitTypedEvent(&types.EventEndGame{
			GameNumber: msg.GameIndex,
			Status:     game.Status,
		})
	}

	return &types.MsgRevealMoveResponse{}, nil
}

func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if _, err := ms.k.addressCodec.StringToBytes(msg.Authority); err != nil {
		return nil, fmt.Errorf("invalid authority: %w", err)
	}

	if authority := ms.k.GetAuthority(); !strings.EqualFold(authority, msg.Authority) {
		return nil, fmt.Errorf("unauthorized. authority does not match the module's authority: got %s, want %s", authority, msg.Authority)
	}

	if err := msg.Params.Validate(); err != nil {
		return nil, err
	}

	if err := ms.k.Params.Set(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
