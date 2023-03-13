package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ollo-station/ollo/x/market/types"
	"time"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the token MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}
func (m msgServer) ListNft(goCtx context.Context, msg *types.MsgListNft) (*types.MsgListNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return nil, err
	}

	nft, err := m.nftKeeper.GetNFT(ctx, msg.DenomId, msg.NftId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrNftNotExists,
			"invalid nft and or denomId, nftId %s, denomId %s", msg.NftId, msg.DenomId)
	}
	if owner.String() != nft.GetOwner().String() {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized address %s", owner)
	}
	if !nft.IsTransferable() {
		return nil, sdkerrors.Wrapf(
			types.ErrNftNonTransferable, "non-transferable nfts not allowed to list in marketplace")
	}

	listing := types.NewNftListing(msg.Id, msg.NftId, msg.DenomId, msg.Price, owner)
	err = m.Keeper.AddNftListing(ctx, listing)
	if err != nil {
		return nil, err
	}

	m.Keeper.CreateListNftEvent(ctx, owner, listing.Id, listing.DenomId, listing.NftId, listing.Price)

	return &types.MsgListNftResponse{}, nil
}

func (m msgServer) EditNftListing(goCtx context.Context,
	msg *types.MsgEditNftListing) (*types.MsgEditNftListingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	listing, found := m.Keeper.GetNftListing(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNftListingDoesNotExists, "listing id %s not exists", msg.Id)
	}
	if owner.String() != listing.Creator {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized address %s", owner)
	}
	if err := types.ValidatePrice(msg.Price); err != nil {
		return nil, err
	}
	listing.Price = msg.Price
	m.Keeper.SetNftListing(ctx, listing)

	m.Keeper.CreateEditNftListingEvent(ctx, owner, listing.Id, listing.Price)

	return &types.MsgEditNftListingResponse{}, nil
}

func (m msgServer) DelistNft(goCtx context.Context,
	msg *types.MsgDelistNft) (*types.MsgDelistNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}
	listing, found := m.Keeper.GetNftListing(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNftListingDoesNotExists, "listing id %s not exists", msg.Id)
	}
	if owner.String() != listing.Creator {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized address %s", owner)
	}
	err = m.nftKeeper.TransferNFT(ctx, listing.GetDenomId(), listing.GetNftId(),
		m.accountKeeper.GetModuleAddress(types.ModuleName), listing.GetOwner())
	if err != nil {
		return nil, err
	}
	m.Keeper.DeleteNftListing(ctx, listing)

	m.Keeper.CreateDeListNftEvent(ctx, owner, listing.Id)

	return &types.MsgDelistNftResponse{}, nil
}

func (m msgServer) BuyNft(goCtx context.Context, msg *types.MsgBuyNft) (*types.MsgBuyNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return nil, err
	}

	listing, found := m.Keeper.GetNftListing(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNftListingDoesNotExists, "listing id %s not exists", msg.Id)
	}
	if err := types.ValidatePrice(msg.Price); err != nil {
		return nil, err
	}
	if !msg.Price.Equal(listing.Price) {
		if msg.Price.Denom != listing.Price.Denom {
			return nil, sdkerrors.Wrapf(types.ErrInvalidPriceDenom, "invalid price denom %s", msg.Price.Denom)
		}
		if msg.Price.Amount.LT(listing.Price.Amount) {
			return nil, sdkerrors.Wrapf(types.ErrNotEnoughAmount,
				"%s is not enough, to buy %s required", msg.Price.String(), listing.Price.String())
		}
		return nil, sdkerrors.Wrapf(types.ErrInvalidPrice,
			"price %s not matched with listing price", msg.Price.String())
	}
	err = m.Keeper.Buy(ctx, listing, buyer)
	if err != nil {
		return nil, err
	}

	m.Keeper.CreateBuyNftEvent(ctx, buyer, listing.Id, listing.NftId, listing.Price)

	return &types.MsgBuyNftResponse{}, nil
}

// CreateAuction
func (m msgServer) CreateNftAuction(goCtx context.Context, msg *types.MsgCreateNftAuction) (*types.MsgCreateNftAuctionResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}
	if err := msg.Validate(ctx.BlockTime()); err != nil {
		return nil, err
	}

	nft, err := m.nftKeeper.GetNFT(ctx, msg.DenomId, msg.NftId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrNftNotExists,
			"invalid nft and or denomId, nftId %s, denomId %s", msg.NftId, msg.DenomId)
	}
	if owner.String() != nft.GetOwner().String() {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized address %s", owner)
	}
	if !nft.IsTransferable() {
		return nil, sdkerrors.Wrapf(
			types.ErrNftNonTransferable, "non-transferable nfts not allowed to list in marketplace")
	}
	var endTime *time.Time
	if msg.Duration != nil {
		endAt := msg.StartTime.Add(*msg.Duration)
		endTime = &endAt
		if endTime.Before(msg.StartTime) || endTime.Equal(msg.StartTime) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDuration, "duration must be positive or nil")
		}
	}
	auctionNumber := m.Keeper.GetNextNftAuctionNumber(ctx)
	auction := types.NewNftAuction(auctionNumber, msg.NftId, msg.DenomId,
		msg.StartTime, endTime, msg.StartPrice,
		owner, msg.IncrementPercentage)
	err = m.Keeper.CreateNftAuction(ctx, auction)
	if err != nil {
		return nil, err
	}

	m.Keeper.CreateNftAuctionEvent(ctx, auction)

	return &types.MsgCreateNftAuctionResponse{
		Auction: &auction,
	}, nil
}

// CancelAuction
func (m msgServer) CancelNftAuction(goCtx context.Context, msg *types.MsgCancelNftAuction) (*types.MsgCancelNftAuctionResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	auction, found := m.Keeper.GetNftAuction(ctx, msg.AuctionId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNftAuctionDoesNotExists, "auction id %d not exists", msg.AuctionId)
	}
	if owner.String() != auction.Owner {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized address %s", owner.String())
	}

	err = m.Keeper.CancelNftAuction(ctx, auction)
	if err != nil {
		return nil, err
	}

	m.Keeper.CancelNftAuctionEvent(ctx, auction)

	return &types.MsgCancelNftAuctionResponse{}, nil
}

// PlaceBid
func (m msgServer) PlaceNftBid(goCtx context.Context, msg *types.MsgPlaceNftBid) (*types.MsgPlaceNftBidResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	bidder, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		return nil, err
	}

	auction, found := m.Keeper.GetNftAuction(ctx, msg.AuctionId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNftAuctionDoesNotExists, "auction id %d not exists", msg.AuctionId)
	}
	if !auction.StartTime.Before(ctx.BlockTime()) {
		return nil, sdkerrors.Wrapf(types.ErrInActiveNftAuction, "cannot place a bid for inactive auction %d, ", auction.Id)
	}
	if msg.Amount.GetDenom() != auction.StartPrice.GetDenom() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPriceDenom, "given auction only accepts bids in %s, ", auction.StartPrice.GetDenom())
	}

	bid := types.NewBid(auction.Id, msg.Amount, ctx.BlockTime(), bidder)

	err = m.Keeper.PlaceNftBid(ctx, auction, bid)
	if err != nil {
		return nil, err
	}

	m.Keeper.PlaceNftAuctionBidEvent(ctx, auction, bid)

	return &types.MsgPlaceNftBidResponse{}, nil
}
