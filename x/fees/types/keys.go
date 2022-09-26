package types

const (
	// ModuleName defines the module name
	ModuleName               = "fees"
	FeeCollectorName         = "fee_collector"
	ExternalFeeCollectorName = "non_native_fee_collector"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_fees"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	BaseDenomKey         = []byte("base_denom")
	FeeTokensStorePrefix = []byte("fee_tokens")
)
