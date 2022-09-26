package keeper

import (
	"ollo/x/wasm/types"
)

var _ types.QueryServer = Keeper{}
