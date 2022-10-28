package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/loan/types"
)

// GetLoansCount get the total number of loans
func (k Keeper) GetLoansCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LoansCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLoansCount set the total number of loans
func (k Keeper) SetLoansCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LoansCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLoans appends a loans in the store with a new id and update the count
func (k Keeper) AppendLoans(
	ctx sdk.Context,
	loans types.Loans,
) uint64 {
	// Create the loans
	count := k.GetLoansCount(ctx)

	// Set the ID of the appended value
	loans.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoansKey))
	appendedValue := k.cdc.MustMarshal(&loans)
	store.Set(GetLoansIDBytes(loans.Id), appendedValue)

	// Update loans count
	k.SetLoansCount(ctx, count+1)

	return count
}

// SetLoans set a specific loans in the store
func (k Keeper) SetLoans(ctx sdk.Context, loans types.Loans) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoansKey))
	b := k.cdc.MustMarshal(&loans)
	store.Set(GetLoansIDBytes(loans.Id), b)
}

// GetLoans returns a loans from its id
func (k Keeper) GetLoans(ctx sdk.Context, id uint64) (val types.Loans, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoansKey))
	b := store.Get(GetLoansIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLoans removes a loans from the store
func (k Keeper) RemoveLoans(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoansKey))
	store.Delete(GetLoansIDBytes(id))
}

// GetAllLoans returns all loans
func (k Keeper) GetAllLoans(ctx sdk.Context) (list []types.Loans) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoansKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Loans
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLoansIDBytes returns the byte representation of the ID
func GetLoansIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLoansIDFromBytes returns ID in uint64 format from a byte array
func GetLoansIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
