package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type GoalVoteHooks struct {
	k         Keeper
	missionID uint64
}

// NewGoalVoteHooks returns a GovHooks that triggers mission completion on voting for a proposal
func (k Keeper) NewGoalVoteHooks(missionID uint64) GoalVoteHooks {
	return GoalVoteHooks{k, missionID}
}

var _ govtypes.GovHooks = GoalVoteHooks{}

// AfterProposalVote completes mission when a vote is cast
func (h GoalVoteHooks) AfterProposalVote(ctx sdk.Context, _ uint64, voterAddr sdk.AccAddress) {
	// TODO: error handling
	_, _ = h.k.CompleteGoal(ctx, h.missionID, voterAddr.String())
}

// AfterProposalSubmission implements GovHooks
func (h GoalVoteHooks) AfterProposalSubmission(_ sdk.Context, _ uint64) {
}

// AfterProposalDeposit implements GovHooks
func (h GoalVoteHooks) AfterProposalDeposit(_ sdk.Context, _ uint64, _ sdk.AccAddress) {
}

// AfterProposalFailedMinDeposit implements GovHooks
func (h GoalVoteHooks) AfterProposalFailedMinDeposit(_ sdk.Context, _ uint64) {
}

// AfterProposalVotingPeriodEnded implements GovHooks
func (h GoalVoteHooks) AfterProposalVotingPeriodEnded(_ sdk.Context, _ uint64) {
}
