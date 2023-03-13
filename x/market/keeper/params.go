package keeper

import (
	// "time"

	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ollo-station/ollo/x/market/types"
)

// GetParams gets the parameters for the marketplace module.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// GetSaleCommission returns the current sale commission of marketplace.
func (k Keeper) GetSaleCommission(ctx sdk.Context) (percent sdk.Dec) {
	k.paramstore.Get(ctx, types.ParamStoreKeySaleCommission, &percent)
	return percent
}

// GetMarketplaceDistributionParams returns the current distribution  of marketplace commission.
func (k Keeper) GetMarketplaceDistributionParams(ctx sdk.Context) (distParams types.Distribution) {
	k.paramstore.Get(ctx, types.ParamStoreKeyDistribution, &distParams)
	return distParams
}

// GetBidCloseDuration returns the closing duration for bid for auctions.
func (k Keeper) GetBidCloseDuration(ctx sdk.Context) (duration time.Duration) {
	k.paramstore.Get(ctx, types.ParamStoreKeyBidCloseDuration, &duration)
	return duration
}
