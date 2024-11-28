package types

import "cosmossdk.io/errors"

var (
	ErrInvalidAddress     = errors.Register(ModuleName, 1, "Invalid address")
	ErrRoundsOutOfBounds  = errors.Register(ModuleName, 2, "game rounds out of bounds")
	ErrInvalidMovesNumber = errors.Register(ModuleName, 3, "invalid moves count")
	ErrInvalidGameNumber  = errors.Register(ModuleName, 4, "invalid game number. Should be greater than 0")
	ErrInvalidGameStatus  = errors.Register(ModuleName, 5, "invalid game status")
	ErrInvalidScore       = errors.Register(ModuleName, 6, "invalid score")
	ErrInvalidOpponent    = errors.Register(ModuleName, 7, "invalid opponent address")
	ErrDuplicatedIndex    = errors.Register(ModuleName, 8, "duplicated index")
)
