package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"ollo/x/claim/types"

	liquiditykeeper "ollo/x/liquidity/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper      types.BankKeeper
		distrKeeper     types.DistrKeeper
		govKeeper       types.GovKeeper
		liquidityKeeper liquiditykeeper.Keeper
		stakingKeeper   types.StakingKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bk types.BankKeeper,
	dk types.DistrKeeper,
	gk types.GovKeeper,
	lk liquiditykeeper.Keeper,
	sk types.StakingKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:             cdc,
		storeKey:        storeKey,
		paramstore:      ps,
		distrKeeper:     dk,
		bankKeeper:      bk,
		govKeeper:       gk,
		liquidityKeeper: lk,
		stakingKeeper:   sk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
