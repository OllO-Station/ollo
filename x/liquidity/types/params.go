package types

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	farmingtypes "ollo/x/farming/types"
)

const (
	// CancelOrderLifeSpan is the lifespan of order cancellation.
	CancelOrderLifeSpan int64 = 0

	// MinReserveCoinNum is the minimum number of reserve coins in each liquidity pool.
	MinReserveCoinNum uint32 = 2

	// MaxReserveCoinNum is the maximum number of reserve coins in each liquidity pool.
	MaxReserveCoinNum uint32 = 2

	// DefaultUnitBatchHeight is the default number of blocks in one batch. This param is used for scalability.
	DefaultUnitBatchHeight uint32 = 1

	// DefaultPoolTypeID is the default pool type id. The only supported pool type id is 1.
	DefaultPoolTypeID uint32 = 1

	// DefaultSwapTypeID is the default swap type id. The only supported swap type (instant swap) id is 1.
	DefaultSwapTypeID uint32 = 1

	// DefaultCircuitBreakerEnabled is the default circuit breaker status. This param is used for a contingency plan.
	DefaultCircuitBreakerEnabled = false
)

// Parameter store keys
var (
	KeyPoolTypes              = []byte("PoolTypes")
	KeyMinInitDepositAmount   = []byte("MinInitDepositAmount")
	KeyInitPoolCoinMintAmount = []byte("InitPoolCoinMintAmount")
	KeyMaxReserveCoinAmount   = []byte("MaxReserveCoinAmount")
	KeySwapFeeRate            = []byte("SwapFeeRate")
	KeyPoolCreationFee        = []byte("PoolCreationFee")
	KeyUnitBatchHeight        = []byte("UnitBatchHeight")
	KeyWithdrawFeeRate        = []byte("WithdrawFeeRate")
	KeyMaxOrderAmountRatio    = []byte("MaxOrderAmountRatio")
	KeyCircuitBreakerEnabled  = []byte("CircuitBreakerEnabled")
)

var (
	DefaultMinInitDepositAmount   = sdk.NewInt(1000000)
	DefaultInitPoolCoinMintAmount = sdk.NewInt(1000000)
	DefaultMaxReserveCoinAmount   = sdk.ZeroInt()
	DefaultSwapFeeRate            = sdk.NewDecWithPrec(3, 3) // "0.003000000000000000"
	DefaultWithdrawFeeRate        = sdk.ZeroDec()
	DefaultMaxOrderAmountRatio    = sdk.NewDecWithPrec(1, 1) // "0.100000000000000000"
	DefaultPoolCreationFee        = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(40000000)))
	DefaultPoolType               = PoolType{
		Id:                1,
		Name:              "StandardLiquidityPool",
		MinReserveCoinNum: MinReserveCoinNum,
		MaxReserveCoinNum: MaxReserveCoinNum,
		Description:       "Standard liquidity pool with pool price function X/Y, ESPM constraint, and two kinds of reserve coins",
	}
	DefaultPoolTypes = []PoolType{DefaultPoolType}

	MinOfferCoinAmount = sdk.NewInt(100)
)

var _ paramstypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns the default liquidity module parameters.
func DefaultParams() Params {
	return Params{
		PoolTypes:              DefaultPoolTypes,
		MinInitDepositAmount:   DefaultMinInitDepositAmount,
		InitPoolCoinMintAmount: DefaultInitPoolCoinMintAmount,
		MaxReserveCoinAmount:   DefaultMaxReserveCoinAmount,
		PoolCreationFee:        DefaultPoolCreationFee,
		SwapFeeRate:            DefaultSwapFeeRate,
		WithdrawFeeRate:        DefaultWithdrawFeeRate,
		MaxOrderAmountRatio:    DefaultMaxOrderAmountRatio,
		UnitBatchHeight:        DefaultUnitBatchHeight,
		CircuitBreakerEnabled:  DefaultCircuitBreakerEnabled,
	}
}

// ParamSetPairs implements paramstypes.ParamSet.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyPoolTypes, &p.PoolTypes, validatePoolTypes),
		paramstypes.NewParamSetPair(KeyMinInitDepositAmount, &p.MinInitDepositAmount, validateMinInitDepositAmount),
		paramstypes.NewParamSetPair(KeyInitPoolCoinMintAmount, &p.InitPoolCoinMintAmount, validateInitPoolCoinMintAmount),
		paramstypes.NewParamSetPair(KeyMaxReserveCoinAmount, &p.MaxReserveCoinAmount, validateMaxReserveCoinAmount),
		paramstypes.NewParamSetPair(KeyPoolCreationFee, &p.PoolCreationFee, validatePoolCreationFee),
		paramstypes.NewParamSetPair(KeySwapFeeRate, &p.SwapFeeRate, validateSwapFeeRate),
		paramstypes.NewParamSetPair(KeyWithdrawFeeRate, &p.WithdrawFeeRate, validateWithdrawFeeRate),
		paramstypes.NewParamSetPair(KeyMaxOrderAmountRatio, &p.MaxOrderAmountRatio, validateMaxOrderAmountRatio),
		paramstypes.NewParamSetPair(KeyUnitBatchHeight, &p.UnitBatchHeight, validateUnitBatchHeight),
		paramstypes.NewParamSetPair(KeyCircuitBreakerEnabled, &p.CircuitBreakerEnabled, validateCircuitBreakerEnabled),
	}
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Validate validates parameters.
func (p Params) Validate() error {
	for _, v := range []struct {
		value     interface{}
		validator func(interface{}) error
	}{
		{p.PoolTypes, validatePoolTypes},
		{p.MinInitDepositAmount, validateMinInitDepositAmount},
		{p.InitPoolCoinMintAmount, validateInitPoolCoinMintAmount},
		{p.MaxReserveCoinAmount, validateMaxReserveCoinAmount},
		{p.PoolCreationFee, validatePoolCreationFee},
		{p.SwapFeeRate, validateSwapFeeRate},
		{p.WithdrawFeeRate, validateWithdrawFeeRate},
		{p.MaxOrderAmountRatio, validateMaxOrderAmountRatio},
		{p.UnitBatchHeight, validateUnitBatchHeight},
		{p.CircuitBreakerEnabled, validateCircuitBreakerEnabled},
	} {
		if err := v.validator(v.value); err != nil {
			return err
		}
	}
	return nil
}

func validatePoolTypes(i interface{}) error {
	v, ok := i.([]PoolType)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("pool types must not be empty")
	}

	for i, p := range v {
		if int(p.Id) != i+1 {
			return fmt.Errorf("pool type ids must be sorted")
		}
		if p.MaxReserveCoinNum > MaxReserveCoinNum || MinReserveCoinNum > p.MinReserveCoinNum {
			return fmt.Errorf("min, max reserve coin num value of pool types are out of bounds")
		}
	}

	if len(v) > 1 || !v[0].Equal(DefaultPoolType) {
		return fmt.Errorf("the only supported pool type is 1")
	}

	return nil
}

//nolint:staticcheck,nolintlint
func validateMinInitDepositAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("minimum initial deposit amount must not be nil")
	}

	if !v.IsPositive() {
		return fmt.Errorf("minimum initial deposit amount must be positive: %s", v)
	}

	return nil
}

//nolint:staticcheck,nolintlint
func validateInitPoolCoinMintAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("initial pool coin mint amount must not be nil")
	}

	if !v.IsPositive() {
		return fmt.Errorf("initial pool coin mint amount must be positive: %s", v)
	}

	if v.LT(DefaultInitPoolCoinMintAmount) {
		return fmt.Errorf("initial pool coin mint amount must be greater than or equal to 1000000: %s", v)
	}

	return nil
}

//nolint:staticcheck,nolintlint
func validateMaxReserveCoinAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("max reserve coin amount must not be nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("max reserve coin amount must not be negative: %s", v)
	}

	return nil
}

func validateSwapFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec) //nolint:nolintlint
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("swap fee rate must not be nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("swap fee rate must not be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("swap fee rate too large: %s", v)
	}

	return nil
}

func validateWithdrawFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec) //nolint:nolintlint
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("withdraw fee rate must not be nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("withdraw fee rate must not be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("withdraw fee rate too large: %s", v)
	}

	return nil
}

func validateMaxOrderAmountRatio(i interface{}) error {
	v, ok := i.(sdk.Dec) //nolint:nolintlint
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("max order amount ratio must not be nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("max order amount ratio must not be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("max order amount ratio too large: %s", v)
	}

	return nil
}

func validatePoolCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return err
	}

	if v.Empty() {
		return fmt.Errorf("pool creation fee must not be empty")
	}

	return nil
}

func validateUnitBatchHeight(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("unit batch height must be positive: %d", v)
	}

	return nil
}

func validateCircuitBreakerEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// Liquidity params default values
const (
	DefaultBatchSize                    uint32 = 1
	DefaultTickPrecision                uint32 = 4
	DefaultMaxNumMarketMakingOrderTicks        = 10
	DefaultMaxOrderLifespan                    = 24 * time.Hour
	DefaultMaxNumActivePoolsPerPair            = 20
)

// Liquidity params default values
var (
	DefaultFeeCollectorAddress      = farmingtypes.DeriveAddress(AddressType, ModuleName, "FeeCollector")
	DefaultDustCollectorAddress     = farmingtypes.DeriveAddress(AddressType, ModuleName, "DustCollector")
	DefaultMinInitialPoolCoinSupply = sdk.NewInt(1_000_000_000_000)
	DefaultPairCreationFee          = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000))
	// DefaultPoolCreationFee          = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000))
	DefaultMinInitialDepositAmount = sdk.NewInt(1000000)
	DefaultMaxPriceLimitRatio      = sdk.NewDecWithPrec(1, 1) // 10%
	// DefaultSwapFeeRate              = sdk.ZeroDec()
	// DefaultWithdrawFeeRate          = sdk.ZeroDec()
	DefaultDepositExtraGas  = sdk.Gas(60000)
	DefaultWithdrawExtraGas = sdk.Gas(64000)
	DefaultOrderExtraGas    = sdk.Gas(37000)
)

// General constants
const (
	PoolReserveAddressPrefix  = "PoolReserveAddress"
	PairEscrowAddressPrefix   = "PairEscrowAddress"
	ModuleAddressNameSplitter = "|"
	AddressType               = farmingtypes.AddressType32Bytes
)

// GlobalEscrowAddress is an escrow for deposit/withdraw requests.
var GlobalEscrowAddress = farmingtypes.DeriveAddress(AddressType, ModuleName, "GlobalEscrow")

var (
	KeyBatchSize                = []byte("BatchSize")
	KeyTickPrecision            = []byte("TickPrecision")
	KeyFeeCollectorAddress      = []byte("FeeCollectorAddress")
	KeyDustCollectorAddress     = []byte("DustCollectorAddress")
	KeyMinInitialPoolCoinSupply = []byte("MinInitialPoolCoinSupply")
	KeyPairCreationFee          = []byte("PairCreationFee")
	// KeyPoolCreationFee              = []byte("PoolCreationFee")
	KeyMinInitialDepositAmount      = []byte("MinInitialDepositAmount")
	KeyMaxPriceLimitRatio           = []byte("MaxPriceLimitRatio")
	KeyMaxNumMarketMakingOrderTicks = []byte("MaxNumMarketMakingOrderTicks")
	KeyMaxOrderLifespan             = []byte("MaxOrderLifespan")
	// KeySwapFeeRate                  = []byte("SwapFeeRate")
	// KeyWithdrawFeeRate              = []byte("WithdrawFeeRate")
	KeyDepositExtraGas          = []byte("DepositExtraGas")
	KeyWithdrawExtraGas         = []byte("WithdrawExtraGas")
	KeyOrderExtraGas            = []byte("OrderExtraGas")
	KeyMaxNumActivePoolsPerPair = []byte("MaxNumActivePoolsPerPair")
)

var _ paramstypes.ParamSet = (*Params)(nil)

// func ParamKeyTable() paramstypes.KeyTable {
// 	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
// }
//
// // DefaultParams returns a default params for the liquidity module.
// func DefaultParams() Params {
// 	return Params{
// 		BatchSize:                    DefaultBatchSize,
// 		TickPrecision:                DefaultTickPrecision,
// 		FeeCollectorAddress:          DefaultFeeCollectorAddress.String(),
// 		DustCollectorAddress:         DefaultDustCollectorAddress.String(),
// 		MinInitialPoolCoinSupply:     DefaultMinInitialPoolCoinSupply,
// 		PairCreationFee:              DefaultPairCreationFee,
// 		PoolCreationFee:              DefaultPoolCreationFee,
// 		MinInitialDepositAmount:      DefaultMinInitialDepositAmount,
// 		MaxPriceLimitRatio:           DefaultMaxPriceLimitRatio,
// 		MaxNumMarketMakingOrderTicks: DefaultMaxNumMarketMakingOrderTicks,
// 		MaxOrderLifespan:             DefaultMaxOrderLifespan,
// 		SwapFeeRate:                  DefaultSwapFeeRate,
// 		WithdrawFeeRate:              DefaultWithdrawFeeRate,
// 		DepositExtraGas:              DefaultDepositExtraGas,
// 		WithdrawExtraGas:             DefaultWithdrawExtraGas,
// 		OrderExtraGas:                DefaultOrderExtraGas,
// 		MaxNumActivePoolsPerPair:     DefaultMaxNumActivePoolsPerPair,
// 	}
// }
//
// // ParamSetPairs implements ParamSet.
// func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
// 	return paramstypes.ParamSetPairs{
// 		paramstypes.NewParamSetPair(KeyBatchSize, &params.BatchSize, validateBatchSize),
// 		paramstypes.NewParamSetPair(KeyTickPrecision, &params.TickPrecision, validateTickPrecision),
// 		paramstypes.NewParamSetPair(KeyFeeCollectorAddress, &params.FeeCollectorAddress, validateFeeCollectorAddress),
// 		paramstypes.NewParamSetPair(KeyDustCollectorAddress, &params.DustCollectorAddress, validateDustCollectorAddress),
// 		paramstypes.NewParamSetPair(KeyMinInitialPoolCoinSupply, &params.MinInitialPoolCoinSupply, validateMinInitialPoolCoinSupply),
// 		paramstypes.NewParamSetPair(KeyPairCreationFee, &params.PairCreationFee, validatePairCreationFee),
// 		paramstypes.NewParamSetPair(KeyPoolCreationFee, &params.PoolCreationFee, validatePoolCreationFee),
// 		paramstypes.NewParamSetPair(KeyMinInitialDepositAmount, &params.MinInitialDepositAmount, validateMinInitialDepositAmount),
// 		paramstypes.NewParamSetPair(KeyMaxPriceLimitRatio, &params.MaxPriceLimitRatio, validateMaxPriceLimitRatio),
// 		paramstypes.NewParamSetPair(KeyMaxNumMarketMakingOrderTicks, &params.MaxNumMarketMakingOrderTicks, validateMaxNumMarketMakingOrderTicks),
// 		paramstypes.NewParamSetPair(KeyMaxOrderLifespan, &params.MaxOrderLifespan, validateMaxOrderLifespan),
// 		paramstypes.NewParamSetPair(KeySwapFeeRate, &params.SwapFeeRate, validateSwapFeeRate),
// 		paramstypes.NewParamSetPair(KeyWithdrawFeeRate, &params.WithdrawFeeRate, validateWithdrawFeeRate),
// 		paramstypes.NewParamSetPair(KeyDepositExtraGas, &params.DepositExtraGas, validateExtraGas),
// 		paramstypes.NewParamSetPair(KeyWithdrawExtraGas, &params.WithdrawExtraGas, validateExtraGas),
// 		paramstypes.NewParamSetPair(KeyOrderExtraGas, &params.OrderExtraGas, validateExtraGas),
// 		paramstypes.NewParamSetPair(KeyMaxNumActivePoolsPerPair, &params.MaxNumActivePoolsPerPair, validateMaxNumActivePoolsPerPair),
// 	}
// }
//
// // Validate validates Params.
// func (params Params) Validate() error {
// 	for _, field := range []struct {
// 		val          interface{}
// 		validateFunc func(i interface{}) error
// 	}{
// 		{params.BatchSize, validateBatchSize},
// 		{params.TickPrecision, validateTickPrecision},
// 		{params.FeeCollectorAddress, validateFeeCollectorAddress},
// 		{params.DustCollectorAddress, validateDustCollectorAddress},
// 		{params.MinInitialPoolCoinSupply, validateMinInitialPoolCoinSupply},
// 		{params.PairCreationFee, validatePairCreationFee},
// 		{params.PoolCreationFee, validatePoolCreationFee},
// 		{params.MinInitialDepositAmount, validateMinInitialDepositAmount},
// 		{params.MaxPriceLimitRatio, validateMaxPriceLimitRatio},
// 		{params.MaxNumMarketMakingOrderTicks, validateMaxNumMarketMakingOrderTicks},
// 		{params.MaxOrderLifespan, validateMaxOrderLifespan},
// 		{params.SwapFeeRate, validateSwapFeeRate},
// 		{params.WithdrawFeeRate, validateWithdrawFeeRate},
// 		{params.DepositExtraGas, validateExtraGas},
// 		{params.WithdrawExtraGas, validateExtraGas},
// 		{params.OrderExtraGas, validateExtraGas},
// 		{params.MaxNumActivePoolsPerPair, validateMaxNumActivePoolsPerPair},
// 	} {
// 		if err := field.validateFunc(field.val); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func validateBatchSize(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("batch size must be positive: %d", v)
	}

	return nil
}

func validateTickPrecision(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateFeeCollectorAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if _, err := sdk.AccAddressFromBech32(v); err != nil {
		return fmt.Errorf("invalid fee collector address: %w", err)
	}

	return nil
}

func validateDustCollectorAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if _, err := sdk.AccAddressFromBech32(v); err != nil {
		return fmt.Errorf("invalid dust collector address: %w", err)
	}

	return nil
}

func validateMinInitialPoolCoinSupply(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("min initial pool coin supply must not be nil")
	}

	if !v.IsPositive() {
		return fmt.Errorf("min initial pool coin supply must be positive: %s", v)
	}

	return nil
}

func validatePairCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return fmt.Errorf("invalid pair creation fee: %w", err)
	}

	return nil
}

//	func validatePoolCreationFee(i interface{}) error {
//		v, ok := i.(sdk.Coins)
//		if !ok {
//			return fmt.Errorf("invalid parameter type: %T", i)
//		}
//
//		if err := v.Validate(); err != nil {
//			return fmt.Errorf("invalid pool creation fee: %w", err)
//		}
//
//		return nil
//	}
func validateMinInitialDepositAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("minimum initial deposit amount must not be negative: %s", v)
	}

	return nil
}

func validateMaxPriceLimitRatio(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("max price limit ratio must not be negative: %s", v)
	}

	return nil
}

func validateMaxNumMarketMakingOrderTicks(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max number of market making order ticks must be positive: %d", v)
	}

	return nil
}

func validateMaxOrderLifespan(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("max order lifespan must not be negative: %s", v)
	}

	return nil
}

// func validateSwapFeeRate(i interface{}) error {
// 	v, ok := i.(sdk.Dec)
// 	if !ok {
// 		return fmt.Errorf("invalid parameter type: %T", i)
// 	}
//
// 	if v.IsNegative() {
// 		return fmt.Errorf("swap fee rate must not be negative: %s", v)
// 	}
//
// 	return nil
// }
//
// func validateWithdrawFeeRate(i interface{}) error {
// 	v, ok := i.(sdk.Dec)
// 	if !ok {
// 		return fmt.Errorf("invalid parameter type: %T", i)
// 	}
//
// 	if v.IsNegative() {
// 		return fmt.Errorf("withdraw fee rate must not be negative: %s", v)
// 	}
//
// 	return nil
// }

func validateExtraGas(i interface{}) error {
	_, ok := i.(sdk.Gas)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateMaxNumActivePoolsPerPair(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
