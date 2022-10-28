package keeper

import (
	"ollo/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
