package keeper

import (
	// "fmt"

	"fmt"

	"github.com/ollo-station/ollo/x/claim/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	airdropSupplyRoute = "airdrop-supply"
	claimRecordRoute   = "claim-record"
)

// // RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	// ir.RegisterRoute(types.ModuleName, airdropSupplyRoute,
	// 	AirdropSupplyInvariant(k))
	// ir.RegisterRoute(types.ModuleName, claimRecordRoute,
	// ClaimRecordInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := AirdropSupplyInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return ClaimRecordInvariant(k)(ctx)
	}
}

// // AirdropSupplyInvariant invariant checks that airdrop supply is equal to the remaining claimable
// // amounts in claim records
func AirdropSupplyInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		missions := k.GetAllGoal(ctx)
		claimRecords := k.GetAllClaimRecord(ctx)
		airdropSupply, _ := k.GetAirdropSupply(ctx)

		missionMap := make(map[uint64]types.Goal)
		for _, mission := range missions {
			missionMap[mission.Id] = mission
		}

		err := types.CheckAirdropSupply(airdropSupply, missionMap, claimRecords)
		if err != nil {
			return err.Error(), true
		}

		return "", false
	}
}

// // ClaimRecordInvariant invariant checks that claim record was claimed but not completed
func ClaimRecordInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		missions := k.GetAllGoal(ctx)
		claimRecords := k.GetAllClaimRecord(ctx)

		for _, claimRecord := range claimRecords {
			for _, mission := range missions {
				if !claimRecord.IsGoalCompleted(mission.Id) &&
					claimRecord.IsGoalClaimed(mission.Id) {

					return fmt.Sprintf("mission %d claimed but not completed", mission.Id), true
				}
			}
		}
		return "", false
	}
}
