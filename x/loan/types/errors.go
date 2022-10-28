package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/loan module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrWrongLoanState       = sdkerrors.Register(ModuleName, 1502, "invalid loan state")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrDeadline             = sdkerrors.Register(ModuleName, 1503, "deadline")
)
