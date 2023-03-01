package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/x/claim/types"
)

// SetClaimRecord set a specific claimRecord in the store from its index
func (k Keeper) SetClaimRecord(ctx sdk.Context, claimRecord types.ClaimRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimRecordKeyPrefix))
	b := k.cdc.MustMarshal(&claimRecord)
	store.Set(types.ClaimRecordKey(
		claimRecord.Address,
	), b)
}

// GetClaimRecord returns a claimRecord from its index
func (k Keeper) GetClaimRecord(ctx sdk.Context, address string) (val types.ClaimRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimRecordKeyPrefix))

	b := store.Get(types.ClaimRecordKey(address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClaimRecord removes a claimRecord from the store
func (k Keeper) RemoveClaimRecord(
	ctx sdk.Context,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimRecordKeyPrefix))
	store.Delete(types.ClaimRecordKey(address))
}

// GetAllClaimRecord returns all claimRecord
func (k Keeper) GetAllClaimRecord(ctx sdk.Context) (list []types.ClaimRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClaimRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetInitialClaim set initialClaim in the store
func (k Keeper) SetInitialClaim(ctx sdk.Context, initialClaim types.InitialClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKey))
	b := k.cdc.MustMarshal(&initialClaim)
	store.Set([]byte{0}, b)
}

// GetInitialClaim returns initialClaim
func (k Keeper) GetInitialClaim(ctx sdk.Context) (val types.InitialClaim, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInitialClaim removes initialClaim from the store
func (k Keeper) RemoveInitialClaim(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKey))
	store.Delete([]byte{0})
}
