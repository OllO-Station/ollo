package simulation_test

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abnormal scenarios are not tested here.
// func TestRandomizedGenState(t *testing.T) {
// 	interfaceRegistry := codectypes.NewInterfaceRegistry()
// 	cdc := codec.NewProtoCodec(interfaceRegistry)
// 	s := rand.NewSource(1)
// 	r := rand.New(s)

// 	simState := module.SimulationState{
// 		AppParams:    make(simtypes.AppParams),
// 		Cdc:          cdc,
// 		Rand:         r,
// 		NumBonded:    3,
// 		Accounts:     simtypes.RandomAccounts(r, 3),
// 		InitialStake: 1000,
// 		GenState:     make(map[string]json.RawMessage),
// 	}

// 	simulation.RandomizedGenState(&simState)

// 	var genState types.GenesisState
// 	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &genState)

// 	dec1 := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(36122540)))
// 	dec3 := uint32(5)
// 	dec4 := "cosmos1h292smhhttwy0rl3qr4p6xsvpvxc4v05s6rxtczwq3cs6qc462mqejwy8x"
// 	dec5 := uint32(1347)

// 	require.Equal(t, dec1, genState.Params.PrivatePlanCreationFee)
// 	require.Equal(t, dec3, genState.Params.NextEpochDays)
// 	require.Equal(t, dec4, genState.Params.FarmingFeeCollector)
// 	require.Equal(t, dec5, genState.Params.MaxNumPrivatePlans)
// }

// // TestRandomizedGenState tests abnormal scenarios of applying RandomizedGenState.
// func TestRandomizedGenState1(t *testing.T) {
// 	interfaceRegistry := codectypes.NewInterfaceRegistry()
// 	cdc := codec.NewProtoCodec(interfaceRegistry)

// 	s := rand.NewSource(1)
// 	r := rand.New(s)

// 	// all these tests will panic
// 	tests := []struct {
// 		simState module.SimulationState
// 		panicMsg string
// 	}{
// 		{ // panic => reason: incomplete initialization of the simState
// 			module.SimulationState{}, "invalid memory address or nil pointer dereference"},
// 		{ // panic => reason: incomplete initialization of the simState
// 			module.SimulationState{
// 				AppParams: make(simtypes.AppParams),
// 				Cdc:       cdc,
// 				Rand:      r,
// 			}, "assignment to entry in nil map"},
// 	}

// 	for _, tt := range tests {
// 		require.Panicsf(t, func() { simulation.RandomizedGenState(&tt.simState) }, tt.panicMsg)
// 	}
// }
