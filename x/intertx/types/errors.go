package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/intertx module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
    ErrInvalidVersion = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrIBCAccountAlreadyExist = sdkerrors.Register(ModuleName, 1502, "interchain account already registered")
	ErrIBCAccountNotExist     = sdkerrors.Register(ModuleName, 1503, "interchain account not exist")
)
