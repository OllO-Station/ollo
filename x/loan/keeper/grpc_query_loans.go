package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ollo/x/loan/types"
)

func (k Keeper) LoansAll(c context.Context, req *types.QueryAllLoansRequest) (*types.QueryAllLoansResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var loanss []types.Loans
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	loansStore := prefix.NewStore(store, types.KeyPrefix(types.LoansKey))

	pageRes, err := query.Paginate(loansStore, req.Pagination, func(key []byte, value []byte) error {
		var loans types.Loans
		if err := k.cdc.Unmarshal(value, &loans); err != nil {
			return err
		}

		loanss = append(loanss, loans)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLoansResponse{Loans: loanss, Pagination: pageRes}, nil
}

func (k Keeper) Loans(c context.Context, req *types.QueryGetLoansRequest) (*types.QueryGetLoansResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	loans, found := k.GetLoans(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLoansResponse{Loans: loans}, nil
}
