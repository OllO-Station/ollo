package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PlanI interface {
	GetId() uint64
	GetReserveAddr() sdk.AccAddress
}
