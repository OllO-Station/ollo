package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "market"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_market"

	// Version defines the current version the IBC module supports
	Version = "market-1"

	// PortID is the default port id that module binds to
	PortID = "market"

	// QueryRoute is the querier route for the market store.
	QuerierRoute string = ModuleName
)

var (
	PrefixNftListingId         = []byte{0x01}
	PrefixNftListingOwner      = []byte{0x02}
	PrefixNftListingsCount     = []byte{0x03}
	PrefixNftListingNFTID      = []byte{0x04}
	PrefixNftListingPriceDenom = []byte{0x05}
	PrefixNftAuctionId         = []byte{0x06}
	PrefixNftAuctionOwner      = []byte{0x07}
	PrefixNftAuctionNFTID      = []byte{0x08}
	PrefixNftAuctionPriceDenom = []byte{0x09}
	PrefixNextNftAuctionNumber = []byte{0x10}
	PrefixBidByNftAuctionId    = []byte{0x11}
	PrefixBidByNftBidder       = []byte{0x12}
	PrefixInactiveNftAuction   = []byte{0x13}
	PrefixActiveNftAuction     = []byte{0x14}
)

const (
	MinListingIdLength   = 4
	MaxListingIdLength   = 64
	MaxSplits            = 5
	MaxWhitelistAccounts = 10
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("market-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func KeyNftListingIdPrefix(id string) []byte {
	return append(PrefixNftListingId, []byte(id)...)
}

func KeyNftListingOwnerPrefix(owner sdk.AccAddress, id string) []byte {
	return append(append(PrefixNftListingOwner, owner.Bytes()...), []byte(id)...)
}

func KeyNftListingNFTIDPrefix(nftId string) []byte {
	return append(PrefixNftListingNFTID, []byte(nftId)...)
}

func KeyNftListingPriceDenomPrefix(priceDenom, id string) []byte {
	return append(append(PrefixNftListingPriceDenom, []byte(priceDenom)...), []byte(id)...)
}

func KeyNftAuctionIdPrefix(id uint64) []byte {
	return append(PrefixNftAuctionId, sdk.Uint64ToBigEndian(id)...)
}

func KeyNftAuctionOwnerPrefix(owner sdk.AccAddress, id uint64) []byte {
	return append(append(PrefixNftAuctionOwner, owner.Bytes()...), sdk.Uint64ToBigEndian(id)...)
}

func KeyNftAuctionNFTIDPrefix(nftId string) []byte {
	return append(PrefixNftAuctionNFTID, []byte(nftId)...)
}

func KeyNftAuctionPriceDenomPrefix(priceDenom string, id uint64) []byte {
	return append(append(PrefixNftAuctionPriceDenom, []byte(priceDenom)...), sdk.Uint64ToBigEndian(id)...)
}
func KeyBidPrefix(id uint64) []byte {
	return append(PrefixBidByNftAuctionId, sdk.Uint64ToBigEndian(id)...)
}
func KeyInActiveNftAuctionPrefix(id uint64) []byte {
	return append(PrefixInactiveNftAuction, sdk.Uint64ToBigEndian(id)...)
}

func KeyActiveNftAuctionPrefix(id uint64) []byte {
	return append(PrefixActiveNftAuction, sdk.Uint64ToBigEndian(id)...)
}
