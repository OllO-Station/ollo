package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type FeeI interface {
	GetFeeID() string
	GetReserveAddr() sdk.AccAddress
	GetDenom() string
	GetReserveAmount() sdk.Int
	GetFeeAmount() sdk.Int
}
