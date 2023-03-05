package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	types "github.com/ollo-station/ollo/x/epoch/types"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/epochs keeper providing gRPC method
// handlers.
type Querier struct {
	Keeper
}

// NewQuerier initializes new querier.
func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}

// EpochInfos provide running epochInfos.
func (q Querier) Epochs(c context.Context, _ *types.QueryEpochsRequest) (*types.QueryEpochsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryEpochsResponse{
		Epochs: q.Keeper.AllEpochs(ctx),
	}, nil
}
func (q Querier) Epoch(c context.Context, r *types.QueryEpochsRequest) (*types.QueryEpochsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryEpochsResponse{
		Epochs: q.Keeper.GetEpoch(ctx, r.Id),
	}, nil
}

// CurrentEpoch provides current epoch of specified identifier.
func (q Querier) CurrentEpoch(c context.Context, req *types.QueryCurrentEpochRequest) (*types.QueryCurrentEpochResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "identifier is empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	info := q.Keeper.GetEpoch(ctx, req.Id)
	if info.Id != req.Id {
		return nil, errors.New("not available identifier")
	}

	return &types.QueryCurrentEpochResponse{
		CurrentEpoch: info.CurrentEpoch,
	}, nil
}
