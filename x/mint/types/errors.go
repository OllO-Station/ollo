package types

// DONTCOVER

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type (
	InvalidProportionError struct {
		Proportion sdk.Dec
	}
	InsufficientVestingBalanceError struct {
		Balance   sdk.Int
		Requested sdk.Int
	}
)

// x/inflation module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)

func (e InvalidProportionError) Error() string {
	return fmt.Sprintf("invalid proportion: %s", e.Proportion)
}

func (e InsufficientVestingBalanceError) Error() string {
	return fmt.Sprintf("insufficient vesting balance: %s, requested: %s", e.Balance, e.Requested)
}
