package keeper

import (
	// "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/x/bank"
)

const (
	defaultLockPeriod = 60 * 60 * 24 * 7 // 1 week
)

func BasicLockInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return "", false
	}
}

func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	// ir.RegisterRoute(types.ModuleName, "nonnegative-outstanding",
	//     NonnegativeOutstandingInvariant(k))
}
