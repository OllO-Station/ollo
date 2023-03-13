package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/ollo-station/ollo/x/market/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// Params queries params of marketplace module
func (k Keeper) Params(
	c context.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) NftListing(
	goCtx context.Context,
	req *types.QueryNftListingRequest,
) (*types.QueryNftListingResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	listing, found := k.GetNftListing(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "listing %d not found", req.Id)
	}

	return &types.QueryNftListingResponse{Listing: &listing}, nil
}

func (k Keeper) NftListings(
	goCtx context.Context,
	req *types.QueryNftListingsRequest,
) (*types.QueryNftListingsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var listings []types.NftListing
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	var owner sdk.AccAddress
	var err error
	if len(req.Owner) > 0 {
		owner, err = sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				fmt.Sprintf("invalid owner address (%s)", err),
			)
		}
		listingStore := prefix.NewStore(
			store,
			append(types.PrefixNftListingOwner, owner.Bytes()...),
		)
		pageRes, err = query.Paginate(
			listingStore,
			req.Pagination,
			func(key []byte, value []byte) error {
				var listingId gogotypes.StringValue
				k.cdc.MustUnmarshal(value, &listingId)
				listing, found := k.GetNftListing(ctx, listingId.Value)
				if found {
					listings = append(listings, listing)
				}
				return nil
			},
		)

	} else if len(req.Denom) > 0 {
		listingStore := prefix.NewStore(store, types.KeyNftListingPriceDenomPrefix(req.Denom, ""))
		pageRes, err = query.Paginate(listingStore, req.Pagination, func(key []byte, value []byte) error {
			var listingId gogotypes.StringValue
			k.cdc.MustUnmarshal(value, &listingId)
			listing, found := k.GetNftListing(ctx, listingId.Value)
			if found {
				listings = append(listings, listing)
			}
			return nil
		})
	} else {

		listingStore := prefix.NewStore(store, types.PrefixNftListingId)
		pageRes, err = query.Paginate(listingStore, req.Pagination, func(key []byte, value []byte) error {
			var listing types.NftListing
			k.cdc.MustUnmarshal(value, &listing)
			listings = append(listings, listing)
			return nil
		})
	}
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	var ls []*types.NftListing
	for _, l := range listings {
		ls = append(ls, &l)
	}
	return &types.QueryNftListingsResponse{Listings: ls, Pagination: pageRes}, nil
}

func (k Keeper) NftListingsByOwner(
	goCtx context.Context,
	req *types.QueryNftListingsByOwnerRequest,
) (*types.QueryNftListingsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var owner sdk.AccAddress
	var err error
	if len(req.Owner) > 0 {
		owner, err = sdk.AccAddressFromBech32(req.Owner)
		if err != nil || owner == nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				fmt.Sprintf("invalid owner address (%s)", err),
			)
		}
	}

	var listings []types.NftListing
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	listingStore := prefix.NewStore(store, append(types.PrefixNftListingOwner, owner.Bytes()...))
	pageRes, err = query.Paginate(
		listingStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var listingId gogotypes.StringValue
			k.cdc.MustUnmarshal(value, &listingId)
			listing, found := k.GetNftListing(ctx, listingId.Value)
			if found {
				listings = append(listings, listing)
			}
			return nil
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	var ls []*types.NftListing
	for _, l := range listings {
		ls = append(ls, &l)
	}
	return &types.QueryNftListingsByOwnerResponse{Listings: ls, Pagination: pageRes}, nil
}

func (k Keeper) NftListingsByDenom(
	goCtx context.Context,
	req *types.QueryNftListingsByDenomRequest,
) (*types.QueryNftListingsByDenomResponse, error) {

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var err error

	var listings []types.NftListing
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	listingStore := prefix.NewStore(store, types.KeyNftListingPriceDenomPrefix(req.Denom, ""))
	pageRes, err = query.Paginate(
		listingStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var listingId gogotypes.StringValue
			k.cdc.MustUnmarshal(value, &listingId)
			listing, found := k.GetNftListing(ctx, listingId.Value)
			if found {
				listings = append(listings, listing)
			}
			return nil
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	var ls []*types.NftListing
	for _, l := range listings {
		ls = append(ls, &l)
	}
	return &types.QueryNftListingsByDenomResponse{Listings: ls, Pagination: pageRes}, nil
}

func (k Keeper) NftListingByNft(
	goCtx context.Context,
	req *types.QueryNftListingByNftRequest,
) (*types.QueryNftListingByNftResponse, error) {

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if req.NftId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "need nft id to request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	listingId, found := k.GetNftListingIdByNftId(ctx, req.NftId)
	if found {
		listing, err := k.NftListingByNft(goCtx, &types.QueryNftListingByNftRequest{
			NftId: listingId,
		})
		if err != nil {
			return nil, err
		}
		return listing, nil
	}
	return nil, status.Errorf(codes.NotFound, "listing not found with given nft id")
}

func (k Keeper) NftAuctions(
	goCtx context.Context,
	req *types.QueryNftAuctionsRequest,
) (*types.QueryNftAuctionsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var filteredAuctions []types.NftAuction
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)
	auctionStore := prefix.NewStore(store, types.PrefixNftAuctionId)
	pageRes, err := query.FilteredPaginate(
		auctionStore,
		req.Pagination,
		func(key []byte, value []byte, accumulate bool) (bool, error) {
			var al types.NftAuction
			k.cdc.MustUnmarshal(value, &al)
			matchOwner, matchPriceDenom, matchStatus := true, true, true
			// match status (if supplied/valid)
			if types.ValidAuctionStatus(req.Status) {
				if req.Status == types.NftAuctionStatusActive {
					matchStatus = al.StartTime.Before(time.Now())
				} else {
					matchStatus = al.StartTime.After(time.Now())
				}
			}

			// match owner address (if supplied)
			if len(req.Owner) > 0 {
				owner, err := sdk.AccAddressFromBech32(req.Owner)
				if err != nil {
					return false, err
				}

				matchOwner = al.Owner == owner.String()
			}

			// match Price Denom (if supplied)
			if len(req.Denom) > 0 {
				matchPriceDenom = al.StartPrice.Denom == req.Denom
			}

			if matchOwner && matchPriceDenom && matchStatus {
				if accumulate {
					filteredAuctions = append(filteredAuctions, al)
				}

				return true, nil
			}

			return false, nil
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}
	var ls []*types.NftAuction
	for _, l := range filteredAuctions {
		ls = append(ls, &l)
	}

	return &types.QueryNftAuctionsResponse{Auctions: ls, Pagination: pageRes}, nil
}

func (k Keeper) NftAuction(
	goCtx context.Context,
	req *types.QueryNftAuctionRequest,
) (*types.QueryNftAuctionResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	auction, found := k.GetNftAuction(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "auction %d not found", req.Id)
	}
	return &types.QueryNftAuctionResponse{Auction: &auction}, nil
}

func (k Keeper) NftAuctionsByOwner(
	goCtx context.Context,
	req *types.QueryNftAuctionsByOwnerRequest,
) (*types.QueryNftAuctionsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var owner sdk.AccAddress
	var err error
	if len(req.Owner) > 0 {
		owner, err = sdk.AccAddressFromBech32(req.Owner)
		if err != nil || owner == nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				fmt.Sprintf("invalid owner address (%s)", err),
			)
		}
	}

	var auctions []types.NftAuction
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	auctionStore := prefix.NewStore(store, append(types.PrefixNftAuctionOwner, owner.Bytes()...))
	pageRes, err = query.Paginate(
		auctionStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var auctionId gogotypes.UInt64Value
			k.cdc.MustUnmarshal(value, &auctionId)
			auction, found := k.GetNftAuction(ctx, auctionId.Value)
			if found {
				auctions = append(auctions, auction)
			}
			return nil
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	var ls []*types.NftAuction
	for _, l := range auctions {
		ls = append(ls, &l)
	}
	return &types.QueryNftAuctionsByOwnerResponse{Auctions: ls, Pagination: pageRes}, nil
}

func (k Keeper) NftAuctionsByDenom(
	goCtx context.Context,
	req *types.QueryNftAuctionsByDenomRequest,
) (*types.QueryNftAuctionsByDenomResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var err error

	var auctions []types.NftAuction
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	auctionStore := prefix.NewStore(
		store,
		append(types.PrefixNftAuctionPriceDenom, []byte(req.Denom)...),
	)
	pageRes, err = query.Paginate(
		auctionStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var auctionId gogotypes.UInt64Value
			k.cdc.MustUnmarshal(value, &auctionId)
			auction, found := k.GetNftAuction(ctx, auctionId.Value)
			if found {
				auctions = append(auctions, auction)
			}
			return nil
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	var ls []*types.NftAuction
	for _, l := range auctions {
		ls = append(ls, &l)
	}
	return &types.QueryNftAuctionsByDenomResponse{Auctions: ls, Pagination: pageRes}, nil
}

func (k Keeper) NftAuctionByNft(
	goCtx context.Context,
	req *types.QueryNftAuctionByNftRequest,
) (*types.QueryNftAuctionByNftResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	if req.NftId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "need nft id to request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetNftAuctionIdByNftId(ctx, req.NftId)
	if found {
		auction, err := k.NftAuctionByNft(goCtx, &types.QueryNftAuctionByNftRequest{
			NftId: req.NftId,
		})
		if err != nil {
			return nil, err
		}
		return auction, nil
	}
	return nil, status.Errorf(codes.NotFound, "auction not found with given nft id")
}

func (k Keeper) NftAuctionBid(
	goCtx context.Context,
	req *types.QueryNftAuctionBidRequest,
) (*types.QueryNftAuctionBidResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	bid, found := k.GetNftAuctionBid(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "bid not found for auction %d", req.Id)
	}
	return &types.QueryNftAuctionBidResponse{Bid: &bid}, nil
}

func (k Keeper) NftAuctionBids(
	goCtx context.Context,
	req *types.QueryNftAuctionBidsRequest,
) (*types.QueryNftAuctionBidsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var bids []types.NftAuctionBid
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	if len(req.Bidder) > 0 {
		_, err := sdk.AccAddressFromBech32(req.Bidder)
		if err != nil {
			return nil, err
		}
	}

	bidStore := prefix.NewStore(store, types.PrefixBidByNftAuctionId)
	pageRes, err := query.Paginate(bidStore, req.Pagination, func(key []byte, value []byte) error {
		var bid types.NftAuctionBid
		k.cdc.MustUnmarshal(value, &bid)
		if len(req.Bidder) > 0 {
			if bid.Bidder == req.Bidder {
				bids = append(bids, bid)
			}
		} else {
			bids = append(bids, bid)
		}
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}
	var ls []*types.NftAuctionBid
	for _, l := range bids {
		ls = append(ls, &l)
	}
	return &types.QueryNftAuctionBidsResponse{Bids: ls, Pagination: pageRes}, nil
}
