package types

// NewGenesisState creates a new GenesisState object

// DefaultGenesis creates a default GenesisState object
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Minter: DefaultInitialMinter(),
		Params: DefaultParams(),
	}
}

// Validate validates the provided genesis state to ensure the
// expected invariants holds.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	return gs.Minter.Validate()
}

func NewGenesisState(m Minter, p Params, reductionEpoch int64) *GenesisState {
	return &GenesisState{
		Minter:             m,
		Params:             p,
		LastEpochReduction: reductionEpoch,
	}
}

func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}
	if err := data.Minter.Validate(); err != nil {
		return err
	}
	return nil
}
