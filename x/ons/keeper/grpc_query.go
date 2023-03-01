package keeper

import (
	"github.com/ollo-station/ollo/x/ons/types"
)

var _ types.QueryServer = Keeper{}
