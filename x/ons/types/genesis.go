package types

import (
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:    PortID,

		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
    State: NameState{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in whois
	// nameMap := make(map[string]int)

	// for i, elem := range gs.State.Names {
	// 	if _, ok := nameMap[elem.Name]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Name] = i
	// }
	// for i, elem := range gs.State.NameTags {
	// 	if _, ok := nameMap[elem.Name]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Name] = i
 //  }
	// for i, elem := range gs.State.Threads {
	// 	if _, ok := nameMap[elem.Name]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Name] = i
 //  }
	// for i, elem := range gs.State.ThreadTags {
	// 	if _, ok := nameMap[elem.Tag]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Tag] = i
 //  }
	// for i, elem := range gs.State.ThreadMessage {
	// 	if _, ok := nameMap[elem.Content]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Content] = i
 //  }
	// for i, elem := range gs.State.ThreadMessageTags {
	// 	if _, ok := nameMap[elem.Tag]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Tag] = i
 //  }
	// for i, elem := range gs.State.ActiveLoans {
	// 	if _, ok := nameMap[elem.Name]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Name] = i
 //  }
	// for i, elem := range gs.State.ActionTag {
	// 	if _, ok := nameMap[elem.Tag]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Tag] = i
 //  }
	// for i, elem := range gs.State.BuyOffers {
	// 	if _, ok := nameMap[elem.Name]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Name] = i
 //  }
	// for i, elem := range gs.State.SellOffers {
	// 	if _, ok := nameMap[elem.Name]; ok {
	// 		return fmt.Errorf("duplicated index for whois")
	// 	}
 //    nameMap[elem.Name] = i
  // }
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
