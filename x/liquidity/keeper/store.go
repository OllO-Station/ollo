package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/liquidity/types"
)

// GetLastPairId returns the last pair id.
func (k Keeper) GetLastPairId(ctx sdk.Context) (id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastPairIdKey)
	if bz == nil {
		id = 0 // initialize the pair id
	} else {
		var val gogotypes.UInt64Value
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return
}

// SetLastPairId stores the last pair id.
func (k Keeper) SetLastPairId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.LastPairIdKey, bz)
}

// GetPair returns pair object for the given pair id.
func (k Keeper) GetPair(ctx sdk.Context, id uint64) (pair types.Pair, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPairKey(id))
	if bz == nil {
		return
	}
	pair = types.MustUnmarshalPair(k.cdc, bz)
	return pair, true
}

// GetPairByDenoms returns a types.Pair for given denoms.
func (k Keeper) GetPairByDenoms(ctx sdk.Context, baseCoinDenom, quoteCoinDenom string) (pair types.Pair, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPairIndexKey(baseCoinDenom, quoteCoinDenom))
	if bz == nil {
		return
	}
	var val gogotypes.UInt64Value
	k.cdc.MustUnmarshal(bz, &val)
	pair, found = k.GetPair(ctx, val.Value)
	return
}

// SetPair stores the particular pair.
func (k Keeper) SetPair(ctx sdk.Context, pair types.Pair) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalPair(k.cdc, pair)
	store.Set(types.GetPairKey(pair.Id), bz)
}

// SetPairIndex stores a pair index.
func (k Keeper) SetPairIndex(ctx sdk.Context, baseCoinDenom, quoteCoinDenom string, pairId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: pairId})
	store.Set(types.GetPairIndexKey(baseCoinDenom, quoteCoinDenom), bz)
}

// SetPairLookupIndex stores a pair lookup index for given denoms.
func (k Keeper) SetPairLookupIndex(ctx sdk.Context, denomA string, denomB string, pairId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPairsByDenomsIndexKey(denomA, denomB, pairId), []byte{})
}

// IterateAllPairs iterates over all the stored pairs and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllPairs(ctx sdk.Context, cb func(pair types.Pair) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PairKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		pair := types.MustUnmarshalPair(k.cdc, iter.Value())
		stop, err := cb(pair)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetAllPairs returns all pairs in the store.
func (k Keeper) GetAllPairs(ctx sdk.Context) (pairs []types.Pair) {
	pairs = []types.Pair{}
	_ = k.IterateAllPairs(ctx, func(pair types.Pair) (stop bool, err error) {
		pairs = append(pairs, pair)
		return false, nil
	})
	return pairs
}

// GetLastPoolId returns the last pool id.
func (k Keeper) GetLastPoolId(ctx sdk.Context) (id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastPoolIdKey)
	if bz == nil {
		id = 0 // initialize the pool id
	} else {
		var val gogotypes.UInt64Value
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return
}

// SetLastPoolId stores the last pool id.
func (k Keeper) SetLastPoolId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.LastPoolIdKey, bz)
}

// GetPool returns pool object for the given pool id.
func (k Keeper) GetPool(ctx sdk.Context, id uint64) (pool types.Pool, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPoolKey(id))
	if bz == nil {
		return
	}
	pool = types.MustUnmarshalPool(k.cdc, bz)
	return pool, true
}

// GetPoolByReserveAddress returns pool object for the given reserve account address.
func (k Keeper) GetPoolByReserveAddress(ctx sdk.Context, reserveAddr sdk.AccAddress) (pool types.Pool, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPoolByReserveAddressIndexKey(reserveAddr))
	if bz == nil {
		return
	}
	var val gogotypes.UInt64Value
	k.cdc.MustUnmarshal(bz, &val)
	poolId := val.GetValue()
	return k.GetPool(ctx, poolId)
}

// SetPool stores the particular pool.
func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalPool(k.cdc, pool)
	store.Set(types.GetPoolKey(pool.Id), bz)
}

// SetPoolByReserveIndex stores a pool by reserve account index key.
func (k Keeper) SetPoolByReserveIndex(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: pool.Id})
	store.Set(types.GetPoolByReserveAddressIndexKey(pool.GetReserveAddress()), bz)
}

// SetPoolsByPairIndex stores a pool by pair index key.
func (k Keeper) SetPoolsByPairIndex(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPoolsByPairIndexKey(pool.PairId, pool.Id), []byte{})
}

// IterateAllPools iterates over all the stored pools and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllPools(ctx sdk.Context, cb func(pool types.Pool) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PoolKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		pool := types.MustUnmarshalPool(k.cdc, iter.Value())
		stop, err := cb(pool)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// IteratePoolsByPair iterates over all the stored pools by the pair and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IteratePoolsByPair(ctx sdk.Context, pairId uint64, cb func(pool types.Pool) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetPoolsByPairIndexKeyPrefix(pairId))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		poolId := types.ParsePoolsByPairIndexKey(iter.Key())
		pool, _ := k.GetPool(ctx, poolId)
		stop, err := cb(pool)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetAllPools returns all pools in the store.
func (k Keeper) GetAllPools(ctx sdk.Context) (pools []types.Pool) {
	pools = []types.Pool{}
	_ = k.IterateAllPools(ctx, func(pool types.Pool) (stop bool, err error) {
		pools = append(pools, pool)
		return false, nil
	})
	return
}

// GetPoolsByPair returns pools within the pair.
func (k Keeper) GetPoolsByPair(ctx sdk.Context, pairId uint64) (pools []types.Pool) {
	_ = k.IteratePoolsByPair(ctx, pairId, func(pool types.Pool) (stop bool, err error) {
		pools = append(pools, pool)
		return false, nil
	})
	return
}

// GetRequestDeposit returns the particular deposit request.
func (k Keeper) GetRequestDeposit(ctx sdk.Context, poolId, id uint64) (req types.RequestDeposit, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRequestDepositKey(poolId, id))
	if bz == nil {
		return
	}
	req = types.MustUnmarshalRequestDeposit(k.cdc, bz)
	return req, true
}

// SetRequestDeposit stores deposit request for the batch execution.
func (k Keeper) SetRequestDeposit(ctx sdk.Context, req types.RequestDeposit) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalRequestDeposit(k.cdc, req)
	store.Set(types.GetRequestDepositKey(req.PoolId, req.Id), bz)
}

func (k Keeper) SetRequestDepositIndex(ctx sdk.Context, req types.RequestDeposit) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetRequestDepositIndexKey(req.GetDepositor(), req.PoolId, req.Id), []byte{})
}

// IterateAllRequestDeposits iterates through all deposit requests in the store
// and call cb for each request.
func (k Keeper) IterateAllRequestDeposits(ctx sdk.Context, cb func(req types.RequestDeposit) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.RequestDepositKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		req := types.MustUnmarshalRequestDeposit(k.cdc, iter.Value())
		stop, err := cb(req)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// IterateRequestDepositsByDepositor iterates through deposit requests in the
// store by a depositor and call cb on each order.
func (k Keeper) IterateRequestDepositsByDepositor(ctx sdk.Context, depositor sdk.AccAddress, cb func(req types.RequestDeposit) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetRequestDepositIndexKeyPrefix(depositor))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		_, poolId, reqId := types.ParseRequestDepositIndexKey(iter.Key())
		req, _ := k.GetRequestDeposit(ctx, poolId, reqId)
		stop, err := cb(req)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetAllRequestDeposits returns all deposit requests in the store.
func (k Keeper) GetAllRequestDeposits(ctx sdk.Context) (reqs []types.RequestDeposit) {
	reqs = []types.RequestDeposit{}
	_ = k.IterateAllRequestDeposits(ctx, func(req types.RequestDeposit) (stop bool, err error) {
		reqs = append(reqs, req)
		return false, nil
	})
	return
}

// GetRequestDepositsByDepositor returns deposit requests by the depositor.
func (k Keeper) GetRequestDepositsByDepositor(ctx sdk.Context, depositor sdk.AccAddress) (reqs []types.RequestDeposit) {
	_ = k.IterateRequestDepositsByDepositor(ctx, depositor, func(req types.RequestDeposit) (stop bool, err error) {
		reqs = append(reqs, req)
		return false, nil
	})
	return
}

// DeleteRequestDeposit deletes a deposit request.
func (k Keeper) DeleteRequestDeposit(ctx sdk.Context, req types.RequestDeposit) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetRequestDepositKey(req.PoolId, req.Id))
	k.DeleteRequestDepositIndex(ctx, req)
}

func (k Keeper) DeleteRequestDepositIndex(ctx sdk.Context, req types.RequestDeposit) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetRequestDepositIndexKey(req.GetDepositor(), req.PoolId, req.Id))
}

// GetRequestWithdraw returns the particular withdraw request.
func (k Keeper) GetRequestWithdraw(ctx sdk.Context, poolId, id uint64) (req types.RequestWithdraw, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRequestWithdrawKey(poolId, id))
	if bz == nil {
		return
	}
	req = types.MustUnmarshalRequestWithdraw(k.cdc, bz)
	return req, true
}

// SetRequestWithdraw stores withdraw request for the batch execution.
func (k Keeper) SetRequestWithdraw(ctx sdk.Context, req types.RequestWithdraw) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshaRequestWithdraw(k.cdc, req)
	store.Set(types.GetRequestWithdrawKey(req.PoolId, req.Id), bz)
}

func (k Keeper) SetRequestWithdrawIndex(ctx sdk.Context, req types.RequestWithdraw) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetRequestWithdrawIndexKey(req.GetWithdrawer(), req.PoolId, req.Id), []byte{})
}

// IterateAllRequestWithdraws iterates through all withdraw requests in the store
// and call cb for each request.
func (k Keeper) IterateAllRequestWithdraws(ctx sdk.Context, cb func(req types.RequestWithdraw) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.RequestWithdrawKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		req := types.MustUnmarshalRequestWithdraw(k.cdc, iter.Value())
		stop, err := cb(req)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// IterateRequestWithdrawsByWithdrawer iterates through withdraw requests in the
// store by a withdrawer and call cb on each order.
func (k Keeper) IterateRequestWithdrawsByWithdrawer(ctx sdk.Context, withdrawer sdk.AccAddress, cb func(req types.RequestWithdraw) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetRequestWithdrawIndexKeyPrefix(withdrawer))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		_, poolId, reqId := types.ParseRequestWithdrawIndexKey(iter.Key())
		req, _ := k.GetRequestWithdraw(ctx, poolId, reqId)
		stop, err := cb(req)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetAllRequestWithdraws returns all withdraw requests in the store.
func (k Keeper) GetAllRequestWithdraws(ctx sdk.Context) (reqs []types.RequestWithdraw) {
	reqs = []types.RequestWithdraw{}
	_ = k.IterateAllRequestWithdraws(ctx, func(req types.RequestWithdraw) (stop bool, err error) {
		reqs = append(reqs, req)
		return false, nil
	})
	return
}

// GetRequestWithdrawsByWithdrawer returns withdraw requests by the withdrawer.
func (k Keeper) GetRequestWithdrawsByWithdrawer(ctx sdk.Context, withdrawer sdk.AccAddress) (reqs []types.RequestWithdraw) {
	_ = k.IterateRequestWithdrawsByWithdrawer(ctx, withdrawer, func(req types.RequestWithdraw) (stop bool, err error) {
		reqs = append(reqs, req)
		return false, nil
	})
	return
}

// DeleteRequestWithdraw deletes a withdraw request.
func (k Keeper) DeleteRequestWithdraw(ctx sdk.Context, req types.RequestWithdraw) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetRequestWithdrawKey(req.PoolId, req.Id))
	k.DeleteRequestWithdrawIndex(ctx, req)
}

func (k Keeper) DeleteRequestWithdrawIndex(ctx sdk.Context, req types.RequestWithdraw) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetRequestDepositIndexKey(req.GetWithdrawer(), req.PoolId, req.Id))
}

// GetOrder returns the particular order.
func (k Keeper) GetOrder(ctx sdk.Context, pairId, id uint64) (order types.Order, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetOrderKey(pairId, id))
	if bz == nil {
		return
	}
	order = types.MustUnmarshalOrder(k.cdc, bz)
	return order, true
}

// SetOrder stores an order for the batch execution.
func (k Keeper) SetOrder(ctx sdk.Context, order types.Order) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshaOrder(k.cdc, order)
	store.Set(types.GetOrderKey(order.PairId, order.Id), bz)
}

func (k Keeper) SetOrderIndex(ctx sdk.Context, order types.Order) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetOrderIndexKey(order.GetOrderer(), order.PairId, order.Id), []byte{})
}

// IterateAllOrders iterates through all orders in the store and all
// cb for each order.
func (k Keeper) IterateAllOrders(ctx sdk.Context, cb func(order types.Order) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.OrderKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		order := types.MustUnmarshalOrder(k.cdc, iter.Value())
		stop, err := cb(order)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// IterateOrdersByPair iterates through all the orders within the pair
// and call cb for each order.
func (k Keeper) IterateOrdersByPair(ctx sdk.Context, pairId uint64, cb func(order types.Order) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetOrdersByPairKeyPrefix(pairId))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		order := types.MustUnmarshalOrder(k.cdc, iter.Value())
		stop, err := cb(order)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// IterateOrdersByOrderer iterates through orders in the store by an orderer
// and call cb on each order.
func (k Keeper) IterateOrdersByOrderer(ctx sdk.Context, orderer sdk.AccAddress, cb func(order types.Order) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetOrderIndexKeyPrefix(orderer))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		_, pairId, orderId := types.ParseOrderIndexKey(iter.Key())
		order, _ := k.GetOrder(ctx, pairId, orderId)
		stop, err := cb(order)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetAllOrders returns all orders in the store.
func (k Keeper) GetAllOrders(ctx sdk.Context) (orders []types.Order) {
	orders = []types.Order{}
	_ = k.IterateAllOrders(ctx, func(order types.Order) (stop bool, err error) {
		orders = append(orders, order)
		return false, nil
	})
	return
}

// GetOrdersByPair returns orders within the pair.
func (k Keeper) GetOrdersByPair(ctx sdk.Context, pairId uint64) (orders []types.Order) {
	_ = k.IterateOrdersByPair(ctx, pairId, func(order types.Order) (stop bool, err error) {
		orders = append(orders, order)
		return false, nil
	})
	return
}

// GetOrdersByOrderer returns orders by the orderer.
func (k Keeper) GetOrdersByOrderer(ctx sdk.Context, orderer sdk.AccAddress) (orders []types.Order) {
	_ = k.IterateOrdersByOrderer(ctx, orderer, func(order types.Order) (stop bool, err error) {
		orders = append(orders, order)
		return false, nil
	})
	return
}

// DeleteOrder deletes an order.
func (k Keeper) DeleteOrder(ctx sdk.Context, order types.Order) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetOrderKey(order.PairId, order.Id))
	k.DeleteOrderIndex(ctx, order)
}

func (k Keeper) DeleteOrderIndex(ctx sdk.Context, order types.Order) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetOrderIndexKey(order.GetOrderer(), order.PairId, order.Id))
}

// GetMarketMakingOrderId returns the market making order index.
func (k Keeper) GetMarketMakingOrderId(ctx sdk.Context, orderer sdk.AccAddress, pairId uint64) (index types.MarketMakingOrderId, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetMMOrderIndexKey(orderer, pairId))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &index)
	return index, true
}

// SetMarketMakingOrderId stores a market making order index.
func (k Keeper) SetMarketMakingOrderId(ctx sdk.Context, index types.MarketMakingOrderId) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&index)
	store.Set(types.GetMMOrderIndexKey(index.GetOrderer(), index.PairId), bz)
}

// IterateAllMarketMakingOrderIdes iterates through all market making order indexes
// in the store and call cb for each order.
func (k Keeper) IterateAllMarketMakingOrderIdes(ctx sdk.Context, cb func(index types.MarketMakingOrderId) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.MMOrderIndexKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var index types.MarketMakingOrderId
		k.cdc.MustUnmarshal(iter.Value(), &index)
		stop, err := cb(index)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetAllMarketMakingOrderIdes returns all market making order indexes in the store.
func (k Keeper) GetAllMarketMakingOrderIdes(ctx sdk.Context) (indexes []types.MarketMakingOrderId) {
	indexes = []types.MarketMakingOrderId{}
	_ = k.IterateAllMarketMakingOrderIdes(ctx, func(index types.MarketMakingOrderId) (stop bool, err error) {
		indexes = append(indexes, index)
		return false, nil
	})
	return
}

// DeleteMarketMakingOrderId deletes a market making order index.
func (k Keeper) DeleteMarketMakingOrderId(ctx sdk.Context, index types.MarketMakingOrderId) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetMMOrderIndexKey(index.GetOrderer(), index.PairId))
}
