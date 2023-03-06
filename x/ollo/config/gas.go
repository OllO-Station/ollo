package config

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	//
	MainnetGasMinPrice = sdk.NewDec(20_000_000_000)
	//
	MainnetGasMinMultiplier = sdk.NewDecWithPrec(5, 1)
)
