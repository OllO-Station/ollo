package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ollo-station/ollo/x/lend/types"
)

// GetLoanCount get the total number of loan
func (k Keeper) GetLoanCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LoanCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLoanCount set the total number of loan
func (k Keeper) SetLoanCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LoanCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLoan appends a loan in the store with a new id and update the count
func (k Keeper) AppendLoan(
	ctx sdk.Context,
	loan types.Loan,
) uint64 {
	// Create the loan
	count := k.GetLoanCount(ctx)

	// Set the ID of the appended value
	loan.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoanKey))
	appendedValue := k.cdc.MustMarshal(&loan)
	store.Set(GetLoanIDBytes(loan.Id), appendedValue)

	// Update loan count
	k.SetLoanCount(ctx, count+1)

	return count
}

// SetLoan set a specific loan in the store
func (k Keeper) SetLoan(ctx sdk.Context, loan types.Loan) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoanKey))
	b := k.cdc.MustMarshal(&loan)
	store.Set(GetLoanIDBytes(loan.Id), b)
}

// GetLoan returns a loan from its id
func (k Keeper) GetLoan(ctx sdk.Context, id uint64) (val types.Loan, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoanKey))
	b := store.Get(GetLoanIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLoan removes a loan from the store
func (k Keeper) RemoveLoan(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoanKey))
	store.Delete(GetLoanIDBytes(id))
}

// GetAllLoan returns all loan
func (k Keeper) GetAllLoan(ctx sdk.Context) (list []types.Loan) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoanKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Loan
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLoanIDBytes returns the byte representation of the ID
func GetLoanIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLoanIDFromBytes returns ID in uint64 format from a byte array
func GetLoanIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
