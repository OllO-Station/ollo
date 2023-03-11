package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AfterEpochEnd(ctx sdk.Context, id string, num uint64) {
	_ = k.hooks.End(ctx, id, num)
}

func (k Keeper) AfterEpochStart(ctx sdk.Context, id string, num uint64) {
	_ = k.hooks.Start(ctx, id, num)
}
