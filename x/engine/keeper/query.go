package keeper

import (
	"github.com/ollo-station/ollo/x/engine/types"
)

var _ types.QueryServer = Keeper{}
