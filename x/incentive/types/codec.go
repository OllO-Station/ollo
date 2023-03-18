package types

import (
	cdc "github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocdc "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	msgservice "github.com/cosmos/cosmos-sdk/types/msgservice"
	// govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var (
	amino     = cdc.NewLegacyAmino()
	ModuleCdc = cdc.NewAminoCodec(amino)
)

func RegisterLegacyAminoCodec(c *cdc.LegacyAmino) {
	c.RegisterConcrete(&MsgApplyMarketMaker{}, "incentive/MsgApplyMarketMaker", nil)
	c.RegisterConcrete(&MsgClaimIncentive{}, "incentive/MsgClaimIncentive", nil)
	c.RegisterConcrete(&MarketMakerProposal{}, "incentive/MarketMakerProposal", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgApplyMarketMaker{},
		&MsgClaimIncentive{},
	)

	// registry.RegisterImplementations(
	// 	(*govtypes.Content)(nil),
	// 	&MarketMakerProposal{},
	// )

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocdc.RegisterCrypto(amino)
	amino.Seal()
}
