package keeper

import (
	"github.com/ollo-station/ollo/x/lend/types"
)

var _ types.QueryServer = Keeper{}
