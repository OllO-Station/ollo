package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/dex/types"
)

// SetBuyOrderBook set a specific buyOrderBook in the store from its index
func (k Keeper) SetBuyOrderBook(ctx sdk.Context, buyOrderBook types.BuyOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKeyPrefix))
	b := k.cdc.MustMarshal(&buyOrderBook)
	store.Set(types.BuyOrderBookKey(
		buyOrderBook.Index,
	), b)
}

// GetBuyOrderBook returns a buyOrderBook from its index
func (k Keeper) GetBuyOrderBook(
	ctx sdk.Context,
	index string,

) (val types.BuyOrderBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKeyPrefix))

	b := store.Get(types.BuyOrderBookKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBuyOrderBook removes a buyOrderBook from the store
func (k Keeper) RemoveBuyOrderBook(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKeyPrefix))
	store.Delete(types.BuyOrderBookKey(
		index,
	))
}

// GetAllBuyOrderBook returns all buyOrderBook
func (k Keeper) GetAllBuyOrderBook(ctx sdk.Context) (list []types.BuyOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BuyOrderBook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
