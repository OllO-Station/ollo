package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/prices/types"
)

// BandPricesResult returns the BandPrices result by RequestId
func (k Keeper) BandPricesResult(c context.Context, req *types.QueryBandPricesRequest) (*types.QueryBandPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	result, err := k.GetBandPricesResult(ctx, types.OracleRequestID(req.RequestId))
	if err != nil {
		return nil, err
	}
	return &types.QueryBandPricesResponse{Result: &result}, nil
}

// LastBandPricesId returns the last BandPrices request Id
func (k Keeper) LastBandPricesId(c context.Context, req *types.QueryLastBandPricesIdRequest) (*types.QueryLastBandPricesIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	id := k.GetLastBandPricesID(ctx)
	return &types.QueryLastBandPricesIdResponse{RequestId: id}, nil
}
