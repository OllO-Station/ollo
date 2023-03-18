package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govv1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	liquiditytypes "github.com/ollo-station/ollo/x/liquidity/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type (
	// AccountKeeper defines the expected account keeper (noalias)
	AccountKeeper interface {
		GetModuleAddress(name string) sdk.AccAddress
		GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
		GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	}
	// BankKeeper defines the expected bank keeper (noalias)
	BankKeeper interface {
		SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
		BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
		MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
		GetSupply(ctx sdk.Context, denom string) sdk.Coin
		SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
		SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
		SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
		SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	}
	// AccountKeeper defines the expected account keeper (noalias)
	StakingKeeper interface {
		Validator(ctx sdk.Context, addr sdk.ValAddress) stakingtypes.ValidatorI
		ValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) stakingtypes.ValidatorI
		GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
		GetAllValidators(ctx sdk.Context) (validators []stakingtypes.Validator)
		GetBondedValidatorsByPower(ctx sdk.Context) []stakingtypes.Validator
		GetLastTotalPower(ctx sdk.Context) sdk.Int
		GetLastValidatorPower(ctx sdk.Context, addr sdk.ValAddress) int64

		Delegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) stakingtypes.DelegationI
		GetDelegation(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			valAddr sdk.ValAddress,
		) (delegation stakingtypes.Delegation, found bool)
		IterateDelegations(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			fn func(index int64, delegation stakingtypes.DelegationI) (stop bool),
		)
		Delegate(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			bondAmt sdk.Int,
			tokenSrc stakingtypes.BondStatus,
			validator stakingtypes.Validator,
			subtractAmt bool,
		) (newshares sdk.Dec, err error)

		BondDenom(ctx sdk.Context) (res string)
		Unbond(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			valAddr sdk.ValAddress,
			shares sdk.Dec,
		) (amount sdk.Int, err error)
		UnbondingTime(ctx sdk.Context) (res time.Duration)
		SetUnbondingDelegationEntry(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			valAddr sdk.ValAddress,
			creationHeight int64,
			minTime time.Time,
			balance sdk.Int,
		) stakingtypes.UnbondingDelegation
		InsertUBDQueue(
			ctx sdk.Context,
			ubd stakingtypes.UnbondingDelegation,
			completionTime time.Time,
		)
		ValidateUnbondAmount(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			valAddr sdk.ValAddress,
			amt sdk.Int,
		) (shares sdk.Dec, err error)
		GetAllUnbondingDelegations(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
		) []stakingtypes.UnbondingDelegation
		BeginRedelegation(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			valAddrFrom sdk.ValAddress,
			valAddrTo sdk.ValAddress,
			sharesAmount sdk.Dec,
		) (completionTime time.Time, err error)
		GetAllRedelegations(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			valAddrFrom sdk.ValAddress,
			valAddrTo sdk.ValAddress,
		) []stakingtypes.Redelegation
		HasReceivingRedelegation(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			valAddr sdk.ValAddress,
		) bool
		BlockValidatorUpdates(ctx sdk.Context) []abci.ValidatorUpdate
		HasMaxUnbondingDelegationEntries(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			valAddr sdk.ValAddress,
		) bool
	}
	// DistributionKeeper defines the expected distribution keeper (noalias)
	DistrKeeper interface {
		IncrementValidatorPeriod(ctx sdk.Context, val stakingtypes.ValidatorI) uint64
		CalculateDelegationRewards(
			ctx sdk.Context,
			val stakingtypes.ValidatorI,
			del stakingtypes.DelegationI,
			endingPeriod uint64,
		) (rewards sdk.DecCoins)
		WithdrawDelegationRewards(
			ctx sdk.Context,
			delAddr sdk.AccAddress,
			valAddr sdk.ValAddress,
		) (rewards sdk.DecCoins, err error)
	}

	// GovKeeper defines the expected gov keeper (noalias)
	GovKeeper interface {
		Tally(ctx sdk.Context, proposal govv1types.Proposal) (passing bool, tallyResults govv1types.TallyResult)
		AddVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress, opts govv1types.WeightedVoteOptions) error
		GetVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) (vote govv1types.Vote, found bool)
		GetVotes(ctx sdk.Context, proposalID uint64) (votes govv1types.Votes)
		GetProposal(ctx sdk.Context, proposalID uint64) (proposal govv1types.Proposal, found bool)
		GetProposals(ctx sdk.Context) (proposals govv1types.Proposals)
		SetProposal(ctx sdk.Context, proposal govv1types.Proposal)
	}

	// LiquidityKeeper defines the expected liquidity keeper (noalias)
	LiquidityKeeper interface {
		GetPool(ctx sdk.Context, id uint64) (pool liquiditytypes.Pool, found bool)
		GetPoolCoinSupply(ctx sdk.Context, pool liquiditytypes.Pool) sdk.Int
		GetPoolBalances(ctx sdk.Context, pool liquiditytypes.Pool) (rx sdk.Coin, ry sdk.Coin)
		IterateAllPools(ctx sdk.Context, cb func(pool liquiditytypes.Pool) (stop bool, err error)) error
		GetPair(ctx sdk.Context, id uint64) (pair liquiditytypes.Pair, found bool)
	}

	// SlashingKeeper defines the expected slashing keeper (noalias)
	SlashingKeeper interface {
		IsTombstoned(ctx sdk.Context, consAddr sdk.ConsAddress) bool
	}

	// StakingHooks defines the expected staking hooks (noalias)
	StakingHooks interface {
		AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress)
		AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress)
		BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
		BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
		AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
		BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec)
	}
)
