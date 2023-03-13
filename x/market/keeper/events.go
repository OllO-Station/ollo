package keeper

import (
	// "fmt"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ollo-station/ollo/x/market/types"
)

func (k *Keeper) CreateListNftEvent(
	ctx sdk.Context,
	owner sdk.AccAddress,
	listingId string,
	denomId string,
	nftId string,
	price sdk.Coin,
) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventListNFT,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyNftListingId, listingId),
			sdk.NewAttribute(types.AttributeKeyDenomId, denomId),
			sdk.NewAttribute(types.AttributeKeyNftId, nftId),
			// sdk.NewAttribute(types.AttributeKeyPrice, price.String()),
		),
	})
}

func (k *Keeper) CreateDeListNftEvent(ctx sdk.Context, sender sdk.AccAddress, listingId string) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventDeListNFT,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyNftListingId, listingId),
		),
	})
}

func (k *Keeper) CreateEditNftListingEvent(ctx sdk.Context, sender sdk.AccAddress, listingId string, price sdk.Coin) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventEditNFTListing,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyNftListingId, listingId),
			sdk.NewAttribute(types.AttributeKeyAmount, price.String()),
		),
	})
}

func (k *Keeper) CreateBuyNftEvent(ctx sdk.Context, buyer sdk.AccAddress, listId, nftId string, price sdk.Coin) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventBuyNFT,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyBuyer, buyer.String()),
			sdk.NewAttribute(types.AttributeKeyNftListingId, listId),
			sdk.NewAttribute(types.AttributeKeyNftId, nftId),
			sdk.NewAttribute(types.AttributeKeyAmount, price.String()),
		),
	})
}

func (k *Keeper) CreateRoyaltyShareTransferEvent(ctx sdk.Context, sender, recipient sdk.AccAddress, amount sdk.Coin) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventRoyaltyShareTransfer,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	})
}

func (k *Keeper) CreateSplitShareTransferEvent(ctx sdk.Context, sender, recipient sdk.AccAddress, amount sdk.Coin) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventSplitShareTransfer,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	})
}

func (k *Keeper) CreateCommissionTransferEvent(ctx sdk.Context, sender, recipient sdk.AccAddress, amount sdk.Coin) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventSaleCommissionTransfer,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
	})
}

func (k *Keeper) CreateNftAuctionEvent(ctx sdk.Context, auction types.NftAuction) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventCreateAuction,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyOwner, auction.GetOwner().String()),
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(auction.GetId())),
			sdk.NewAttribute(types.AttributeKeyDenomId, auction.GetDenomId()),
			sdk.NewAttribute(types.AttributeKeyNftId, auction.GetNftId()),
			sdk.NewAttribute(types.AttributeKeyStartPrice, auction.GetStartPrice().String()),
		),
	})
}

func (k *Keeper) CancelNftAuctionEvent(ctx sdk.Context, auction types.NftAuction) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventCancelAuction,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyOwner, auction.GetOwner().String()),
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(auction.GetId())),
			sdk.NewAttribute(types.AttributeKeyDenomId, auction.GetDenomId()),
			sdk.NewAttribute(types.AttributeKeyNftId, auction.GetNftId()),
		),
	})
}

func (k *Keeper) PlaceNftAuctionBidEvent(ctx sdk.Context, auction types.NftAuction, bid types.NftAuctionBid) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventPlaceBid,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyBidder, bid.GetBidder().String()),
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(auction.GetId())),
			sdk.NewAttribute(types.AttributeKeyDenomId, auction.GetDenomId()),
			sdk.NewAttribute(types.AttributeKeyNftId, auction.GetNftId()),
			sdk.NewAttribute(types.AttributeKeyAmount, bid.GetAmount().String()),
		),
	})
}

func (k *Keeper) RemoveNftAuctionEvent(ctx sdk.Context, auction types.NftAuction) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventRemoveAuction,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(auction.GetId())),
			sdk.NewAttribute(types.AttributeKeyDenomId, auction.GetDenomId()),
			sdk.NewAttribute(types.AttributeKeyNftId, auction.GetNftId()),
		),
	})
}

func (k *Keeper) ProcessNftAuctionBidEvent(ctx sdk.Context, auction types.NftAuction, bid types.NftAuctionBid) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventProcessBid,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyAuctionId, fmt.Sprint(auction.GetId())),
			sdk.NewAttribute(types.AttributeKeyDenomId, auction.GetDenomId()),
			sdk.NewAttribute(types.AttributeKeyNftId, auction.GetNftId()),
			sdk.NewAttribute(types.AttributeKeyBidder, bid.GetBidder().String()),
			sdk.NewAttribute(types.AttributeKeyAmount, bid.GetAmount().String()),
		),
	})
}
