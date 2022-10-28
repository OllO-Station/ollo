package keeper

import (
	"ollo/x/reserve/types"
)

var _ types.QueryServer = Keeper{}
