package keeper

import (
	"github.com/ollo-station/ollo/x/market/types"
)

var _ types.QueryServer = Keeper{}
