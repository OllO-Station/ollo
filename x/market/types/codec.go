package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/ollo-station/ollo/x/market/exported"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

var (
	Amino = codec.NewLegacyAmino()
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgListNft{}, "ollo/market/v1/MsgListNft", nil)
	cdc.RegisterConcrete(&MsgEditNftListing{}, "ollo/market/v1/MsgEditNftListing", nil)
	cdc.RegisterConcrete(&MsgDelistNft{}, "ollo/market/v1/MsgDelistNft", nil)
	cdc.RegisterConcrete(&MsgBuyNft{}, "ollo/market/v1/MsgBuyNft", nil)
	cdc.RegisterConcrete(&MsgCreateNftAuction{}, "ollo/market/v1/MsgCreateNftAuction", nil)
	cdc.RegisterConcrete(&MsgCancelNftAuction{}, "ollo/market/v1/MsgCancelNftAuction", nil)
	cdc.RegisterConcrete(&MsgPlaceNftBid{}, "ollo/market/v1/MsgPlaceNftBid", nil)

	cdc.RegisterInterface((*exported.NftListingI)(nil), nil)
	cdc.RegisterConcrete(&NftListing{}, "ollo/market/v1/Listing", nil)
	cdc.RegisterInterface((*exported.NftAuctionI)(nil), nil)
	cdc.RegisterConcrete(&NftAuction{}, "ollo/market/v1/AuctionListing", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgListNft{},
		&MsgEditNftListing{},
		&MsgDelistNft{},
		&MsgBuyNft{},
		&MsgCreateNftAuction{},
		&MsgCancelNftAuction{},
		&MsgPlaceNftBid{},
	)

	registry.RegisterImplementations((*exported.NftListingI)(nil),
		&NftListing{},
	)
	registry.RegisterImplementations((*exported.NftAuctionI)(nil),
		&NftAuction{},
	)
}

var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func MustMarshalListingID(cdc codec.BinaryCodec, listingId string) []byte {
	listingIdWrap := gogotypes.StringValue{Value: listingId}
	return cdc.MustMarshal(&listingIdWrap)
}

func MustUnMarshalListingID(cdc codec.BinaryCodec, value []byte) string {
	var listingIdWrap gogotypes.StringValue
	cdc.MustUnmarshal(value, &listingIdWrap)
	return listingIdWrap.Value
}
