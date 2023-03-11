package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ollo-station/ollo/x/lend/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LoanAll(goCtx context.Context, req *types.QueryAllLoanRequest) (*types.QueryAllLoanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var loans []types.Loan
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	loanStore := prefix.NewStore(store, types.KeyPrefix(types.LoanKey))

	pageRes, err := query.Paginate(loanStore, req.Pagination, func(key []byte, value []byte) error {
		var loan types.Loan
		if err := k.cdc.Unmarshal(value, &loan); err != nil {
			return err
		}

		loans = append(loans, loan)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLoanResponse{Loan: loans, Pagination: pageRes}, nil
}

func (k Keeper) Loan(goCtx context.Context, req *types.QueryGetLoanRequest) (*types.QueryGetLoanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	loan, found := k.GetLoan(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLoanResponse{Loan: loan}, nil
}
