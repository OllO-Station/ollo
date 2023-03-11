package keeper

import (
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ollo-station/ollo/x/epoch/types"
	log "github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	hooks    types.EpochHook
}

func NewKeeper(storeKey storetypes.StoreKey) *Keeper {
	return &Keeper{
		storeKey: storeKey,
	}
}

//	func (k *Keeper) HasHook() bool {
//		return len(k.hooks) > 0
//	}
func (k *Keeper) SetHooks(h []types.EpochHook) *Keeper {
	// k.hooks = h
	return k
}
func (k *Keeper) AddHook(h types.EpochHook) *Keeper {
	// k.hooks = append(k.hooks, h)
	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// func (k Keeper) GetEpoch(ctx sdk.Context) (epoch uint64, err error) {
// 	store := ctx.KVStore(k.storeKey)
// 	fmt.Println("store", store)
// 	// bz := store.Get(types.EpochHook)
// 	// if bz == nil {
// 	// 	return 0, fmt.Errorf("epoch not found")
// 	// }
// 	// k.cdc.MustUnmarshalBinaryBare(bz, &epoch)
// 	return epoch, nil
// }
