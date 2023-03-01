package keeper

import (
	"github.com/ollo-station/ollo/x/loan/types"
)

var _ types.QueryServer = Keeper{}
