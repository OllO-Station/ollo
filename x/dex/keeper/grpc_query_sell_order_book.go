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

func (k Keeper) SellOrderBookAll(c context.Context, req *types.QueryAllSellOrderBookRequest) (*types.QueryAllSellOrderBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sellOrderBooks []types.SellOrderBook
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sellOrderBookStore := prefix.NewStore(store, types.KeyPrefix(types.SellOrderBookKeyPrefix))

	pageRes, err := query.Paginate(sellOrderBookStore, req.Pagination, func(key []byte, value []byte) error {
		var sellOrderBook types.SellOrderBook
		if err := k.cdc.Unmarshal(value, &sellOrderBook); err != nil {
			return err
		}

		sellOrderBooks = append(sellOrderBooks, sellOrderBook)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSellOrderBookResponse{SellOrderBook: sellOrderBooks, Pagination: pageRes}, nil
}

func (k Keeper) SellOrderBook(c context.Context, req *types.QueryGetSellOrderBookRequest) (*types.QueryGetSellOrderBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSellOrderBook(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSellOrderBookResponse{SellOrderBook: val}, nil
}
