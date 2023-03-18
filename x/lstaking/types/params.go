package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	farmingtypes "github.com/ollo-station/ollo/x/farming/types"
)

var (
	_                           paramstypes.ParamSet = (*Params)(nil)
	KeyLiquidBondDenom                               = []byte("liquid_bond_denom")
	KeyValidatorWhitelist                            = []byte("validator_whitelist")
	KeyUnstakeFeeRate                                = []byte("unstake_fee_rate")
	KeyMinLiquidStakeAmount                          = []byte("min_liquid_stake_amount")
	DefaultLiquidBondDenom                           = "uwise"
	DefaultUnstakeFeeRate                            = sdk.NewDecWithPrec(1, 3)
	DefaultMinLiquidStakeAmount                      = sdk.NewInt(1_000_000)
	RebalanceTrigger                                 = sdk.NewDecWithPrec(1, 3)
	RewardTrigger                                    = sdk.NewDecWithPrec(1, 3)
	LiquidStakeReserveAcc                            = farmingtypes.DeriveAddress(
		farmingtypes.AddressType32Bytes,
		ModuleName,
		"liquid_stake_reserve_acc",
	)
)

func ParamsKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func DefaultParams() Params {
	return Params{}
}

func (p Params) Validate() error {
	return nil
}

func (p Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{}
}
