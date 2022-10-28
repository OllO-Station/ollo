package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	liquiditytypes "ollo/x/liquidity/types"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	// stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

type StakingKeeper interface {
	GetParams(sdk sdk.Context) stakingtypes.Params
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	// MintCoins is used only for simulation test codes
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// DistrKeeper is the keeper of the distribution store
type DistrKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type GovKeeper interface {
	IterateProposals(ctx sdk.Context, cb func(proposal govv1.Proposal) (stop bool))
	GetVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) (vote govv1.Vote, found bool)
	// Note that this function is only used before the UpgradeHeight defined in app/upgrades/v1.1.0
	IterateAllVotes(ctx sdk.Context, cb func(vote govv1.Vote) (stop bool))
}

// LiquidityKeeper defines the expected interface needed to check the condition.
type LiquidityKeeper interface {
	GetDepositRequestsByDepositor(ctx sdk.Context, depositor sdk.AccAddress) (reqs []liquiditytypes.MsgDeposit)
	GetOrdersByOrderer(ctx sdk.Context, orderer sdk.AccAddress) (orders []liquiditytypes.Order)
}
