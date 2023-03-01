package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ollo-station/ollo/x/claim/types"
)

func (k Keeper) GoalAll(c context.Context, req *types.QueryAllGoalRequest) (*types.QueryAllGoalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var goals []types.Goal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	goalStore := prefix.NewStore(store, types.KeyPrefix(types.GoalKey))

	pageRes, err := query.Paginate(goalStore, req.Pagination, func(key []byte, value []byte) error {
		var goal types.Goal
		if err := k.cdc.Unmarshal(value, &goal); err != nil {
			return err
		}

		goals = append(goals, goal)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGoalResponse{Goal: goals, Pagination: pageRes}, nil
}

func (k Keeper) Goal(c context.Context, req *types.QueryGetGoalRequest) (*types.QueryGetGoalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	goal, found := k.GetGoal(ctx, req.GoalID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetGoalResponse{Goal: goal}, nil
}
