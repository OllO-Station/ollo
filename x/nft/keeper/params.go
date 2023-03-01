package keeper

import (
	"github.com/ollo-station/ollo/x/nft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	// k.paramstore.SetParamSet(ctx, &params)
}

// // GetFeeCollector returns the current fee collector address parameter.
// func (k Keeper) GetFeeCollector(ctx sdk.Context) sdk.AccAddress {
// 	var feeCollectorAddr string
// 	k.paramstore.Get(ctx, types.KeyFeeCollectorAddress, &feeCollectorAddr)
// 	addr, err := sdk.AccAddressFromBech32(feeCollectorAddr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return addr
// }

// // GetDustCollector returns the current dust collector address parameter.
// func (k Keeper) GetDustCollector(ctx sdk.Context) sdk.AccAddress {
// 	var dustCollectorAddr string
// 	k.paramstore.Get(ctx, types.KeyDustCollectorAddress, &dustCollectorAddr)
// 	addr, err := sdk.AccAddressFromBech32(dustCollectorAddr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return addr
// }
