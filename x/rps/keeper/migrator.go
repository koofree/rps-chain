package keeper

import (
	"cosmossdk.io/collections"
	v2 "github.com/0xlb/rpschain/x/rps/migrations/v2"
	"github.com/0xlb/rpschain/x/rps/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct that migrates the rps module from v1 to v2.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator for the state migration.
func NewMigrator(k Keeper) Migrator {
	return Migrator{
		keeper: k,
	}
}

// Migrate1to2 migrates the rps module from v1 to v2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	// migrate games - add the expiration height for existing games
	logger := ctx.Logger().With("upgrade", "v2")
	logger.Debug("Migrate1to2... starting")

	if err := m.keeper.Games.Walk(ctx, nil, func(id uint64, game types.Game) (stop bool, err error) {
		logger := ctx.Logger().With("upgrade", "v2")
		logger.Debug("Migrate1to2... gameNumber is %d", game.GameNumber)

		mg := v2.MigrateGame(ctx, game)
		if err := m.keeper.Games.Set(ctx, id, mg); err != nil {
			return true, err
		}
		// add these to the active game list
		if err := m.keeper.ActiveGamesQueue.Set(ctx, collections.Join(mg.ExpirationHeight, id)); err != nil {
			return true, err
		}

		return false, nil
	}); err != nil {
		return err
	}

	return m.keeper.Params.Set(ctx, types.DefaultParams())
}
