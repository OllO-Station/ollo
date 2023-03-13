package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/ollo-station/ollo/x/market/exported"
)

var (
	_ proto.Message        = &NftListing{}
	_ exported.NftListingI = &NftListing{}
)

func NewNftListing(id, nftId, denomId string, price sdk.Coin, creator sdk.AccAddress) NftListing {
	return NftListing{
		Id:      id,
		NftId:   nftId,
		DenomId: denomId,
		Price:   price,
		Creator: creator.String(),
	}
}

func (l NftListing) GetId() string {
	return l.Id
}

func (l NftListing) GetDenomId() string {
	return l.DenomId
}

func (l NftListing) GetNftId() string {
	return l.NftId
}

func (l NftListing) GetPrice() sdk.Coin {
	return l.Price
}

func (l NftListing) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(l.Creator)
	return owner
}
