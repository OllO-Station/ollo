package keeper

import (
	"github.com/ollo-station/ollo/x/vault/types"
)

var _ types.QueryServer = Keeper{}
