package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ollo/x/dex/types"
)

func (k Keeper) BuyOrderBookAll(c context.Context, req *types.QueryAllBuyOrderBookRequest) (*types.QueryAllBuyOrderBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var buyOrderBooks []types.BuyOrderBook
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	buyOrderBookStore := prefix.NewStore(store, types.KeyPrefix(types.BuyOrderBookKeyPrefix))

	pageRes, err := query.Paginate(buyOrderBookStore, req.Pagination, func(key []byte, value []byte) error {
		var buyOrderBook types.BuyOrderBook
		if err := k.cdc.Unmarshal(value, &buyOrderBook); err != nil {
			return err
		}

		buyOrderBooks = append(buyOrderBooks, buyOrderBook)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBuyOrderBookResponse{BuyOrderBook: buyOrderBooks, Pagination: pageRes}, nil
}

func (k Keeper) BuyOrderBook(c context.Context, req *types.QueryGetBuyOrderBookRequest) (*types.QueryGetBuyOrderBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBuyOrderBook(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBuyOrderBookResponse{BuyOrderBook: val}, nil
}
