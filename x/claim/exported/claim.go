package exported

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type EpochI interface {
	GetId() uint64
	GetStartTime() time.Time
	GetEndTime() time.Time
	GetReserveAddr() sdk.AccAddress
}
