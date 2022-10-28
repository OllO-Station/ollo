package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRequestLoan{}, "loan/RequestLoan", nil)
	cdc.RegisterConcrete(&MsgApproveLoan{}, "loan/ApproveLoan", nil)
	cdc.RegisterConcrete(&MsgRepayLoan{}, "loan/RepayLoan", nil)
	cdc.RegisterConcrete(&MsgLiquidateLoan{}, "loan/LiquidateLoan", nil)
	cdc.RegisterConcrete(&MsgCancelLoan{}, "loan/CancelLoan", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRepayLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLiquidateLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelLoan{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
