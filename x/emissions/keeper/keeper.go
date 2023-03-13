package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ollo-station/ollo/x/emissions/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper      types.BankKeeper
		distrKeeper     types.DistrKeeper
		accountKeeper   types.AccountKeeper
		stakingKeeper   types.StakingKeeper
		epochingKeeper  types.EpochingKeeper
		mintKeeper      types.MintKeeper
		govKeeper       types.GovKeeper
		liquidityKeeper types.LiquidityKeeper
		lendKeeper      types.LendKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
	distrKeeper types.DistrKeeper,
	accountKeeper types.AccountKeeper,
	stakingKeeper types.StakingKeeper,
	epochingKeeper types.EpochingKeeper,
	mintKeeper types.MintKeeper,
	govKeeper types.GovKeeper,
	liquidityKeeper types.LiquidityKeeper,
	lendKeeper types.LendKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		bankKeeper:      bankKeeper,
		distrKeeper:     distrKeeper,
		accountKeeper:   accountKeeper,
		stakingKeeper:   stakingKeeper,
		epochingKeeper:  epochingKeeper,
		mintKeeper:      mintKeeper,
		govKeeper:       govKeeper,
		liquidityKeeper: liquidityKeeper,
		lendKeeper:      lendKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
