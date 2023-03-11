package keeper

import (
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	// errorsignite "errors"
	"github.com/ollo-station/ollo/x/mint/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         storetypes.StoreKey
	paramSpace       paramtypes.Subspace
	stakingKeeper    types.StakingKeeper
	accountKeeper    types.AccountKeeper
	bankKeeper       types.BankKeeper
	distrKeeper      types.DistrKeeper
	epochKeeper      types.EpochKeeper
	feeCollectorName string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, paramSpace paramtypes.Subspace,
	sk types.StakingKeeper, ak types.AccountKeeper, bk types.BankKeeper, dk types.DistrKeeper, ek types.EpochKeeper,
	feeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:              cdc,
		storeKey:         key,
		paramSpace:       paramSpace,
		stakingKeeper:    sk,
		accountKeeper:    ak,
		bankKeeper:       bk,
		distrKeeper:      dk,
		epochKeeper:      ek,
		feeCollectorName: feeCollectorName,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// GetMinter gets the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

// SetMinter sets the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.MinterKey, b)
}

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// StakingTokenSupply to be used in BeginBlocker.
func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdkmath.Int {
	return k.stakingKeeper.StakingTokenSupply(ctx)
}

// BondedRatio implements an alias call to the underlying staking keeper's
// BondedRatio to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	return k.stakingKeeper.BondedRatio(ctx)
}

// MintCoin implements an alias call to the underlying supply keeper's
// MintCoin to be used in BeginBlocker.
func (k Keeper) MintCoin(ctx sdk.Context, coin sdk.Coin) error {
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
}

// GetProportion gets the balance of the `MintedDenom` from minted coins and returns coins according to the `AllocationRatio`.
func (k Keeper) GetProportion(ctx sdk.Context, mintedCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(mintedCoin.Denom, sdk.NewDecFromInt(mintedCoin.Amount).Mul(ratio).TruncateInt())
}

// DistributeMintedCoin implements distribution of minted coins from mint
// to be used in BeginBlocker.
func (k Keeper) DistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin) error {
	params := k.GetParams(ctx)
	proportions := params.DistributionProportions

	// allocate staking rewards into fee collector account to be moved to on next begin blocker by staking module
	stakingRewardsCoins := sdk.NewCoins(k.GetProportion(ctx, mintedCoin, proportions.Staking))
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, stakingRewardsCoins)
	if err != nil {
		return err
	}

	fundedAddrsCoin := k.GetProportion(ctx, mintedCoin, proportions.FundedAddresses)
	fundedAddrsCoins := sdk.NewCoins(fundedAddrsCoin)
	if len(params.FundedAddresses) == 0 {
		// fund community pool when rewards address is empty
		if err = k.distrKeeper.FundCommunityPool(
			ctx,
			fundedAddrsCoins,
			k.accountKeeper.GetModuleAddress(types.ModuleName),
		); err != nil {
			return err
		}
	} else {
		// allocate developer rewards to developer addresses by weight
		for _, w := range params.FundedAddresses {
			fundedAddrCoins := sdk.NewCoins(k.GetProportion(ctx, fundedAddrsCoin, w.Weight))
			devAddr, err := sdk.AccAddressFromBech32(w.Address)
			if err != nil {
				// return errorsignite.Critical(err.Error())
				return err
			}
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, devAddr, fundedAddrCoins)
			if err != nil {
				return err
			}
		}
	}

	// subtract from original provision to ensure no coins left over after the allocations
	communityPoolCoins := sdk.NewCoins(mintedCoin).Sub(stakingRewardsCoins...).Sub(fundedAddrsCoins...)
	err = k.distrKeeper.FundCommunityPool(ctx, communityPoolCoins, k.accountKeeper.GetModuleAddress(types.ModuleName))
	if err != nil {
		return err
	}

	return err
}

func (k Keeper) DistributeToModule(ctx sdk.Context, recvModule string, coin sdk.Coin, prop sdk.Dec) (sdk.Int, error) {
	distrCoin := k.GetProportion(ctx, coin, prop)
	if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, recvModule, sdk.NewCoins(distrCoin)); err != nil {
		return sdk.Int{}, err
	}
	return distrCoin.Amount, nil
}

func getProportions(coin sdk.Coin, prop sdk.Dec) (sdk.Coin, error) {
	if prop.GT(sdk.OneDec()) {
		return sdk.Coin{}, types.InvalidProportionError{Proportion: prop}
	}
	return sdk.NewCoin(coin.Denom, sdk.NewDecFromInt(coin.Amount).Mul(prop).TruncateInt()), nil
}

func (k Keeper) GetEpochLastReduction(c sdk.Context) uint64 {
	s := c.KVStore(k.storeKey)
	b := s.Get([]byte(types.LastEpochReductionKey))
	if b == nil {
		return 0
	}
	return sdk.BigEndianToUint64(b)
}
func (k Keeper) SetEpochLastReduction(c sdk.Context, num uint64) {
	s := c.KVStore(k.storeKey)
	s.Set([]byte(types.LastEpochReductionKey), sdk.Uint64ToBigEndian(num))
}
