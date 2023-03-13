package types

var (
	EventListNFT                = "list_nft"
	EventEditNFTListing         = "edit_listing"
	EventDeListNFT              = "de_list_nft"
	EventBuyNFT                 = "buy_nft"
	EventRoyaltyShareTransfer   = "royalty_share_transfer"
	EventSplitShareTransfer     = "split_share_transfer"
	EventSaleCommissionTransfer = "sale_commission_transfer"

	EventCreateAuction = "create_auction"
	EventCancelAuction = "cancel_auction"
	EventPlaceBid      = "place_bid"
	EventRemoveAuction = "remove_auction"
	EventProcessBid    = "process_bid"

	AttributeValueCategory   = ModuleName
	AttributeKeyNftListingId = "listing-id"
	AttributeKeyDenomId      = "denom-id"
	AttributeKeyNftId        = "nft-id"
	AttributeKeyBuyer        = "buyer"
	AttributeKeyOwner        = "owner"
	AttributeKeyRecipient    = "recipient"
	AttributeKeyAmount       = "amount"
	AttributeKeyAuctionId    = "auction-id"
	AttributeKeyStartPrice   = "start-price"
	AttributeKeyBidder       = "bidder"

	AttributeKeySender = "sender"
)
