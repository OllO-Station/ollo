package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName is the name of the liquidity module
	ModuleName = "liquidity"
)

var (
	PoolByReserveAccIndexKeyPrefix = []byte{0x12}

	PoolBatchIndexKeyPrefix = []byte{0x21} // Last PoolBatchIndex
)

// - PoolByReserveAccIndex: `0x12 | ReserveAcc -> Id`
// GetPoolByReserveAccIndexKey returns kv indexing key of the pool indexed by reserve account
func GetPoolByReserveAccIndexKey(reserveAcc sdk.AccAddress) []byte {
	return append(PoolByReserveAccIndexKeyPrefix, reserveAcc.Bytes()...)
}

// GetPoolBatchIndexKey returns kv indexing key of the latest index value of the pool batch
func GetPoolBatchIndexKey(poolID uint64) []byte {
	key := make([]byte, 9)
	key[0] = PoolBatchIndexKeyPrefix[0]
	copy(key[1:9], sdk.Uint64ToBigEndian(poolID))
	return key
}
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey) error {
	store := ctx.KVStore(storeKey)

	// old key format v042:
	// PoolByReserveAccIndex: `0x12 | ReserveAcc -> ProtocolBuffer(uint64)`
	// PoolBatchIndex: `0x21 | PoolId -> ProtocolBuffer(uint64)`
	// new key format v043:
	// PoolByReserveAccIndex: `0x12 | ReserveAccLen (1 byte) | ReserveAcc -> ProtocolBuffer(uint64)`
	// PoolBatchIndex: deprecated
	MigratePrefixAddress(store, PoolByReserveAccIndexKeyPrefix)
	DeleteDeprecatedPrefix(store, PoolBatchIndexKeyPrefix)
	return nil
}

// MigratePrefixAddress is a helper function that migrates all keys of format:
// prefix_bytes | address_bytes
// into format:
// prefix_bytes | address_len (1 byte) | address_bytes
func MigratePrefixAddress(store sdk.KVStore, prefixBz []byte) {
	oldStore := prefix.NewStore(store, prefixBz)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		// Set new key on store. Values don't change.
		store.Set(append(prefixBz, address.MustLengthPrefix(oldStoreIter.Key())...), oldStoreIter.Value())
		oldStore.Delete(oldStoreIter.Key())
	}
}

// DeleteDeprecatedPrefix is a helper function that deletes all keys which started the prefix
func DeleteDeprecatedPrefix(store sdk.KVStore, prefixBz []byte) {
	oldStore := prefix.NewStore(store, prefixBz)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		oldStore.Delete(oldStoreIter.Key())
	}
}

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return MigrateStore(ctx, m.keeper.storeKey)
}
