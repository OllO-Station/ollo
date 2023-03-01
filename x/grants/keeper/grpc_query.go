package keeper

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/ollo-station/ollo/x/grants/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Params queries the parameters of the fundraising module.
func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.Keeper.paramSpace.GetParamSet(ctx, &params)
	return &types.QueryParamsResponse{Params: params}, nil
}

// Auctions queries all auctions.
func (k Querier) Auctions(c context.Context, req *types.QueryAuctionsRequest) (*types.QueryAuctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.Type != "" && !(req.Type == types.AuctionTypeFixedPrice.String() || req.Type == types.AuctionTypeBatch.String()) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid auction type %s", req.Type)
	}

	if req.Status != "" && !(req.Status == types.AuctionStatusStandBy.String() || req.Status == types.AuctionStatusStarted.String() ||
		req.Status == types.AuctionStatusVesting.String() || req.Status == types.AuctionStatusFinished.String() ||
		req.Status == types.AuctionStatusCancelled.String()) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid auction status %s", req.Status)
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	auctionStore := prefix.NewStore(store, types.AuctionKeyPrefix)

	var auctions []*codectypes.Any
	pageRes, err := query.FilteredPaginate(auctionStore, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		auction, err := types.UnmarshalAuction(k.cdc, value)
		if err != nil {
			return false, err
		}

		auctionAny, err := types.PackAuction(auction)
		if err != nil {
			return false, err
		}

		if req.Type != "" && auction.GetType().String() != req.Type {
			return false, nil
		}

		if req.Status != "" && auction.GetStatus().String() != req.Status {
			return false, nil
		}

		if accumulate {
			auctions = append(auctions, auctionAny)
		}

		return true, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAuctionsResponse{Auctions: auctions, Pagination: pageRes}, nil
}

// Auction queries the specific auction.
func (k Querier) Auction(c context.Context, req *types.QueryAuctionRequest) (*types.QueryAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	auction, found := k.Keeper.GetAuction(ctx, req.AuctionId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "auction %d not found", req.AuctionId)
	}

	auctionAny, err := types.PackAuction(auction)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAuctionResponse{Auction: auctionAny}, nil
}

// AllowedBidder queries the specific allowed bidder information.
func (k Querier) AllowedBidder(c context.Context, req *types.QueryAllowedBidderRequest) (*types.QueryAllowedBidderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.AuctionId == 0 {
		return nil, status.Error(codes.InvalidArgument, "auction id cannot be 0")
	}

	if req.Bidder == "" {
		return nil, status.Error(codes.InvalidArgument, "empty bidder address")
	}

	ctx := sdk.UnwrapSDKContext(c)

	_, found := k.GetAuction(ctx, req.AuctionId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "auction %d not found", req.AuctionId)
	}

	bidderAddr, err := sdk.AccAddressFromBech32(req.Bidder)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bidder address %s is not valid", req.Bidder)
	}

	allowedBidder, found := k.GetAllowedBidder(ctx, req.AuctionId, bidderAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "allowed bidder by auction id %d and bidder address %s doesn't exist", req.AuctionId, req.Bidder)
	}

	return &types.QueryAllowedBidderResponse{AllowedBidder: allowedBidder}, nil
}

// AllowedBidders queries all allowed bidders for the auction.
func (k Querier) AllowedBidders(c context.Context, req *types.QueryAllowedBiddersRequest) (*types.QueryAllowedBiddersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.AuctionId == 0 {
		return nil, status.Error(codes.InvalidArgument, "auction id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)

	_, found := k.GetAuction(ctx, req.AuctionId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "auction %d not found", req.AuctionId)
	}

	store := ctx.KVStore(k.storeKey)
	abStore := prefix.NewStore(store, types.GetAllowedBiddersByAuctionKeyPrefix(req.AuctionId))

	var allowedBidders []types.AllowedBidder
	pageRes, err := query.FilteredPaginate(abStore, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		var allowedBidder types.AllowedBidder
		err := k.cdc.Unmarshal(value, &allowedBidder)
		if err != nil {
			return false, nil
		}

		if accumulate {
			allowedBidders = append(allowedBidders, allowedBidder)
		}

		return true, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllowedBiddersResponse{AllowedBidders: allowedBidders, Pagination: pageRes}, nil
}

// Bids queries all bids for the auction.
func (k Querier) Bids(c context.Context, req *types.QueryBidsRequest) (*types.QueryBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_, found := k.Keeper.GetAuction(ctx, req.AuctionId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "auction %d not found", req.AuctionId)
	}

	var bids []types.Bid
	var pageRes *query.PageResponse
	var err error

	store := ctx.KVStore(k.storeKey)
	switch {
	case req.Bidder != "" && req.IsMatched == "":
		bids, pageRes, err = queryBidsByBidder(ctx, k, store, req)
	case req.Bidder == "" && req.IsMatched != "":
		bids, pageRes, err = queryBidsByIsMatched(ctx, k, store, req)
	case req.Bidder != "" && req.IsMatched != "":
		bids, pageRes, err = queryBidsByBidder(ctx, k, store, req)
	default:
		bids, pageRes, err = queryAllBids(ctx, k, store, req)
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBidsResponse{Bids: bids, Pagination: pageRes}, nil
}

// Bid queries the specific bid from the auction id and bid id.
func (k Querier) Bid(c context.Context, req *types.QueryBidRequest) (*types.QueryBidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	bid, found := k.Keeper.GetBid(ctx, req.AuctionId, req.BidId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "bid from auction id %d and bid id %d not found", req.AuctionId, req.BidId)
	}

	return &types.QueryBidResponse{Bid: bid}, nil
}

// Vestings queries all vesting queues for the auction.
func (k Querier) Vestings(c context.Context, req *types.QueryVestingsRequest) (*types.QueryVestingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	auction, found := k.Keeper.GetAuction(ctx, req.AuctionId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "auction %d not found", req.AuctionId)
	}

	queues := k.Keeper.GetVestingQueuesByAuctionId(ctx, auction.GetId())

	return &types.QueryVestingsResponse{Vestings: queues}, nil
}

func queryAllBids(ctx sdk.Context, k Querier, store sdk.KVStore, req *types.QueryBidsRequest) (bids []types.Bid, pageRes *query.PageResponse, err error) {
	bidStore := prefix.NewStore(store, types.BidKeyPrefix)

	pageRes, err = query.FilteredPaginate(bidStore, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		var bid types.Bid
		if err := k.cdc.Unmarshal(value, &bid); err != nil {
			return false, nil
		}

		if bid.AuctionId != req.AuctionId {
			return false, nil
		}

		if accumulate {
			bids = append(bids, bid)
		}

		return true, nil
	})
	if err != nil {
		return nil, nil, status.Error(codes.Internal, err.Error())
	}

	return bids, pageRes, err
}

func queryBidsByBidder(ctx sdk.Context, k Querier, store sdk.KVStore, req *types.QueryBidsRequest) (bids []types.Bid, pageRes *query.PageResponse, err error) {
	bidderAddr, err := sdk.AccAddressFromBech32(req.Bidder)
	if err != nil {
		return nil, nil, err
	}

	bidStore := prefix.NewStore(store, types.GetBidIndexByBidderPrefix(bidderAddr))

	pageRes, err = query.FilteredPaginate(bidStore, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		auctionId, bidId := types.SplitAuctionIdBidIdKey(key)
		bid, _ := k.GetBid(ctx, auctionId, bidId)

		if req.Bidder != bid.Bidder {
			return false, nil
		}

		if req.IsMatched != "" {
			isMatched, err := strconv.ParseBool(req.IsMatched)
			if err != nil {
				return false, err
			}

			if bid.IsMatched != isMatched {
				return false, nil
			}
		}

		if accumulate {
			bids = append(bids, bid)
		}

		return true, nil
	})

	if err != nil {
		return nil, nil, status.Error(codes.Internal, err.Error())
	}

	return bids, pageRes, err
}

func queryBidsByIsMatched(ctx sdk.Context, k Querier, store sdk.KVStore, req *types.QueryBidsRequest) (bids []types.Bid, pageRes *query.PageResponse, err error) {
	isMatched, err := strconv.ParseBool(req.IsMatched)
	if err != nil {
		return nil, nil, err
	}

	bidStore := prefix.NewStore(store, types.BidKeyPrefix)

	pageRes, err = query.FilteredPaginate(bidStore, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		var bid types.Bid
		if err := k.cdc.Unmarshal(value, &bid); err != nil {
			return false, nil
		}

		if bid.AuctionId != req.AuctionId {
			return false, nil
		}

		if bid.IsMatched != isMatched {
			return false, nil
		}

		if accumulate {
			bids = append(bids, bid)
		}

		return true, nil
	})
	if err != nil {
		return nil, nil, status.Error(codes.Internal, err.Error())
	}

	return bids, pageRes, err
}
