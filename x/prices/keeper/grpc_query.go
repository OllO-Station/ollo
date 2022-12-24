package keeper

import (
	"ollo/x/prices/types"
)

var _ types.QueryServer = Keeper{}
