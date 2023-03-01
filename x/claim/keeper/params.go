package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/x/claim/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.DecayInformation(ctx),
		k.AirdropStart(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// DecayInformation returns the param that defines decay information
func (k Keeper) DecayInformation(ctx sdk.Context) (totalSupplyRange types.DecayInformation) {
	k.paramstore.Get(ctx, types.KeyDecayInformation, &totalSupplyRange)
	return
}

// AirdropStart returns the param that defines airdrop start
func (k Keeper) AirdropStart(ctx sdk.Context) (airdropStart time.Time) {
	k.paramstore.Get(ctx, types.KeyAirdropStart, &airdropStart)
	return
}
