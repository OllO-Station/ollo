package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "claim"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_claim"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	GoalKey = "Goal-value-"
)

const (
	AirdropSupplyKey = "AirdropSupply-value-"
)

const (
	InitialClaimKey = "InitialClaim-value-"
)

var _ binary.ByteOrder

const (
	// ClaimRecordKeyPrefix is the prefix to retrieve all ClaimRecord
	ClaimRecordKeyPrefix = "ClaimRecord/value/"
)

// ClaimRecordKey returns the store key to retrieve a ClaimRecord from the index fields
func ClaimRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
