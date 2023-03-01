package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	proto "github.com/gogo/protobuf/proto"

	"github.com/ollo-station/ollo/x/nft/exported"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// RegisterLegacyAminoCodec concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgIssueDenom{}, "ollo/nft/v1/MsgIssueDenom", nil)
	cdc.RegisterConcrete(&MsgSendNFT{}, "ollo/nft/v1/MsgSendNFT", nil)
	cdc.RegisterConcrete(&MsgEditNFT{}, "ollo/nft/v1/MsgEditNFT", nil)
	cdc.RegisterConcrete(&MsgMintNFT{}, "ollo/nft/v1/MsgMintNFT", nil)
	cdc.RegisterConcrete(&MsgBurnNFT{}, "ollo/nft/v1/MsgBurnNFT", nil)
	cdc.RegisterConcrete(&MsgTransferDenom{}, "ollo/nft/v1/MsgTransferDenom", nil)

	cdc.RegisterInterface((*exported.NFT)(nil), nil)
	cdc.RegisterConcrete(&BaseNFT{}, "ollo/nft/v1/BaseNFT", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgIssueDenom{},
		&MsgSendNFT{},
		&MsgEditNFT{},
		&MsgMintNFT{},
		&MsgBurnNFT{},
		&MsgTransferDenom{},
	)

	registry.RegisterImplementations(
		(*exported.NFT)(nil),
		&BaseNFT{},
	)

	registry.RegisterImplementations(
		(*proto.Message)(nil),
		&DenomMetadata{},
		&NFTMetadata{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
