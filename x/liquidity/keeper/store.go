package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"

	"ollo/x/liquidity/types"
)

// GetPool reads from kvstore and returns a specific pool
func (k Keeper) GetPool(ctx sdk.Context, poolID uint64) (pool types.Pool, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPoolKey(poolID)

	value := store.Get(key)
	if value == nil {
		return pool, false
	}

	pool = types.MustUnmarshalPool(k.cdc, value)

	return pool, true
}

// SetPool sets to kvstore a specific pool
func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalPool(k.cdc, pool)
	store.Set(types.GetPoolKey(pool.Id), b)
}

// delete from kvstore a specific liquidityPool
func (k Keeper) DeletePool(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	Key := types.GetPoolKey(pool.Id)
	store.Delete(Key)
}

// IterateAllPools iterate through all of the liquidityPools
func (k Keeper) IterateAllPools(ctx sdk.Context, cb func(pool types.Pool) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PoolKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		pool := types.MustUnmarshalPool(k.cdc, iterator.Value())
		if cb(pool) {
			break
		}
	}
}

// GetAllPools returns all liquidityPools used during genesis dump
func (k Keeper) GetAllPools(ctx sdk.Context) (pools []types.Pool) {
	k.IterateAllPools(ctx, func(liquidityPool types.Pool) bool {
		pools = append(pools, liquidityPool)
		return false
	})

	return pools
}

// GetNextPoolIDWithUpdate returns and increments the global Pool ID counter.
// If the global account number is not set, it initializes it with value 0.
func (k Keeper) GetNextPoolIDWithUpdate(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	poolID := k.GetNextPoolID(ctx)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: poolID + 1})
	store.Set(types.GlobalLiquidityPoolIDKey, bz)
	return poolID
}

// GetNextPoolID returns next pool id for new pool, using index of latest pool id
func (k Keeper) GetNextPoolID(ctx sdk.Context) uint64 {
	var poolID uint64
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GlobalLiquidityPoolIDKey)
	if bz == nil {
		// initialize the LiquidityPoolID
		poolID = 1
	} else {
		val := gogotypes.UInt64Value{}

		err := k.cdc.Unmarshal(bz, &val)
		if err != nil {
			panic(err)
		}

		poolID = val.GetValue()
	}
	return poolID
}

// GetPoolByReserveAccIndex reads from kvstore and return a specific liquidityPool indexed by given reserve account
func (k Keeper) GetPoolByReserveAccIndex(ctx sdk.Context, reserveAcc sdk.AccAddress) (pool types.Pool, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPoolByReserveAccIndexKey(reserveAcc)

	value := store.Get(key)
	if value == nil {
		return pool, false
	}

	val := gogotypes.UInt64Value{}
	err := k.cdc.Unmarshal(value, &val)
	if err != nil {
		return pool, false
	}
	poolID := val.GetValue()
	return k.GetPool(ctx, poolID)
}

// SetPoolByReserveAccIndex sets Index by ReserveAcc for pool duplication check
func (k Keeper) SetPoolByReserveAccIndex(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: pool.Id})
	store.Set(types.GetPoolByReserveAccIndexKey(pool.GetReserveAccount()), b)
}

// SetPoolAtomic sets pool with set global pool id index +1 and index by reserveAcc
func (k Keeper) SetPoolAtomic(ctx sdk.Context, pool types.Pool) types.Pool {
	pool.Id = k.GetNextPoolIDWithUpdate(ctx)
	k.SetPool(ctx, pool)
	k.SetPoolByReserveAccIndex(ctx, pool)
	return pool
}

// GetPoolBatch returns a specific pool batch
func (k Keeper) GetPoolBatch(ctx sdk.Context, poolID uint64) (poolBatch types.PoolBatch, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPoolBatchKey(poolID)

	value := store.Get(key)
	if value == nil {
		return poolBatch, false
	}

	poolBatch = types.MustUnmarshalPoolBatch(k.cdc, value)

	return poolBatch, true
}

// GetAllPoolBatches returns all batches of the all existed liquidity pools
func (k Keeper) GetAllPoolBatches(ctx sdk.Context) (poolBatches []types.PoolBatch) {
	k.IterateAllPoolBatches(ctx, func(poolBatch types.PoolBatch) bool {
		poolBatches = append(poolBatches, poolBatch)
		return false
	})

	return poolBatches
}

// IterateAllPoolBatches iterate through all of the pool batches
func (k Keeper) IterateAllPoolBatches(ctx sdk.Context, cb func(poolBatch types.PoolBatch) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PoolBatchKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		poolBatch := types.MustUnmarshalPoolBatch(k.cdc, iterator.Value())
		if cb(poolBatch) {
			break
		}
	}
}

// DeletePoolBatch deletes batch of the pool, it used for test case
func (k Keeper) DeletePoolBatch(ctx sdk.Context, poolBatch types.PoolBatch) {
	store := ctx.KVStore(k.storeKey)
	batchKey := types.GetPoolBatchKey(poolBatch.PoolId)
	store.Delete(batchKey)
}

// SetPoolBatch sets batch of the pool, with current state
func (k Keeper) SetPoolBatch(ctx sdk.Context, poolBatch types.PoolBatch) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalPoolBatch(k.cdc, poolBatch)
	store.Set(types.GetPoolBatchKey(poolBatch.PoolId), b)
}

// GetPoolBatchDepositMsgState returns a specific DepositMsgState
func (k Keeper) GetPoolBatchDepositMsgState(ctx sdk.Context, poolID, msgIndex uint64) (state types.DepositMsgState, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPoolBatchDepositMsgStateIndexKey(poolID, msgIndex)

	value := store.Get(key)
	if value == nil {
		return state, false
	}

	state = types.MustUnmarshalDepositMsgState(k.cdc, value)
	return state, true
}

// SetPoolBatchDepositMsgState sets deposit msg state of the pool batch, with current state
func (k Keeper) SetPoolBatchDepositMsgState(ctx sdk.Context, poolID uint64, state types.DepositMsgState) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalDepositMsgState(k.cdc, state)
	store.Set(types.GetPoolBatchDepositMsgStateIndexKey(poolID, state.MsgIndex), b)
}

// SetPoolBatchDepositMsgStatesByPointer sets deposit batch msgs of the pool batch, with current state using pointers
func (k Keeper) SetPoolBatchDepositMsgStatesByPointer(ctx sdk.Context, poolID uint64, states []*types.DepositMsgState) {
	store := ctx.KVStore(k.storeKey)
	for _, state := range states {
		if poolID != state.Msg.PoolId {
			continue
		}
		b := types.MustMarshalDepositMsgState(k.cdc, *state)
		store.Set(types.GetPoolBatchDepositMsgStateIndexKey(poolID, state.MsgIndex), b)
	}
}

// SetPoolBatchDepositMsgStates sets deposit batch msgs of the pool batch, with current state
func (k Keeper) SetPoolBatchDepositMsgStates(ctx sdk.Context, poolID uint64, states []types.DepositMsgState) {
	store := ctx.KVStore(k.storeKey)
	for _, state := range states {
		if poolID != state.Msg.PoolId {
			continue
		}
		b := types.MustMarshalDepositMsgState(k.cdc, state)
		store.Set(types.GetPoolBatchDepositMsgStateIndexKey(poolID, state.MsgIndex), b)
	}
}

// IterateAllPoolBatchDepositMsgStates iterate through all of the DepositMsgStates in the batch
func (k Keeper) IterateAllPoolBatchDepositMsgStates(ctx sdk.Context, poolBatch types.PoolBatch, cb func(state types.DepositMsgState) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.GetPoolBatchDepositMsgStatesPrefix(poolBatch.PoolId)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		state := types.MustUnmarshalDepositMsgState(k.cdc, iterator.Value())
		if cb(state) {
			break
		}
	}
}

// IterateAllDepositMsgStates iterate through all of the DepositMsgState of all batches
func (k Keeper) IterateAllDepositMsgStates(ctx sdk.Context, cb func(state types.DepositMsgState) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.PoolBatchDepositMsgStateIndexKeyPrefix
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		state := types.MustUnmarshalDepositMsgState(k.cdc, iterator.Value())
		if cb(state) {
			break
		}
	}
}

// GetAllDepositMsgStates returns all BatchDepositMsgs for all batches.
func (k Keeper) GetAllDepositMsgStates(ctx sdk.Context) (states []types.DepositMsgState) {
	k.IterateAllDepositMsgStates(ctx, func(state types.DepositMsgState) bool {
		states = append(states, state)
		return false
	})
	return states
}

// GetAllPoolBatchDepositMsgs returns all BatchDepositMsgs indexed by the pool batch
func (k Keeper) GetAllPoolBatchDepositMsgs(ctx sdk.Context, poolBatch types.PoolBatch) (states []types.DepositMsgState) {
	k.IterateAllPoolBatchDepositMsgStates(ctx, poolBatch, func(state types.DepositMsgState) bool {
		states = append(states, state)
		return false
	})
	return states
}

// GetAllPoolBatchDepositMsgStatesNotToBeDeleted returns all Not toDelete BatchDepositMsgs indexed by the liquidityPoolBatch
func (k Keeper) GetAllPoolBatchDepositMsgStatesNotToBeDeleted(ctx sdk.Context, poolBatch types.PoolBatch) (states []types.DepositMsgState) {
	k.IterateAllPoolBatchDepositMsgStates(ctx, poolBatch, func(state types.DepositMsgState) bool {
		if !state.ToBeDeleted {
			states = append(states, state)
		}
		return false
	})
	return states
}

// GetAllRemainingPoolBatchDepositMsgStates returns all remaining DepositMsgStates after endblock,
// which are executed but not to be deleted
func (k Keeper) GetAllRemainingPoolBatchDepositMsgStates(ctx sdk.Context, poolBatch types.PoolBatch) (states []*types.DepositMsgState) {
	k.IterateAllPoolBatchDepositMsgStates(ctx, poolBatch, func(state types.DepositMsgState) bool {
		if state.Executed && !state.ToBeDeleted {
			states = append(states, &state)
		}
		return false
	})
	return states
}

// delete deposit batch msgs of the liquidity pool batch which has state ToBeDeleted
func (k Keeper) DeleteAllReadyPoolBatchDepositMsgStates(ctx sdk.Context, poolBatch types.PoolBatch) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetPoolBatchDepositMsgStatesPrefix(poolBatch.PoolId))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		state := types.MustUnmarshalDepositMsgState(k.cdc, iterator.Value())
		if state.ToBeDeleted {
			store.Delete(iterator.Key())
		}
	}
}

// return a specific liquidityPoolBatchWithdrawMsg
func (k Keeper) GetPoolBatchWithdrawMsgState(ctx sdk.Context, poolID, msgIndex uint64) (state types.WithdrawMsgState, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPoolBatchWithdrawMsgStateIndexKey(poolID, msgIndex)

	value := store.Get(key)
	if value == nil {
		return state, false
	}

	state = types.MustUnmarshalWithdrawMsgState(k.cdc, value)
	return state, true
}

// set withdraw batch msg of the liquidity pool batch, with current state
func (k Keeper) SetPoolBatchWithdrawMsgState(ctx sdk.Context, poolID uint64, state types.WithdrawMsgState) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalWithdrawMsgState(k.cdc, state)
	store.Set(types.GetPoolBatchWithdrawMsgStateIndexKey(poolID, state.MsgIndex), b)
}

// set withdraw batch msgs of the liquidity pool batch, with current state using pointers
func (k Keeper) SetPoolBatchWithdrawMsgStatesByPointer(ctx sdk.Context, poolID uint64, states []*types.WithdrawMsgState) {
	store := ctx.KVStore(k.storeKey)
	for _, state := range states {
		if poolID != state.Msg.PoolId {
			continue
		}
		b := types.MustMarshalWithdrawMsgState(k.cdc, *state)
		store.Set(types.GetPoolBatchWithdrawMsgStateIndexKey(poolID, state.MsgIndex), b)
	}
}

// set withdraw batch msgs of the pool batch, with current state
func (k Keeper) SetPoolBatchWithdrawMsgStates(ctx sdk.Context, poolID uint64, states []types.WithdrawMsgState) {
	store := ctx.KVStore(k.storeKey)
	for _, state := range states {
		if poolID != state.Msg.PoolId {
			continue
		}
		b := types.MustMarshalWithdrawMsgState(k.cdc, state)
		store.Set(types.GetPoolBatchWithdrawMsgStateIndexKey(poolID, state.MsgIndex), b)
	}
}

// IterateAllPoolBatchWithdrawMsgStates iterate through all of the LiquidityPoolBatchWithdrawMsgs
func (k Keeper) IterateAllPoolBatchWithdrawMsgStates(ctx sdk.Context, poolBatch types.PoolBatch, cb func(state types.WithdrawMsgState) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.GetPoolBatchWithdrawMsgsPrefix(poolBatch.PoolId)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		state := types.MustUnmarshalWithdrawMsgState(k.cdc, iterator.Value())
		if cb(state) {
			break
		}
	}
}

// IterateAllWithdrawMsgStates iterate through all of the WithdrawMsgState of all batches
func (k Keeper) IterateAllWithdrawMsgStates(ctx sdk.Context, cb func(state types.WithdrawMsgState) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.PoolBatchWithdrawMsgStateIndexKeyPrefix
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		state := types.MustUnmarshalWithdrawMsgState(k.cdc, iterator.Value())
		if cb(state) {
			break
		}
	}
}

// GetAllWithdrawMsgStates returns all BatchWithdrawMsgs for all batches
func (k Keeper) GetAllWithdrawMsgStates(ctx sdk.Context) (states []types.WithdrawMsgState) {
	k.IterateAllWithdrawMsgStates(ctx, func(state types.WithdrawMsgState) bool {
		states = append(states, state)
		return false
	})
	return states
}

// GetAllPoolBatchWithdrawMsgStates returns all BatchWithdrawMsgs indexed by the liquidityPoolBatch
func (k Keeper) GetAllPoolBatchWithdrawMsgStates(ctx sdk.Context, poolBatch types.PoolBatch) (states []types.WithdrawMsgState) {
	k.IterateAllPoolBatchWithdrawMsgStates(ctx, poolBatch, func(state types.WithdrawMsgState) bool {
		states = append(states, state)
		return false
	})
	return states
}

// GetAllPoolBatchWithdrawMsgStatesNotToBeDeleted returns all Not to delete BatchWithdrawMsgs indexed by the liquidityPoolBatch
func (k Keeper) GetAllPoolBatchWithdrawMsgStatesNotToBeDeleted(ctx sdk.Context, poolBatch types.PoolBatch) (states []types.WithdrawMsgState) {
	k.IterateAllPoolBatchWithdrawMsgStates(ctx, poolBatch, func(state types.WithdrawMsgState) bool {
		if !state.ToBeDeleted {
			states = append(states, state)
		}
		return false
	})
	return states
}

// GetAllRemainingPoolBatchWithdrawMsgStates returns All only remaining BatchWithdrawMsgs after endblock, executed but not toDelete
func (k Keeper) GetAllRemainingPoolBatchWithdrawMsgStates(ctx sdk.Context, poolBatch types.PoolBatch) (states []*types.WithdrawMsgState) {
	k.IterateAllPoolBatchWithdrawMsgStates(ctx, poolBatch, func(state types.WithdrawMsgState) bool {
		if state.Executed && !state.ToBeDeleted {
			states = append(states, &state)
		}
		return false
	})
	return states
}

// delete withdraw batch msgs of the liquidity pool batch which has state ToBeDeleted
func (k Keeper) DeleteAllReadyPoolBatchWithdrawMsgStates(ctx sdk.Context, poolBatch types.PoolBatch) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetPoolBatchWithdrawMsgsPrefix(poolBatch.PoolId))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		state := types.MustUnmarshalWithdrawMsgState(k.cdc, iterator.Value())
		if state.ToBeDeleted {
			store.Delete(iterator.Key())
		}
	}
}

// return a specific SwapMsgState given the pool_id with the msg_index
func (k Keeper) GetPoolBatchSwapMsgState(ctx sdk.Context, poolID, msgIndex uint64) (state types.SwapMsgState, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPoolBatchSwapMsgStateIndexKey(poolID, msgIndex)

	value := store.Get(key)
	if value == nil {
		return state, false
	}

	state = types.MustUnmarshalSwapMsgState(k.cdc, value)
	return state, true
}

// set swap batch msg of the liquidity pool batch, with current state
func (k Keeper) SetPoolBatchSwapMsgState(ctx sdk.Context, poolID uint64, state types.SwapMsgState) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalSwapMsgState(k.cdc, state)
	store.Set(types.GetPoolBatchSwapMsgStateIndexKey(poolID, state.MsgIndex), b)
}

// Delete swap batch msg of the liquidity pool batch, it used for test case
func (k Keeper) DeletePoolBatchSwapMsgState(ctx sdk.Context, poolID uint64, msgIndex uint64) {
	store := ctx.KVStore(k.storeKey)
	batchKey := types.GetPoolBatchSwapMsgStateIndexKey(poolID, msgIndex)
	store.Delete(batchKey)
}

// IterateAllPoolBatchSwapMsgStates iterate through all of the LiquidityPoolBatchSwapMsgs
func (k Keeper) IterateAllPoolBatchSwapMsgStates(ctx sdk.Context, poolBatch types.PoolBatch, cb func(state types.SwapMsgState) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.GetPoolBatchSwapMsgStatesPrefix(poolBatch.PoolId)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		state := types.MustUnmarshalSwapMsgState(k.cdc, iterator.Value())
		if cb(state) {
			break
		}
	}
}

// IterateAllSwapMsgStates iterate through all of the SwapMsgState of all batches
func (k Keeper) IterateAllSwapMsgStates(ctx sdk.Context, cb func(state types.SwapMsgState) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.PoolBatchSwapMsgStateIndexKeyPrefix
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		state := types.MustUnmarshalSwapMsgState(k.cdc, iterator.Value())
		if cb(state) {
			break
		}
	}
}

// GetAllSwapMsgStates returns all BatchSwapMsgs of all batches
func (k Keeper) GetAllSwapMsgStates(ctx sdk.Context) (states []types.SwapMsgState) {
	k.IterateAllSwapMsgStates(ctx, func(state types.SwapMsgState) bool {
		states = append(states, state)
		return false
	})
	return states
}

// delete swap batch msgs of the liquidity pool batch which has state ToBeDeleted
func (k Keeper) DeleteAllReadyPoolBatchSwapMsgStates(ctx sdk.Context, poolBatch types.PoolBatch) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetPoolBatchSwapMsgStatesPrefix(poolBatch.PoolId))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		state := types.MustUnmarshalSwapMsgState(k.cdc, iterator.Value())
		if state.ToBeDeleted {
			store.Delete(iterator.Key())
		}
	}
}

// GetAllPoolBatchSwapMsgStatesAsPointer returns all BatchSwapMsgs pointer indexed by the liquidityPoolBatch
func (k Keeper) GetAllPoolBatchSwapMsgStatesAsPointer(ctx sdk.Context, poolBatch types.PoolBatch) (states []*types.SwapMsgState) {
	k.IterateAllPoolBatchSwapMsgStates(ctx, poolBatch, func(state types.SwapMsgState) bool {
		states = append(states, &state)
		return false
	})
	return states
}

// GetAllPoolBatchSwapMsgStates returns all BatchSwapMsgs indexed by the liquidityPoolBatch
func (k Keeper) GetAllPoolBatchSwapMsgStates(ctx sdk.Context, poolBatch types.PoolBatch) (states []types.SwapMsgState) {
	k.IterateAllPoolBatchSwapMsgStates(ctx, poolBatch, func(state types.SwapMsgState) bool {
		states = append(states, state)
		return false
	})
	return states
}

// GetAllNotProcessedPoolBatchSwapMsgStates returns All only not processed swap msgs, not executed with not succeed and not toDelete BatchSwapMsgs indexed by the liquidityPoolBatch
func (k Keeper) GetAllNotProcessedPoolBatchSwapMsgStates(ctx sdk.Context, poolBatch types.PoolBatch) (states []*types.SwapMsgState) {
	k.IterateAllPoolBatchSwapMsgStates(ctx, poolBatch, func(state types.SwapMsgState) bool {
		if !state.Executed && !state.Succeeded && !state.ToBeDeleted {
			states = append(states, &state)
		}
		return false
	})
	return states
}

// GetAllRemainingPoolBatchSwapMsgStates returns All only remaining after endblock swap msgs, executed but not toDelete
func (k Keeper) GetAllRemainingPoolBatchSwapMsgStates(ctx sdk.Context, poolBatch types.PoolBatch) (states []*types.SwapMsgState) {
	k.IterateAllPoolBatchSwapMsgStates(ctx, poolBatch, func(state types.SwapMsgState) bool {
		if state.Executed && !state.ToBeDeleted {
			states = append(states, &state)
		}
		return false
	})
	return states
}

// GetAllPoolBatchSwapMsgStatesNotToBeDeleted returns All only not to delete swap msgs
func (k Keeper) GetAllPoolBatchSwapMsgStatesNotToBeDeleted(ctx sdk.Context, poolBatch types.PoolBatch) (states []*types.SwapMsgState) {
	k.IterateAllPoolBatchSwapMsgStates(ctx, poolBatch, func(state types.SwapMsgState) bool {
		if !state.ToBeDeleted {
			states = append(states, &state)
		}
		return false
	})
	return states
}

// set swap batch msgs of the liquidity pool batch, with current state using pointers
func (k Keeper) SetPoolBatchSwapMsgStatesByPointer(ctx sdk.Context, poolID uint64, states []*types.SwapMsgState) {
	store := ctx.KVStore(k.storeKey)
	for _, state := range states {
		if poolID != state.Msg.PoolId {
			continue
		}
		b := types.MustMarshalSwapMsgState(k.cdc, *state)
		store.Set(types.GetPoolBatchSwapMsgStateIndexKey(poolID, state.MsgIndex), b)
	}
}

// set swap batch msgs of the liquidity pool batch, with current state
func (k Keeper) SetPoolBatchSwapMsgStates(ctx sdk.Context, poolID uint64, states []types.SwapMsgState) {
	store := ctx.KVStore(k.storeKey)
	for _, state := range states {
		if poolID != state.Msg.PoolId {
			continue
		}
		b := types.MustMarshalSwapMsgState(k.cdc, state)
		store.Set(types.GetPoolBatchSwapMsgStateIndexKey(poolID, state.MsgIndex), b)
	}
}

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

// GetPoolsByPair returns pools within the pair.
func (k Keeper) GetPoolsByPair(ctx sdk.Context, pairId uint64) (pools []types.Pool) {
	_ = k.IteratePoolsByPair(ctx, pairId, func(pool types.Pool) (stop bool, err error) {
		pools = append(pools, pool)
		return false, nil
	})
	return
}
