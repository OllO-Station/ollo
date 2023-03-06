package config

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethermint "github.com/evmos/ethermint/types"
)

const (
	//
	DisplayDenom = "ollo"
	BaseDenom    = "uollo"

	//
	WiseDisplayDenom = "wise"
	WiseBaseDenom    = "uwise"
)

func RegisterDenoms() {
	if err := sdk.RegisterDenom(DisplayDenom, sdk.OneDec()); err != nil {
		panic(err)
	}
	if err := sdk.RegisterDenom(BaseDenom, sdk.NewDecWithPrec(1, ethermint.BaseDenomUnit)); err != nil {
		panic(err)
	}

	if err := sdk.RegisterDenom(WiseDisplayDenom, sdk.OneDec()); err != nil {
		panic(err)
	}
	if err := sdk.RegisterDenom(WiseBaseDenom, sdk.NewDecWithPrec(1, ethermint.BaseDenomUnit)); err != nil {
		panic(err)
	}
}

func SetBip44CoinType(config *sdk.Config) {
	config.SetCoinType(ethermint.Bip44CoinType)
	config.SetPurpose(sdk.Purpose)
	config.SetFullFundraiserPath(ethermint.BIP44HDPath)
}
