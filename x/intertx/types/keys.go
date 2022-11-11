package types

const (
	// ModuleName defines the module name
	ModuleName = "intertx"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_intertx"

    // Version defines the current version the IBC module supports
Version = "intertx-1"

// PortID is the default port id that module binds to
PortID = "intertx"
	QuerierRoute = ModuleName
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("intertx-port-")
)

func KeyPrefix(p string) []byte {
    return []byte(p)
}
