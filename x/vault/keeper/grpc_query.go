package keeper

import (
	"ollo/x/vault/types"
)

var _ types.QueryServer = Keeper{}
