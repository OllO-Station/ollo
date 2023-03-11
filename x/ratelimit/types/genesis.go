package types

func (gs GenesisState) Validate() error {
	if e := gs.Params.Validate(); e != nil {
		return e
	}
	return nil
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}
