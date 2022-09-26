package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:            PortID,
		SellOrderBookList: []SellOrderBook{},
		BuyOrderBookList:  []BuyOrderBook{},
		DenomTraceList:    []DenomTrace{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in sellOrderBook
	sellOrderBookIndexMap := make(map[string]struct{})

	for _, elem := range gs.SellOrderBookList {
		index := string(SellOrderBookKey(elem.Index))
		if _, ok := sellOrderBookIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for sellOrderBook")
		}
		sellOrderBookIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in buyOrderBook
	buyOrderBookIndexMap := make(map[string]struct{})

	for _, elem := range gs.BuyOrderBookList {
		index := string(BuyOrderBookKey(elem.Index))
		if _, ok := buyOrderBookIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for buyOrderBook")
		}
		buyOrderBookIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in denomTrace
	denomTraceIndexMap := make(map[string]struct{})

	for _, elem := range gs.DenomTraceList {
		index := string(DenomTraceKey(elem.Index))
		if _, ok := denomTraceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for denomTrace")
		}
		denomTraceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
