package keeper

import (
	"github.com/ollo-station/ollo/x/emissions/types"
)

var _ types.QueryServer = Keeper{}
