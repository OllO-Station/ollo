package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/x/reserve/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) GetDenomWhitelist(ctx context.Context, req *types.QueryGetDenomWhitelistRequest) (*types.QueryGetDenomWhitelistResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	authorityMetadata, err := k.GetWhitelist(sdkCtx, req.GetDenom())
	if err != nil {
		return nil, err
	}

	return &types.QueryGetDenomWhitelistResponse{Whitelist: authorityMetadata}, nil
}

func (k Keeper) DenomsFromCreator(ctx context.Context, req *types.QueryDenomsFromCreatorRequest) (*types.QueryDenomsFromCreatorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	denoms := k.getDenomsFromCreator(sdkCtx, req.GetCreator())
	return &types.QueryDenomsFromCreatorResponse{Denoms: denoms}, nil
}

func (k Keeper) GetDenoms(ctx context.Context, req *types.QueryGetDenomsRequest) (*types.QueryGetDenomsResponse, error) {
	// sdkCtx := sdk.UnwrapSDKContext(ctx)
	// denoms := k.getDenomsFromCreator(sdkCtx, req.GetCreator())
	return &types.QueryGetDenomsResponse{}, nil
}

func (k Keeper) GetDenom(ctx context.Context, req *types.QueryGetDenomRequest) (*types.QueryGetDenomResponse, error) {
	// sdkCtx := sdk.UnwrapSDKContext(ctx)
	// denoms := k.getDenomsFromCreator(sdkCtx, req.GetCreator())
	return &types.QueryGetDenomResponse{}, nil
}
