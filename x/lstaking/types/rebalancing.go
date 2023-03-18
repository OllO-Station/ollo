package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func WeightedDiv(
	input sdk.Int,
) (outputs []sdk.Int, crumb sdk.Int) {
	totalWeight := sdk.ZeroInt()
	if !totalWeight.IsPositive() {
		return []sdk.Int{}, sdk.ZeroInt()
	}
	totalOut := sdk.ZeroInt()
	// unitInput := sdk.NewDecFromInt(input).QuoTruncate(sdk.NewDecFromInt(totalWeight))
	return outputs, input.Sub(totalOut)
}

func CurrentWeightedDiv(
	input sdk.Dec,
	totalLiquidTokens sdk.Int,
	liquidTokenMap map[string]sdk.Int,
) (outputs []sdk.Dec, crumb sdk.Dec) {
	if !totalLiquidTokens.IsPositive() {
		return []sdk.Dec{}, sdk.ZeroDec()
	}
	totalOut := sdk.ZeroDec()
	// unitInp := input.QuoTruncate(sdk.NewDecFromInt(totalLiquidTokens))
	return outputs, input.Sub(totalOut)

}
