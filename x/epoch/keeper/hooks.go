package keeper

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) EpochEnd(ctx sdk.Context, id string, num uint64) {
    _ = k.hooks.EpochEnd(ctx, id, num)
}

func (k Keeper) EpochStart(ctx sdk.Context, id string, num uint64) {
    _ = k.hooks.EpochStart(ctx, id, num)
}
