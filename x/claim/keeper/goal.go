package keeper

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"

	"ollo/x/claim/types"
)

// SetMission set a specific mission in the store
func (k Keeper) SetMission(ctx sdk.Context, goal types.Goal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	b := k.cdc.MustMarshal(&goal)
	store.Set(types.GetMissionIDBytes(goal.Id), b)
}

// GetMission returns a mission from its id
func (k Keeper) GetMission(ctx sdk.Context, id uint64) (val types.Goal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	b := store.Get(types.GetMissionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMission removes a mission from the store
func (k Keeper) RemoveMission(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	store.Delete(types.GetMissionIDBytes(id))
}

// GetAllMission returns all mission
func (k Keeper) GetAllMission(ctx sdk.Context) (list []types.Goal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Goal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CompleteMission saves the completion of the mission. The claim will
// be called automatically if the airdrop start has already been reached.
// If not, it will only save the mission as completed.
func (k Keeper) CompleteMission(
	ctx sdk.Context,
	missionID uint64,
	address string,
) (claimed math.Int, err error) {
	// retrieve mission
	if _, found := k.GetMission(ctx, missionID); !found {
		return claimed, errors.Wrapf(types.ErrMissionNotFound, "mission %d not found", missionID)
	}

	// retrieve claim record of the user
	claimRecord, found := k.GetClaimRecord(ctx, address)
	if !found {
		return claimed, errors.Wrapf(types.ErrClaimRecordNotFound, "claim record not found for address %s", address)
	}

	// check if the mission is already completed for the claim record
	if claimRecord.IsGoalCompleted(missionID) {
		return claimed, errors.Wrapf(
			types.ErrMissionCompleted,
			"mission %d completed for address %s",
			missionID,
			address,
		)
	}
	claimRecord.CompletedGoals = append(claimRecord.CompletedGoals, missionID)

	k.SetClaimRecord(ctx, claimRecord)

	err = ctx.EventManager().EmitTypedEvent(&types.EventMissionCompleted{
		MissionId: missionID,
		Address:   address,
	})
	if err != nil {
		return claimed, err
	}

	// try to claim the mission if airdrop start is reached
	airdropStart := k.AirdropStart(ctx)
	if ctx.BlockTime().After(airdropStart) {
		return k.ClaimMission(ctx, claimRecord, missionID)
	}

	return claimed, nil
}

// ClaimMission distributes the claimable portion of the airdrop to the user
// the method fails if the mission has already been claimed or not completed
func (k Keeper) ClaimMission(
	ctx sdk.Context,
	claimRecord types.ClaimRecord,
	missionID uint64,
) (claimed math.Int, err error) {
	airdropSupply, found := k.GetAirdropSupply(ctx)
	if !found {
		return claimed, errors.Wrap(types.ErrAirdropSupplyNotFound, "airdrop supply is not defined")
	}

	// retrieve mission
	mission, found := k.GetMission(ctx, missionID)
	if !found {
		return claimed, errors.Wrapf(types.ErrMissionNotFound, "mission %d not found", missionID)
	}

	// check if the mission is not completed for the claim record
	if !claimRecord.IsGoalCompleted(missionID) {
		return claimed, errors.Wrapf(
			types.ErrMissionNotCompleted,
			"mission %d is not completed for address %s",
			missionID,
			claimRecord.Address,
		)
	}
	if claimRecord.IsGoalClaimed(missionID) {
		return claimed, errors.Wrapf(
			types.ErrMissionAlreadyClaimed,
			"mission %d is already claimed for address %s",
			missionID,
			claimRecord.Address,
		)
	}
	claimRecord.ClaimedGoals = append(claimRecord.ClaimedGoals, missionID)

	// calculate claimable from mission weight and claim
	claimableAmount := claimRecord.ClaimableFromGoal(mission)
	claimable := sdk.NewCoins(sdk.NewCoin(airdropSupply.Denom, claimableAmount))

	// calculate claimable after decay factor
	decayInfo := k.DecayInformation(ctx)
	claimable = decayInfo.ApplyDecayFactor(claimable, ctx.BlockTime())

	// check final claimable non-zero
	if claimable.Empty() {
		return claimed, types.ErrNoClaimable
	}

	// decrease airdrop supply
	claimed = claimable.AmountOf(airdropSupply.Denom)
	airdropSupply.Amount = airdropSupply.Amount.Sub(claimed)
	if airdropSupply.Amount.IsNegative() {
		return claimed, errors.ErrInsufficientFunds.Wrap("airdrop supply is lower than total claimable")
	}

	// send claimable to the user
	claimer, err := sdk.AccAddressFromBech32(claimRecord.Address)
	if err != nil {
		return claimed, errors.ErrInvalidAddress.Wrap(err.Error())
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimer, claimable); err != nil {
		return claimed, errors.New("can't send claimable coins %s", 2, err.Error())
	}

	// update store
	k.SetAirdropSupply(ctx, airdropSupply)
	k.SetClaimRecord(ctx, claimRecord)

	return claimed, ctx.EventManager().EmitTypedEvent(&types.EventMissionClaimed{
		MissionId: missionID,
		Claimer:   claimRecord.Address,
	})
}
