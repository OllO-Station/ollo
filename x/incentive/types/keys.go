package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	address "github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName   = "incentive"
	RouterKey    = ModuleName
	StoreKey     = ModuleName
	QuerierRoute = ModuleName
)

var (
	MarketMakerKeyPrefix              = []byte{0xc0}
	MarketMakerIndexByPairIdKeyPrefix = []byte{0xc1}
	DepositKeyPrefix                  = []byte{0xc2}
	IncentiveKeyPrefix                = []byte{0xc5}
)

func GetDepositKey(mmAddr sdk.AccAddress, pairId uint64) []byte {
	res := append(DepositKeyPrefix, address.MustLengthPrefix(mmAddr)...)
	return append(res, sdk.Uint64ToBigEndian(pairId)...)
}

func GetIncentiveKey(mmAddr sdk.AccAddress) []byte {
	return append(IncentiveKeyPrefix, mmAddr...)
}

func GetMarketMakerKey(mmAddr sdk.AccAddress, pairId uint64) []byte {
	res := append(MarketMakerKeyPrefix, address.MustLengthPrefix(mmAddr)...)
	return append(res, sdk.Uint64ToBigEndian(pairId)...)
}

func GetMarketMakerByAddrPrefix(mmAddr sdk.AccAddress) []byte {
	return append(MarketMakerKeyPrefix, address.MustLengthPrefix(mmAddr)...)
}

func GetMarketMakerByPairIdPrefix(pairId uint64) []byte {
	return append(MarketMakerIndexByPairIdKeyPrefix, sdk.Uint64ToBigEndian(pairId)...)
}

func ParseMarketMakerIndexByPairIdKey(key []byte) (pairId uint64, mmAddr sdk.AccAddress) {
	if !bytes.HasPrefix(key, MarketMakerIndexByPairIdKeyPrefix) {
		panic("key does not have proper prefix")
	}
	pairId = sdk.BigEndianToUint64(key[1:9])
	mmAddr = key[9:]
	return
}
func ParseDepositKey(key []byte) (mmAddr sdk.AccAddress, pairId uint64) {
	if !bytes.HasPrefix(key, DepositKeyPrefix) {
		panic("key does not have proper prefix")
	}
	addrLen := key[1]
	mmAddr = key[2 : 2+addrLen]
	pairId = sdk.BigEndianToUint64(key[2+addrLen:])
	return
}
