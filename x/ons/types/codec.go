package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBuyName{}, "ons/BuyName", nil)
	cdc.RegisterConcrete(&MsgSellName{}, "ons/SellName", nil)
	cdc.RegisterConcrete(&MsgSetName{}, "ons/SetName", nil)
	cdc.RegisterConcrete(&MsgAddThread{}, "ons/AddThread", nil)
	cdc.RegisterConcrete(&MsgDeleteThread{}, "ons/DelThread", nil)
	cdc.RegisterConcrete(&MsgDeleteName{}, "ons/DelName", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSellName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetName{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddThread{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteThread{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
