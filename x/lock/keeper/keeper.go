package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

func NewKeeper(cdc codec.Codec, key storetypes.StoreKey, paramSpace paramstypes.Subspace, accountKeeper authkeeper.AccountKeeper, bankKeeper bankkeeper.Keeper) Keeper {
    return Keeper{
        cdc:         cdc,
        storeKey:    key,
        bankKeeper:  bankKeeper,
        paramSpace:  paramSpace,
        accountKeeper: accountKeeper,
    }
}

type Keeper struct {
    cdc         codec.Codec
    storeKey    storetypes.StoreKey
    bankKeeper  bankkeeper.Keeper
    paramSpace  paramstypes.Subspace
    accountKeeper authkeeper.AccountKeeper
}
