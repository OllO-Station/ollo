package types

import (
	"errors"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default claim genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AirdropSupply: sdk.NewCoin("utest", sdkmath.ZeroInt()),
		ClaimRecords:  []ClaimRecord{},
		Goals:         []Goal{},
		InitialClaim:  InitialClaim{},
		Params:        DefaultParams(),
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// check airdrop supply
	err := gs.AirdropSupply.Validate()
	if err != nil {
		return err
	}

	// check missions
	weightSum := sdk.ZeroDec()
	missionMap := make(map[uint64]Goal)
	for _, mission := range gs.Goals {
		err := mission.Validate()
		if err != nil {
			return err
		}

		weightSum = weightSum.Add(mission.Weight)
		if _, ok := missionMap[mission.Id]; ok {
			return errors.New("duplicated id for mission")
		}
		missionMap[mission.Id] = mission
	}

	// ensure mission weight sum is 1
	if len(gs.Goals) > 0 {
		if !weightSum.Equal(sdk.OneDec()) {
			return errors.New("sum of mission weights must be 1")
		}
	}

	// check initial claim mission exist if enabled
	if gs.InitialClaim.Enabled {
		if _, ok := missionMap[gs.InitialClaim.GoalId]; !ok {
			return errors.New("initial claim mission doesn't exist")
		}
	}

	for _, claimRecord := range gs.ClaimRecords {
		err = claimRecord.Validate()
		if err != nil {
			return err
		}

	}

	err = CheckAirdropSupply(gs.AirdropSupply, missionMap, gs.ClaimRecords)
	if err != nil {
		return err
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
