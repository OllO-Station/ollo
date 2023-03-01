package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"

	// "errors"
	"github.com/ollo-station/ollo/x/claim/types"
)

// Claim claims the Airdrop by the mission id if available and reach the airdrop start time
func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve claim record of the user
	claimRecord, found := k.GetClaimRecord(ctx, msg.Claimer)
	if !found {
		return &types.MsgClaimResponse{}, errors.Wrapf(
			types.ErrClaimRecordNotFound,
			"claim record not found for address %s",
			msg.Claimer,
		)
	}

	// check if the claim is an initial claim
	initialClaim, found := k.GetInitialClaim(ctx)
	if found {
		if initialClaim.GoalId == msg.GoalId {
			if !initialClaim.Enabled {
				return nil, types.ErrInitialClaimNotEnabled
			}
			// if is an initial claim, automatically add to completed missions
			// the `ClaimGoal` will update the claim record later
			claimRecord.CompletedGoals = append(claimRecord.CompletedGoals, msg.GoalId)
		}
	}

	// check if airdrop start time already reached
	airdropStart := k.AirdropStart(ctx)
	if ctx.BlockTime().Before(airdropStart) {
		return &types.MsgClaimResponse{}, errors.Wrapf(
			types.ErrAirdropStartNotReached,
			"airdrop start not reached: %s",
			airdropStart.String(),
		)
	}
	claimed, err := k.ClaimGoal(ctx, claimRecord, msg.GoalId)
	if err != nil {
		return nil, err
	}

	return &types.MsgClaimResponse{
		Claimed: claimed,
	}, nil
}
