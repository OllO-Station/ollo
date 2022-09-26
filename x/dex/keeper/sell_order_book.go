package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/dex/types"
)

// SetSellOrderBook set a specific sellOrderBook in the store from its index
func (k Keeper) SetSellOrderBook(ctx sdk.Context, sellOrderBook types.SellOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellOrderBookKeyPrefix))
	b := k.cdc.MustMarshal(&sellOrderBook)
	store.Set(types.SellOrderBookKey(
		sellOrderBook.Index,
	), b)
}

// GetSellOrderBook returns a sellOrderBook from its index
func (k Keeper) GetSellOrderBook(
	ctx sdk.Context,
	index string,

) (val types.SellOrderBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellOrderBookKeyPrefix))

	b := store.Get(types.SellOrderBookKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSellOrderBook removes a sellOrderBook from the store
func (k Keeper) RemoveSellOrderBook(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellOrderBookKeyPrefix))
	store.Delete(types.SellOrderBookKey(
		index,
	))
}

// GetAllSellOrderBook returns all sellOrderBook
func (k Keeper) GetAllSellOrderBook(ctx sdk.Context) (list []types.SellOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellOrderBookKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SellOrderBook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
