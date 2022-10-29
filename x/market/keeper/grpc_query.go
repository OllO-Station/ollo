package keeper

import (
	"ollo/x/market/types"
)

var _ types.QueryServer = Keeper{}
