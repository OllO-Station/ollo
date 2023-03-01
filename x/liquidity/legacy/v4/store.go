package v4

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/ollo-station/ollo/x/liquidity/types"
)

func DeleteMMOrderIndexes(store sdk.KVStore) {
	iter := sdk.KVStorePrefixIterator(store, MMOrderIndexKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

func MigrateStore(
	ctx sdk.Context,
	storeKey storetypes.StoreKey,
	paramSpace paramstypes.Subspace,
) error {
	store := ctx.KVStore(storeKey)
	DeleteMMOrderIndexes(store)
	paramSpace.Set(
		ctx,
		types.KeyMaxNumMarketMakingOrdersPerPair,
		uint32(types.DefaultMaxNumMarketMakingOrdersPerPair),
	)
	return nil
}
