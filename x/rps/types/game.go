package types

import (
	"bytes"

	"cosmossdk.io/errors"
	"github.com/0xlb/rpschain/x/rps/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const MaxRound = 3

func DefaultGames() (games []Game) {
	return
}

func (g Game) getPlayerAAddress() (sdk.AccAddress, error) {
	return getPlayerAddress(g.PlayerA)
}
func (g Game) getPlayerBAddress() (sdk.AccAddress, error) {
	return getPlayerAddress(g.PlayerB)
}

func (g Game) GetPlayerAScore() uint64 {
	return g.Score[0]
}

func (g Game) GetPlayerBScore() uint64 {
	return g.Score[1]
}

func (g Game) GetPlayerALastMove() string {
	movesCount := len(g.PlayerAMoves)
	if movesCount == 0 {
		return ""
	}
	return g.PlayerAMoves[movesCount-1]
}

func (g Game) GetPlayerBLastMove() string {
	movesCount := len(g.PlayerBMoves)
	if movesCount == 0 {
		return ""
	}
	return g.PlayerBMoves[movesCount-1]
}

// GetPlayerType returns the type of the player
func (g Game) GetPlayerType(playerAddr string) (player rules.Player, err error) {
	switch playerAddr {
	case g.PlayerA:
		player = rules.PlayerA
	case g.PlayerB:
		player = rules.PlayerB
	default:
		player = rules.InvalidPlayer
		err = ErrInvalidPlayer
	}
	return
}

// GetPlayerLastMove returns the last move of the player
func (g Game) GetPlayerLastMove(playerAddr string) (player rules.Player, move string, err error) {
	switch playerAddr {
	case g.PlayerA:
		player = rules.PlayerA
		move = g.GetPlayerALastMove()
	case g.PlayerB:
		player = rules.PlayerB
		move = g.GetPlayerBLastMove()
	default:
		player = rules.InvalidPlayer
		err = ErrInvalidPlayer
	}
	return
}

// IsRoundRevealed checks if the round is revealed
// the same amount of moves and the last move is valid
func (g Game) IsRoundRevealed() bool {
	playerAMovesCount, playerBMovesCount := len(g.PlayerAMoves), len(g.PlayerBMoves)
	if playerAMovesCount != playerBMovesCount {
		return false
	}
	// empty list
	if playerAMovesCount == 0 {
		return true
	}

	return rules.IsValidMove(g.GetPlayerALastMove()) && rules.IsValidMove(g.GetPlayerBLastMove())
}

// AddWinToPlayerA adds a win to the player A
func (g *Game) AddWinToPlayerA() {
	g.Score[0]++
}

// AddWinToPlayerB adds a win to the player B
func (g *Game) AddWinToPlayerB() {
	g.Score[1]++
}

// AddPlayerMove adds a move to the player
func (g *Game) AddPlayerMove(playerAddr, move string) error {
	// Player is in the game
	player, err := g.GetPlayerType(playerAddr)
	if err != nil {
		return err
	}

	if ok := isValidHash(move); !ok {
		return ErrInvalidCommitment
	}

	// To submit a commitment, the player must know the previous move of the opponent
	// move to revealed
	var (
		prevMove         string
		opponentPrevMove string
	)

	playerAMoves, playerBMoves := g.GetPlayerAMoves(), g.GetPlayerBMoves()

	// Get the previous move of the opponent
	switch playerAddr {
	case g.PlayerA:
		prevMove = g.GetPlayerALastMove()
		opponentPrevMove = g.GetPlayerBLastMove()
		playerAMoves = append(playerAMoves, move)
	case g.PlayerB:
		prevMove = g.GetPlayerBLastMove()
		opponentPrevMove = g.GetPlayerALastMove()
		playerBMoves = append(playerBMoves, move)
	default:
		return ErrInvalidPlayer
	}

	ok := rules.IsValidMove(prevMove)
	if prevMove != "" && !ok {
		return errors.Wrapf(ErrRevealPreviousMove, "player with address %s has to reveal the move to finish the round", playerAddr)
	}

	playerAMovesCount, playerBMovesCount := len(playerAMoves), len(playerBMoves)
	// If previous move was revealed, but opponent's previous move was not revealed,
	// then cannot submit a new commitment.
	// Unless is evening out the moves (playerAMovesCount == playerBMovesCount)
	ok = rules.IsValidMove(opponentPrevMove)
	if opponentPrevMove != "" && !ok && playerAMovesCount != playerBMovesCount {
		return errors.Wrapf(ErrRevealPreviousMove, "opponent player has to reveal the move to finish the round")
	}

	// Can make the move - depends on:
	//  - rules: game status, rounds count, other player moves
	if ok := rules.CanMakeMove(player, playerAMovesCount, playerBMovesCount); !ok {
		return ErrPlayerCantMakeMove
	}

	// all validations passed
	// update the game object with the new move
	g.PlayerAMoves = playerAMoves
	g.PlayerBMoves = playerBMoves

	return nil
}

func (game *Game) RevealPlayerMove(playerAddr, revealedMove, salt string) error {
	if ok := rules.IsValidMove(revealedMove); !ok {
		return ErrInvalidMove
	}

	// Can only reveal if both players submitted the commitment
	playerAMovesCount, playerBMovesCount := len(game.PlayerAMoves), len(game.PlayerBMoves)
	if playerAMovesCount != playerBMovesCount {
		return ErrPlayerCantRevealMove
	}

	var commitment string
	switch playerAddr {
	case game.PlayerA:
		commitment = game.GetPlayerALastMove()
		game.PlayerAMoves[playerAMovesCount-1] = revealedMove
	case game.PlayerB:
		commitment = game.GetPlayerBLastMove()
		game.PlayerBMoves[playerBMovesCount-1] = revealedMove
	default:
		return ErrInvalidPlayer
	}

	// check that this move wasn't revealed previously
	if ok := isValidHash(commitment); !ok {
		return errors.Wrapf(ErrMoveAlreadyRevealed, "move %s was already revealed", commitment)
	}

	//calculate the hash and compare with the revealed move
	if ok := isMoveRevealed(commitment, revealedMove, salt); !ok {
		return ErrWrongMoveRevealed
	}

	return nil
}

func (g Game) ValidateRounds() error {
	if g.Rounds <= MaxRound && g.Rounds > 0 {
		return nil
	}

	return ErrRoundsOutOfBounds
}

func (g Game) ValidateMovesCount() error {
	if len(g.PlayerAMoves) <= int(g.Rounds) && len(g.PlayerBMoves) <= int(g.Rounds) {
		return nil
	}
	return ErrInvalidMovesNumber
}

func (g Game) ValidateGameNumber() error {
	if g.GameNumber > 0 {
		return nil
	}
	return ErrInvalidGameNumber
}

func (g Game) ValidateStatus() error {
	if rules.IsValidStatus(g.Status) {
		return nil
	}
	return ErrInvalidGameStatus
}

func (g Game) ValidateScore() error {
	scLen := len(g.Score)
	if scLen != 2 {
		return ErrInvalidScore
	}
	if g.Score[0]+g.Score[1] > g.Rounds {
		return ErrInvalidScore
	}
	return nil
}

func (g Game) Validate() error {
	accA, err := g.getPlayerAAddress()
	if err != nil {
		return err
	}
	accB, err := g.getPlayerBAddress()
	if err != nil {
		return err
	}
	if bytes.Equal(accA, accB) {
		return ErrInvalidOpponent
	}
	if err := g.ValidateGameNumber(); err != nil {
		return err
	}
	if err := g.ValidateStatus(); err != nil {
		return err
	}
	if err := g.ValidateRounds(); err != nil {
		return err
	}
	if err := g.ValidateMovesCount(); err != nil {
		return err
	}

	return g.ValidateScore()
}

func (g Game) Ended() bool {
	return g.Status == rules.StatusPlayerAWins ||
		g.Status == rules.StatusPlayerBWins ||
		g.Status == rules.StatusDraw ||
		g.Status == rules.StatusCancelled
}

func getPlayerAddress(address string) (sdk.AccAddress, error) {
	addr, err := sdk.AccAddressFromBech32(address)

	return addr, errors.Wrapf(err, ErrInvalidAddress.Error(), address)
}
