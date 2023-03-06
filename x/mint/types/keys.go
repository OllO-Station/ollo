package types

// MinterKey is the key to use for the keeper store.
var MinterKey = []byte{0x00}

const (
	// module name
	ModuleName = "mint"

	// StoreKey is the default store key for mint
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the minting store.
	QuerierRoute = StoreKey

	LastEpochReductionKey = "last_epoch_reduction"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
