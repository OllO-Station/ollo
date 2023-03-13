package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LockI interface {
	GetLockID() string
	GetOwner() sdk.AccAddress
	GetDenom() string
	GetAmount() sdk.Int
}
