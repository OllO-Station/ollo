package types

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CheckAirdropSupply(airdropSupply sdk.Coin, missionMap map[uint64]Goal, claimRecords []ClaimRecord) error {
	claimSum := sdkmath.ZeroInt()
	claimRecordMap := make(map[string]struct{})

	for _, claimRecord := range claimRecords {

		// check claim record completed missions
		claimable := claimRecord.Claimable
		for _, completedGoal := range claimRecord.CompletedGoals {
			mission, ok := missionMap[completedGoal]
			if !ok {
				return fmt.Errorf("address %s completed a non existing mission %d",
					claimRecord.Address,
					completedGoal,
				)
			}

			// reduce claimable with already claimed funds
			claimable = claimable.Sub(claimRecord.ClaimableFromGoal(mission))
		}

		claimSum = claimSum.Add(claimable)
		if _, ok := claimRecordMap[claimRecord.Address]; ok {
			return errors.New("duplicated address for claim record")
		}
		claimRecordMap[claimRecord.Address] = struct{}{}
	}

	// verify airdropSupply == sum of claimRecords
	if !airdropSupply.Amount.Equal(claimSum) {
		return fmt.Errorf("airdrop supply amount %v not equal to sum of claimable amounts %v",
			airdropSupply.Amount,
			claimSum,
		)
	}

	return nil
}
