package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	ZeroInt     = sdk.ZeroInt()
	ZeroDec     = sdk.ZeroDec()
	OneInt      = sdk.OneInt()
	OneDec      = sdk.OneDec()
	SmallestDec = sdk.SmallestDec()
)

func Key(comps ...[]byte) []byte {
	return bytes.Join(comps, nil)
}
