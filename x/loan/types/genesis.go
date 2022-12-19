package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:    PortID,
		LoansList: []Loans{},
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
	// Check for duplicated ID in loans
	loansIdMap := make(map[uint64]bool)
	loansCount := gs.GetLoansCount()
	for _, elem := range gs.LoansList {
		if _, ok := loansIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for loans")
		}
		if elem.Id >= loansCount {
			return fmt.Errorf("loans id should be lower or equal than the last id")
		}
		loansIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
