package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/ollo-station/ollo/x/market/exported"
	"time"
)

var (
	_ proto.Message        = &NftAuction{}
	_ exported.NftAuctionI = &NftAuction{}
)

func NewNftAuction(id uint64, nftId, denomId string, startTime time.Time, endTime *time.Time, startPrice sdk.Coin,
	owner sdk.AccAddress, incrementPercentage sdk.Dec) NftAuction {
	return NftAuction{
		Id:                  id,
		NftId:               nftId,
		DenomId:             denomId,
		StartTime:           startTime,
		EndTime:             *endTime,
		StartPrice:          startPrice,
		IncrementPercentage: incrementPercentage,
		Owner:               owner.String(),
	}
}

func (al NftAuction) GetId() uint64 {
	return al.Id
}

func (al NftAuction) GetDenomId() string {
	return al.DenomId
}

func (al NftAuction) GetNftId() string {
	return al.NftId
}
func (al NftAuction) GetStartTime() time.Time {
	return al.StartTime
}

func (al NftAuction) GetStartPrice() sdk.Coin {
	return al.StartPrice
}

func (al NftAuction) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(al.Owner)
	return owner
}

func (al NftAuction) GetStatus() string {
	if al.StartTime.Before(time.Now()) {
		return NftAuctionStatusActive.String()
	}
	return NftAuctionStatusInactive.String()
}

func ValidAuctionStatus(status NftAuctionStatus) bool {
	if status == NftAuctionStatusInactive ||
		status == NftAuctionStatusActive {
		return true
	}
	return false
}
