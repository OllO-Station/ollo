package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SellOrderBookKeyPrefix is the prefix to retrieve all SellOrderBook
	SellOrderBookKeyPrefix = "SellOrderBook/value/"
)

// SellOrderBookKey returns the store key to retrieve a SellOrderBook from the index fields
func SellOrderBookKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
