package types

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate checks the claimRecord is valid
func (m ClaimRecord) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return err
	}

	if !m.Claimable.IsPositive() {
		return errors.New("claimable amount must be positive")
	}

	missionIDMap := make(map[uint64]struct{})
	for _, elem := range m.CompletedGoals {
		if _, ok := missionIDMap[elem]; ok {
			return fmt.Errorf("duplicated id for completed mission")
		}
		missionIDMap[elem] = struct{}{}
	}

	return nil
}

// IsGoalCompleted checks if the specified mission ID is completed for the claim record
func (m ClaimRecord) IsGoalCompleted(missionID uint64) bool {
	for _, completed := range m.CompletedGoals {
		if completed == missionID {
			return true
		}
	}
	return false
}

// IsGoalClaimed checks if the specified mission ID is claimed for the claim record
func (m ClaimRecord) IsGoalClaimed(missionID uint64) bool {
	for _, claimed := range m.ClaimedGoals {
		if claimed == missionID {
			return true
		}
	}
	return false
}

// ClaimableFromGoal returns the amount claimable for this claim record from the provided mission completion
func (m ClaimRecord) ClaimableFromGoal(mission Goal) sdkmath.Int {
	return mission.Weight.Mul(sdk.NewDecFromInt(m.Claimable)).TruncateInt()
}
