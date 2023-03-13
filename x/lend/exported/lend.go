package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type LoanI interface {
	GetId() uint64
	GetDenom() string
	GetAmount() sdk.Coin
	GetDeadline() time.Time
	GetPrincipal() sdk.Coin
}
