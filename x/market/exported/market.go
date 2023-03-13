package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// NftListingI is an interface for NftListing
type NftListingI interface {
	GetId() string
	GetDenomId() string
	GetNftId() string
	GetPrice() sdk.Coin
	GetOwner() sdk.AccAddress
}

// AuctionI is an interface for NftAuction
type NftAuctionI interface {
	GetId() uint64
	GetDenomId() string
	GetNftId() string
	GetStartPrice() sdk.Coin
	GetStartTime() time.Time
	// GetEndTime() time.Time
	GetOwner() sdk.AccAddress
	GetStatus() string
}

// NftBidI is an interface for NftBid
type NftBidI interface {
	// GetId() string
	// GetDenom() string
	// GetNftId() string
	GetAuctionId() uint64
	GetBidder() sdk.AccAddress
	GetAmount() sdk.Coin
}
