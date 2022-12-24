package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"ollo/x/grants/types"
	
)

type (
	Keeper struct {
		
		cdc      	codec.BinaryCodec
		storeKey 	storetypes.StoreKey
		memKey   	storetypes.StoreKey
		paramstore	paramtypes.Subspace
		
        accountKeeper types.AccountKeeper
        bankKeeper types.BankKeeper
        distrKeeper types.DistrKeeper
	}
)

func NewKeeper(
    cdc codec.BinaryCodec,
    storeKey,
    memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
    
    accountKeeper types.AccountKeeper,bankKeeper types.BankKeeper,distrKeeper types.DistrKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		
		cdc:      	cdc,
		storeKey: 	storeKey,
		memKey:   	memKey,
		paramstore:	ps,
		accountKeeper: accountKeeper,bankKeeper: bankKeeper,distrKeeper: distrKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
