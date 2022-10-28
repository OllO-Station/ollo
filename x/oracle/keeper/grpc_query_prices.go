package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/oracle/types"
)

// PricesResult returns the Prices result by RequestId
func (k Keeper) PricesResult(c context.Context, req *types.QueryPricesRequest) (*types.QueryPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	result, err := k.GetPricesResult(ctx, types.OracleRequestID(req.RequestId))
	if err != nil {
		return nil, err
	}
	return &types.QueryPricesResponse{Result: &result}, nil
}

// LastPricesId returns the last Prices request Id
func (k Keeper) LastPricesId(c context.Context, req *types.QueryLastPricesIdRequest) (*types.QueryLastPricesIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	id := k.GetLastPricesID(ctx)
	return &types.QueryLastPricesIdResponse{RequestId: id}, nil
}
