package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	types "github.com/ollo-station/ollo/x/lstaking/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper   types.AccountKeeper
	bankKeeper      types.BankKeeper
	distrKeeper     types.DistrKeeper
	stakingKeeper   types.StakingKeeper
	slashingKeeper  types.SlashingKeeper
	liquidityKeeper types.LiquidityKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	paramstore paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	distrKeeper types.DistrKeeper,
	stakingKeeper types.StakingKeeper,
	slashingKeeper types.SlashingKeeper,
	liquidityKeeper types.LiquidityKeeper,
) Keeper {
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types.ParamsKeyTable())
	}
	return Keeper{
		storeKey:        key,
		cdc:             cdc,
		paramstore:      paramstore,
		accountKeeper:   accountKeeper,
		bankKeeper:      bankKeeper,
		distrKeeper:     distrKeeper,
		stakingKeeper:   stakingKeeper,
		slashingKeeper:  slashingKeeper,
		liquidityKeeper: liquidityKeeper,
	}
}

func (k Keeper) Logger(c sdk.Context) log.Logger {
	return c.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the parameters for the liquidstaking module.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// GetCodec return codec.Codec object used by the keeper
func (k Keeper) GetCodec() codec.BinaryCodec { return k.cdc }
