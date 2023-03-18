package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDenom                 = sdkerrors.Register(ModuleName, 2, "invalid denom")
	ErrInvalidBondDenom             = sdkerrors.Register(ModuleName, 3, "invalid bond denom")
	ErrInvalidLiquidBondDenom       = sdkerrors.Register(ModuleName, 4, "invalid liquid bond denom")
	ErrLessThanMinLiquidStakeAmount = sdkerrors.Register(ModuleName, 5, "staking amount should be over params.min_liquid_staking_amount")
	ErrLiquidValidatorDoesNotExist  = sdkerrors.Register(ModuleName, 6, "liquid validators not exists")
)
