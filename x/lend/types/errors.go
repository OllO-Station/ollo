package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrWrongLoanState = sdkerrors.Register(ModuleName, 2, "wrong loan state")
	ErrDeadline       = sdkerrors.Register(ModuleName, 3, "deadline")
)
