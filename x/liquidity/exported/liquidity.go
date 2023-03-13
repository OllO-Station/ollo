package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PairI interface {
	GetId() uint64
	GetDenom() string
	GetQuoteDenom() string
}

type PoolI interface {
	GetId() uint64
	GetPairId() uint64
	GetPoolType() string
	GetReserveAddr() sdk.AccAddress
	GetReserveCoinDenom() string
	GetReserveCoin() sdk.Coin
}
