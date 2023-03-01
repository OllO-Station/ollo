package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterCodec(c *codec.LegacyAmino) {
	c.RegisterConcrete(&MsgCloseLockRequest{}, "ollo/lock/v1/close", nil)
	c.RegisterConcrete(&MsgCreateLockRequest{}, "ollo/lock/v1/create", nil)
	c.RegisterConcrete(&MsgDepositAssetRequest{}, "ollo/lock/v1/deposit", nil)
	c.RegisterConcrete(&MsgWithdrawAssetRequest{}, "ollo/lock/v1/withdraw", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCloseLockRequest{},
		&MsgCreateLockRequest{},
		&MsgDepositAssetRequest{},
		&MsgWithdrawAssetRequest{},
	)
	sdk.RegisterInterfaces(registry)
}
func init() {
	RegisterCodec(amino)
	sdk.RegisterLegacyAminoCodec(amino)
	RegisterCodec(authzcodec.Amino)
}
