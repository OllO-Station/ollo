package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramstypes.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Liquidity params default values
const (
	DefaultBatchSize                    uint32 = 1
	DefaultTickPrecision                uint32 = 4
	DefaultMaxNumMarketMakingOrderTicks        = 10
	DefaultMaxOrderLifespan                    = 24 * time.Hour
)

// Liquidity params default values
var (
	DefaultFeeCollectorAddress      = DeriveAddress(LiquidityAddressType, ModuleName, "FeeCollector")
	DefaultDustCollectorAddress     = DeriveAddress(LiquidityAddressType, ModuleName, "DustCollector")
	DefaultMinInitialPoolCoinSupply = sdk.NewInt(1_000_000_000_000)
	DefaultPairCreationFee          = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000))
	DefaultPoolCreationFee          = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000))
	DefaultMinInitialDepositAmount  = sdk.NewInt(1000000)
	DefaultMaxPriceLimitRatio       = sdk.NewDecWithPrec(1, 1) // 10%
	DefaultSwapFeeRate              = sdk.ZeroDec()
	DefaultWithdrawFeeRate          = sdk.ZeroDec()
	DefaultDepositExtraGas          = sdk.Gas(60000)
	DefaultWithdrawExtraGas         = sdk.Gas(64000)
	DefaultOrderExtraGas            = sdk.Gas(37000)
)

// General constants
const (
	PoolReserveAddressPrefix  = "PoolReserveAddress"
	PairEscrowAddressPrefix   = "PairEscrowAddress"
	ModuleAddressNameSplitter = "|"
	LiquidityAddressType      = AddressType32Bytes

	// MaxNumActivePoolsPerPair is the maximum number of active(not disabled)
	// pools per pair.
	MaxNumActivePoolsPerPair = 50
)

var (
	// GlobalEscrowAddress is an escrow for deposit/withdraw requests.
	GlobalEscrowAddress = DeriveAddress(LiquidityAddressType, ModuleName, "GlobalEscrow")
)

var (
	KeyBatchSize                    = []byte("BatchSize")
	KeyTickPrecision                = []byte("TickPrecision")
	KeyFeeCollectorAddress          = []byte("FeeCollectorAddress")
	KeyDustCollectorAddress         = []byte("DustCollectorAddress")
	KeyMinInitialPoolCoinSupply     = []byte("MinInitialPoolCoinSupply")
	KeyPairCreationFee              = []byte("PairCreationFee")
	KeyPoolCreationFee              = []byte("PoolCreationFee")
	KeyMinInitDepositAmount         = []byte("MinInitialDepositAmount")
	KeyMaxOrderAmountRatio          = []byte("MaxPriceLimitRatio")
	KeyMaxNumMarketMakingOrderTicks = []byte("MaxNumMarketMakingOrderTicks")
	KeyMaxOrderLifespan             = []byte("MaxOrderLifespan")
	KeySwapFeeRate                  = []byte("SwapFeeRate")
	KeyWithdrawFeeRate              = []byte("WithdrawFeeRate")
	KeyDepositExtraGas              = []byte("DepositExtraGas")
	KeyWithdrawExtraGas             = []byte("WithdrawExtraGas")
	KeyOrderExtraGas                = []byte("OrderExtraGas")
)

var _ paramstypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

var poolTypes = []PoolType{
	{
		Id:          1,
		Description: "Standard liquidity pool with pool price function X/Y, ESPM constraint, and two kinds of reserve coins",
		Name:        "StandardLiquidityPool",
	},
	{
		Id:          2,
		Description: "Liquidity pool capped at +/- 3% of price",
		Name:        "IntelligentLiquidityPool",
	},
}

// DefaultParams returns a default params for the liquidity module.
func DefaultParams() Params {
	return Params{
		PoolParams: &PoolParams{
			MinInitPoolSupply:  DefaultMinInitialPoolCoinSupply,
			MinInitPoolDeposit: DefaultMinInitialDepositAmount,
			MaxPriceRatio:      DefaultMaxPriceLimitRatio,
		},
		FeeParams: &FeeParams{
			FeeCollecterAddr:  DefaultFeeCollectorAddress.String(),
			DustCollectorAddr: DefaultDustCollectorAddress.String(),
			SwapRate:          &DefaultSwapFeeRate,
			WithdrawRate:      &DefaultWithdrawFeeRate,
			PairInitFee:       DefaultPairCreationFee,
			PoolInitFee:       DefaultPoolCreationFee,
		},
		MarketParams: &MarketParams{
			Precision:   DefaultTickPrecision,
			BatchSize:   DefaultBatchSize,
			MaxAmmTicks: DefaultMaxNumMarketMakingOrderTicks,
		},
		OrderParams: &OrderParams{
			MaxLifespan:      DefaultMaxOrderLifespan,
			DepositExtraGas:  DefaultDepositExtraGas,
			OrderExtraGas:    DefaultOrderExtraGas,
			WithdrawExtraGas: DefaultWithdrawExtraGas,
		},
	}
}

// ParamSetPairs implements ParamSet.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		// paramstypes.NewParamSetPair(KeyBatchSize, &params.MarketParams.BatchSize, validateBatchSize),
		// paramstypes.NewParamSetPair(KeyTickPrecision, &params.MarketParams.BatchSize, validateTickPrecision),
		// paramstypes.NewParamSetPair(KeyFeeCollectorAddress, &params.FeeParams.FeeCollecterAddr, validateFeeCollectorAddress),
		// paramstypes.NewParamSetPair(KeyDustCollectorAddress, &params.FeeParams.DustCollectorAddr, validateDustCollectorAddress),
		// paramstypes.NewParamSetPair(KeyMinInitialPoolCoinSupply, &params.PoolParams.MinInitPoolSupply, validateInitPoolCoinMintAmount),
		// paramstypes.NewParamSetPair(KeyPairCreationFee, &params.FeeParams.PairInitFee, validatePairCreationFee),
		// paramstypes.NewParamSetPair(KeyPoolCreationFee, &params.FeeParams.PoolInitFee, validatePoolCreationFee),
		// paramstypes.NewParamSetPair(KeyMinInitDepositAmount, &params.PoolParams.MinInitPoolDeposit, validateMinInitDepositAmount),
		// paramstypes.NewParamSetPair(KeyMaxOrderAmountRatio, &params.PoolParams.MaxPriceRatio, validateMaxOrderAmountRatio),
		// paramstypes.NewParamSetPair(KeyMaxNumMarketMakingOrderTicks, &params.MarketParams.MaxAmmTicks, validateMaxNumMarketMakingOrderTicks),
		// paramstypes.NewParamSetPair(KeyMaxOrderLifespan, &params.OrderParams.MaxLifespan, validateMaxOrderLifespan),
		// paramstypes.NewParamSetPair(KeySwapFeeRate, &params.FeeParams.SwapRate, validateSwapFeeRate),
		// paramstypes.NewParamSetPair(KeyWithdrawFeeRate, &params.FeeParams.WithdrawRate, validateWithdrawFeeRate),
		// paramstypes.NewParamSetPair(KeyDepositExtraGas, &params.OrderParams.DepositExtraGas, validateExtraGas),
		// paramstypes.NewParamSetPair(KeyWithdrawExtraGas, &params.OrderParams.WithdrawExtraGas, validateExtraGas),
		// paramstypes.NewParamSetPair(KeyOrderExtraGas, &params.OrderParams.OrderExtraGas, validateExtraGas),
	}
}

// Validate validates Params.
func (params Params) Validate() error {
	for _, field := range []struct {
		val          interface{}
		validateFunc func(i interface{}) error
	}{
		// {params.MarketParams.BatchSize, validateBatchSize},
		// {params.MarketParams.Precision, validateTickPrecision},
		// {params.FeeParams.FeeCollecterAddr, validateFeeCollectorAddress},
		// {params.FeeParams.DustCollectorAddr, validateDustCollectorAddress},
		// {params.PoolParams.MinInitPoolSupply, validateInitPoolCoinMintAmount},
		// {params.FeeParams.PairInitFee, validatePairCreationFee},
		// {params.FeeParams.PoolInitFee, validatePoolCreationFee},
		// {params.PoolParams.MinInitPoolDeposit, validateMinInitDepositAmount},
		// {params.PoolParams.MaxPriceRatio, validateMaxOrderAmountRatio},
		// {params.MarketParams.MaxAmmTicks, validateMaxNumMarketMakingOrderTicks},
		// {params.OrderParams.MaxLifespan, validateMaxOrderLifespan},
		// {params.FeeParams.SwapRate, validateSwapFeeRate},
		// {params.FeeParams.WithdrawRate, validateWithdrawFeeRate},
		// {params.OrderParams.DepositExtraGas, validateExtraGas},
		// {params.OrderParams.WithdrawExtraGas, validateExtraGas},
		// {params.OrderParams.OrderExtraGas, validateExtraGas},
	} {
		if err := field.validateFunc(field.val); err != nil {
			return err
		}
	}
	return nil
}

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

func validateInitPoolCoinMintAmount(i interface{}) error {
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

func validatePoolCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return fmt.Errorf("invalid pool creation fee: %w", err)
	}

	return nil
}

func validateMinInitDepositAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("minimum initial deposit amount must not be negative: %s", v)
	}

	return nil
}

func validateMaxOrderAmountRatio(i interface{}) error {
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

func validateSwapFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("swap fee rate must not be negative: %s", v)
	}

	return nil
}

func validateWithdrawFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("withdraw fee rate must not be negative: %s", v)
	}

	return nil
}

func validateExtraGas(i interface{}) error {
	_, ok := i.(sdk.Gas)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
