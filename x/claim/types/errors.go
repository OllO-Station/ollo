package types

// DONTCOVER

import (
	// "errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/claim module sentinel sdkerrors
var (
	ErrGoalNotFound           = sdkerrors.Register(ModuleName, 2, "mission not found")
	ErrClaimRecordNotFound    = sdkerrors.Register(ModuleName, 3, "claim record not found")
	ErrGoalCompleted          = sdkerrors.Register(ModuleName, 4, "mission already completed")
	ErrAirdropSupplyNotFound  = sdkerrors.Register(ModuleName, 5, "airdrop supply not found")
	ErrInitialClaimNotFound   = sdkerrors.Register(ModuleName, 6, "initial claim information not found")
	ErrInitialClaimNotEnabled = sdkerrors.Register(ModuleName, 7, "initial claim not enabled")
	ErrGoalCompleteFailure    = sdkerrors.Register(ModuleName, 8, "mission failed to complete")
	ErrNoClaimable            = sdkerrors.Register(ModuleName, 9, "no amount to be claimed")
	ErrGoalNotCompleted       = sdkerrors.Register(ModuleName, 10, "mission not completed yet")
	ErrGoalAlreadyClaimed     = sdkerrors.Register(ModuleName, 11, "mission already claimed")
	ErrAirdropStartNotReached = sdkerrors.Register(ModuleName, 12, "airdrop start has not been reached yet")
)
