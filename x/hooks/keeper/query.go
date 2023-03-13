package keeper

import (
	"github.com/ollo-station/ollo/x/hooks/types"
)

var _ types.QueryServer = Keeper{}
