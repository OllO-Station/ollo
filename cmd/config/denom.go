package config

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	DisplayDenom = "ollo"
	BaseDenom    = "uollo"
)

// RegisterDenoms registers the base and display denominations to the SDK.
func RegisterDenoms() {
	if err := sdk.RegisterDenom(DisplayDenom, sdk.OneDec()); err != nil {
		panic(err)
	}
}
