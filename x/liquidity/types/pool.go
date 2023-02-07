package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
)

type PoolI interface {
	proto.Message

	GetAddr() sdk.AccAddress
	String() string
	GetId() uint64
	GetSwapFee(ctx sdk.Context) sdk.Dec
	GetExitFee(ctx sdk.Context) sdk.Dec
	IsActive(ctx sdk.Context) bool
	GetTotalShares() sdk.Dec
	GetTotalLiquidity(ctx sdk.Context) sdk.Coins
	GetKind() PoolKind
	SpotPrice(ctx sdk.Context, quoteDenom string, baseDenom string) (sdk.Dec, error)
}

type RangedPoolI interface {
	PoolI

	GetTokenFirst() string
	GetTokenSecond() string
	GetCurrentPrice() sdk.Dec
	GetCurrentTick() sdk.Int
	GetPrecFactor() sdk.Int
	GetTickSpacing() uint64
	GetLiquidity() sdk.Dec

	SetTick(tick sdk.Int)
	UpdateLiquidity(liquidity sdk.Dec)
	ApplySwap(newLiquidity sdk.Dec, newTick sdk.Int, newPrice sdk.Dec) error
	CalcAmounts(ctx sdk.Context, lowerTick, upperTick int64, sqrtRatioLowerTick, sqrtRatioUpperTick, liquidityDelta sdk.Dec) (amtDenomFirst sdk.Dec, amtDenomSecond sdk.Dec)
	UpdateLiquidityIfActivePosition(ctx sdk.Context, lowerTick, upperTick int64, liquidityDelta sdk.Dec) bool
}
