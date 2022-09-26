package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/claim module sentinel errors
var (
	ErrAlreadyClaimed    = sdkerrors.New(ModuleName, 2, "already claimed condition")
	ErrTerminatedAirdrop = sdkerrors.New(ModuleName, 3, "terminated airdrop event")
	ErrConditionRequired = sdkerrors.New(ModuleName, 4, "condition must be executed first")
)
