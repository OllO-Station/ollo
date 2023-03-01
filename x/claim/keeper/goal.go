package keeper

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ollo-station/ollo/x/claim/types"
)

// SetGoal set a specific mission in the store
func (k Keeper) SetGoal(ctx sdk.Context, goal types.Goal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoalKey))
	b := k.cdc.MustMarshal(&goal)
	store.Set(types.GetGoalIDBytes(goal.Id), b)
}

// GetGoal returns a mission from its id
func (k Keeper) GetGoal(ctx sdk.Context, id uint64) (val types.Goal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoalKey))
	b := store.Get(types.GetGoalIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGoal removes a mission from the store
func (k Keeper) RemoveGoal(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoalKey))
	store.Delete(types.GetGoalIDBytes(id))
}

// GetAllGoal returns all mission
func (k Keeper) GetAllGoal(ctx sdk.Context) (list []types.Goal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoalKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Goal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CompleteGoal saves the completion of the mission. The claim will
// be called automatically if the airdrop start has already been reached.
// If not, it will only save the mission as completed.
func (k Keeper) CompleteGoal(
	ctx sdk.Context,
	missionID uint64,
	address string,
) (claimed math.Int, err error) {
	// retrieve mission
	if _, found := k.GetGoal(ctx, missionID); !found {
		return claimed, errors.Wrapf(types.ErrGoalNotFound, "mission %d not found", missionID)
	}

	// retrieve claim record of the user
	claimRecord, found := k.GetClaimRecord(ctx, address)
	if !found {
		return claimed, errors.Wrapf(types.ErrClaimRecordNotFound, "claim record not found for address %s", address)
	}

	// check if the mission is already completed for the claim record
	if claimRecord.IsGoalCompleted(missionID) {
		return claimed, errors.Wrapf(
			types.ErrGoalCompleted,
			"mission %d completed for address %s",
			missionID,
			address,
		)
	}
	claimRecord.CompletedGoals = append(claimRecord.CompletedGoals, missionID)

	k.SetClaimRecord(ctx, claimRecord)

	err = ctx.EventManager().EmitTypedEvent(&types.EventGoalCompleted{
		GoalID:  missionID,
		Address: address,
	})
	if err != nil {
		return claimed, err
	}

	// try to claim the mission if airdrop start is reached
	airdropStart := k.AirdropStart(ctx)
	if ctx.BlockTime().After(airdropStart) {
		return k.ClaimGoal(ctx, claimRecord, missionID)
	}

	return claimed, nil
}

// ClaimGoal distributes the claimable portion of the airdrop to the user
// the method fails if the mission has already been claimed or not completed
func (k Keeper) ClaimGoal(
	ctx sdk.Context,
	claimRecord types.ClaimRecord,
	missionID uint64,
) (claimed math.Int, err error) {
	airdropSupply, found := k.GetAirdropSupply(ctx)
	if !found {
		return claimed, errors.Wrap(types.ErrAirdropSupplyNotFound, "airdrop supply is not defined")
	}

	// retrieve mission
	mission, found := k.GetGoal(ctx, missionID)
	if !found {
		return claimed, errors.Wrapf(types.ErrGoalNotFound, "mission %d not found", missionID)
	}

	// check if the mission is not completed for the claim record
	if !claimRecord.IsGoalCompleted(missionID) {
		return claimed, errors.Wrapf(
			types.ErrGoalNotCompleted,
			"mission %d is not completed for address %s",
			missionID,
			claimRecord.Address,
		)
	}
	if claimRecord.IsGoalClaimed(missionID) {
		return claimed, errors.Wrapf(
			types.ErrGoalAlreadyClaimed,
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

	return claimed, ctx.EventManager().EmitTypedEvent(&types.EventGoalClaimed{
		GoalID:  missionID,
		Claimer: claimRecord.Address,
	})
}
