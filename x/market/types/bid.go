package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/ollo-station/ollo/x/market/exported"
	"time"
)

var (
	_ proto.Message    = &NftAuctionBid{}
	_ exported.NftBidI = &NftAuctionBid{}
)

func NewBid(auctionId uint64, amount sdk.Coin, bidTime time.Time, bidder sdk.AccAddress) NftAuctionBid {
	return NftAuctionBid{
		AuctionId: auctionId,
		Amount:    amount,
		Time:      bidTime,
		Bidder:    bidder.String(),
	}
}

func (b NftAuctionBid) GetAuctionId() uint64 {
	return b.AuctionId
}

func (b NftAuctionBid) GetAmount() sdk.Coin {
	return b.Amount
}

func (b NftAuctionBid) GetBidder() sdk.AccAddress {
	bidder, _ := sdk.AccAddressFromBech32(b.Bidder)
	return bidder
}
