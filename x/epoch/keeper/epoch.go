package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/ollo-station/ollo/x/epoch/types"
)

func (k Keeper) GetEpoch(ctx sdk.Context, id string) types.Epoch {
	epoch := types.Epoch{}
	store := ctx.KVStore(k.storeKey)
	b := store.Get(append(types.EpochPrefix, []byte(id)...))
	if b == nil {
		return epoch
	}
	err := proto.Unmarshal(b, &epoch)
	if err != nil {
		panic(err)
	}
	return epoch
}
func (k Keeper) AllEpochs(ctx sdk.Context) []types.Epoch {
	epochs := []types.Epoch{}
	k.IterateEpochs(ctx, func(i int64, epoch types.Epoch) (stop bool) {
		epochs = append(epochs, epoch)
		return false
	})
	return epochs
}

func (k Keeper) DeleteEpoch(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(types.EpochPrefix, []byte(id)...))
}

func (k Keeper) AddEpoch(ctx sdk.Context, epoch types.Epoch) error {
	e := epoch.Validate()
	if e != nil {
		return e
	}
	if (k.GetEpoch(ctx, epoch.Id) != types.Epoch{}) {
		return fmt.Errorf("epoch with id %s already exists", epoch.Id)
	}
	if epoch.Start.Equal(time.Time{}) {
		epoch.Start = ctx.BlockTime()
	}
	epoch.CurrentEpochStartHeight = uint64(ctx.BlockHeight())
	k.SetEpoch(ctx, epoch)
	return nil
}

func (k Keeper) SetEpoch(ctx sdk.Context, epoch types.Epoch) {
	store := ctx.KVStore(k.storeKey)
	v, err := proto.Marshal(&epoch)
	if err != nil {
		panic(err)
	}
	store.Set(append(types.EpochPrefix, []byte(epoch.Id)...), v)
}

func (k Keeper) BlocksSinceEpochStart(ctx sdk.Context, id string) (uint64, error) {
	epoch := k.GetEpoch(ctx, id)
	if (epoch == types.Epoch{}) {
		return 0, fmt.Errorf("epoch with id %s does not exist", id)
	}
	return uint64(ctx.BlockHeight()) - epoch.CurrentEpochStartHeight, nil
}

func (k Keeper) IterateEpochs(ctx sdk.Context, fn func(i int64, epoch types.Epoch) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.EpochPrefix)
	defer iterator.Close()
	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		println(string(iterator.Key()))
		epochI := types.Epoch{}
		err := proto.Unmarshal(iterator.Value(), &epochI)
		if err != nil {
			panic(err)
		}
		stop := fn(i, epochI)
		if stop {
			break
		}
		i++
	}
}
