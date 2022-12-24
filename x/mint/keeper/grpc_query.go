package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/mint/types"
)

var _ types.QueryServer = Keeper{}

// Params returns params of the mint module.
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Inflation returns minter.Inflation of the mint module.
func (k Keeper) Inflation(c context.Context, _ *types.QueryInflationRequest) (*types.QueryInflationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	minter := k.GetMinter(ctx)

	return &types.QueryInflationResponse{Inflation: minter.Inflation}, nil
}

// AnnualProvisions returns minter.AnnualProvisions of the mint module.
func (k Keeper) AnnualProvisions(c context.Context, _ *types.QueryAnnualProvisionsRequest) (*types.QueryAnnualProvisionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	minter := k.GetMinter(ctx)

	return &types.QueryAnnualProvisionsResponse{AnnualProvisions: minter.AnnualProvisions}, nil
}

func (k Keeper) Distribution(c context.Context, _ *types.QueryDistributionRequest) (*types.QueryDistributionResponse, error) {
	return &types.QueryDistributionResponse{}, nil
}
func (k Keeper) LastBlockTime(c context.Context, _ *types.QueryLastBlockTimeRequest) (*types.QueryLastBlockTimeResponse, error) {
	return &types.QueryLastBlockTimeResponse{}, nil
}

// 	ctx := sdk.UnwrapSDKContext(c)
// 	minter := k.GetMinter(ctx)

// 	return &types.QueryAnnualProvisionsResponse{AnnualProvisions: minter.AnnualProvisions}, nil
// }
