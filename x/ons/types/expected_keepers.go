package types

import (
	// "context"
	"time"

	// db "github.com/cometbft/cometbft-db"
	// codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	// authz "github.com/cosmos/cosmos-sdk/x/authz"
	// disttypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	// "github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/cosmos/cosmos-sdk/x/group"
	// "github.com/ollo-station/ollo/x/nft/exported"
	// nftexported "github.com/ollo-station/ollo/x/nft/exported"
	// "github.com/ollo-station/ollo/x/token/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	// banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	// clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
)

type GroupKeeper interface {
	Tally(ctx sdk.Context, p group.Proposal, groupID uint64) (group.TallyResult, error)
	TallyProposalsAtVPEnd(ctx sdk.Context) error
	// Methods imported from group should be defined here
}

type GovKeeper interface {
	// Methods imported from gov should be defined here
	GetGovernanceAccount(ctx sdk.Context) authtypes.ModuleAccountI
	InsertActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time)
	RemoveFromActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time)
	// IteratorActiveProposalQueue(ctx sdk.Context, endTime time.Time) db.Iterator
	// IteratorInactiveProposalQueue(ctx sdk.Context, endTime time.Time) db.Iterator
	RemoveFromInactiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time)
	InsertInactiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time)
}

// The full set of methods of the epoching native keeper
type EpochingKeeper interface {
	// ActionStoreKey(epochNumber int64, actionID uint64) []byte
	// GetNewActionID(ctx sdk.Context) uint64
	// GetEpochMsg(ctx sdk.Context, epochNumber int64, actionID uint64) sdk.Msg
	// GetEpochActions(ctx sdk.Context) []sdk.Msg
	// GetNextEpochTime(ctx sdk.Context, epochInterval int64) time.Time
	// GetNextEpochHeight(ctx sdk.Context, epochInterval int64) int64
	// IncreaseEpochNumber(ctx sdk.Context)
	// RestoreEpochAction(ctx sdk.Context, epochNumber int64, action *codectypes.Any)
	// QueueMsgForEpoch(ctx sdk.Context, epochNumber int64, msg sdk.Msg)
	// GetEpochNumber(ctx sdk.Context) int64
	// SetEpochNumber(ctx sdk.Context, epochNumber int64)
	// GetEpochActionByIterator(iterator db.Iterator) sdk.Msg
	// DequeueEpochActions(ctx sdk.Context)
	// GetEpochActionsIterator(ctx sdk.Context) db.Iterator
	// DeleteByKey(ctx sdk.Context, key []byte)
}

type NFTKeeper interface {
	// GetNFT(ctx sdk.Context, denom string, id string) (nftexported.NFT, error)
	// GetNFTs(ctx sdk.Context, denom string) (nfts []exported.NFT, err error)
	// HasNFT(ctx sdk.Context, denomID, tokenID string) bool
	// Authorize(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) error
	// SaveNFT(
	// 	ctx sdk.Context,
	// 	tokenID,
	// 	tokenNm,
	// 	tokenURI,
	// 	tokenUriHash,
	// 	tokenData string,
	// 	receiver sdk.AccAddress,
	// ) error
	// UpdateNFT(
	// 	ctx sdk.Context,
	// 	denomID,
	// 	tokenID,
	// 	tokenNm,
	// 	tokenURI,
	// 	tokenURIHash,
	// 	tokenData string,
	// 	owner sdk.AccAddress,
	// ) error
	// RemoveNFT(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) (nft exported.NFT, err error)
	// TransferOwnership(
	// 	ctx sdk.Context,
	// 	denomId, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData string,
	// 	owner sdk.AccAddress) error
	// SendNFT(ctx sdk.Context, denom string, id string, from sdk.AccAddress, to sdk.AccAddress) error
	// MintNft(ctx sdk.Context, denom string, id string) error
	NFTBurn(ctx sdk.Context, denom, id string, from sdk.AccAddress) error
}

type TokenKeeper interface {
	// Methods imported from token should be defined here
	// IssueToken(
	// 	ctx sdk.Context,
	// 	symbol, name, minUnit string,
	// 	scale uint32,
	// 	initialSupply, maxSupply uint64,
	// 	mintable bool,
	// 	owner sdk.AccAddress,
	// )
	// EditToken(
	// 	ctx sdk.Context,
	// 	symbol, name, maxSupply string,
	// 	mintable types.Bool,
	// 	owner sdk.AccAddress,
	// )
	// TransferTokenOwner(
	// 	sdk sdk.Context,
	// 	symbol string,
	// 	srcOwner, dstOwner sdk.AccAddress,
	// )
	// MintToken(
	// 	ctx sdk.Context,
	// 	symbol string,
	// 	amount uint64,
	// 	recipient, owner sdk.AccAddress,
	// )
	// BurnToken(
	// 	ctx sdk.Context,
	// 	symbol string,
	// 	amount uint64,
	// 	owner sdk.AccAddress,
	// )
}

type StakingKeeper interface {
	// BondDenom(ctx sdk.Context) (res string)
	// // GetValidator get a single validator
	// GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	// // GetBondedValidatorsByPower get the current group of bonded validators sorted by power-rank
	// GetBondedValidatorsByPower(ctx sdk.Context) []stakingtypes.Validator
	// // GetAllDelegatorDelegations return all delegations for a delegator
	// GetAllDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress) []stakingtypes.Delegation
	// // GetDelegation return a specific delegation
	// GetDelegation(ctx sdk.Context,
	// 	delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation stakingtypes.Delegation, found bool)
	// // HasReceivingRedelegation check if validator is receiving a redelegation
	// HasReceivingRedelegation(ctx sdk.Context,
	// 	delAddr sdk.AccAddress, valDstAddr sdk.ValAddress) bool
	// StakingTokenSupply(ctx sdk.Context) sdk.Int
	// BondedRatio(ctx sdk.Context) sdk.Dec
}

//	type DistrKeeper interface {
//		FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
//		DelegationRewards(c context.Context, req *disttypes.QueryDelegationRewardsRequest) (*disttypes.QueryDelegationRewardsResponse, error)
//	}
type AuthzKeeper interface {
	// DispatchActions(ctx sdk.Context, grantee sdk.AccAddress, msgs []sdk.Msg) error
	// DeleteGrant(ctx sdk.Context, grantee, granter sdk.AccAddress, msgType string) error
	// GetAuthorization(ctx sdk.Context, grantee, granter sdk.AccAddress, msgType string) (authz.Authorization, error)
	// // IterateGrants(ctx sdk.Context, handler func(granter, grantee sdk.AccAddress, grant authz.Grant)) bool
	// GetAuthorizations(ctx sdk.Context, granter, grantee sdk.AccAddress) ([]*authz.Grant, error)
	// SaveGrant(
	// 	ctx sdk.Context,
	// 	grantee, granter sdk.AccAddress,
	// 	authorization authz.Authorization,
	// 	expiration *time.Time) error
	// DequeueAndDeleteExpiredGrants(ctx sdk.Context) error
}

type FeeGrantKeeper interface {
	// GrantAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress, allowance feegrant.FeeAllowanceI) error
	// UpdateAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress, allowance feegrant.FeeAllowanceI) error
	// GetAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress) (feegrant.FeeAllowanceI, error)
	// // IterateAllFeeAllowances(ctx sdk.Context, cb func(grant feegrant.Grant) bool) error
	// UseGrantedFees(ctx sdk.Context, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) error
	// Methods imported from feegrant should be defined here
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	// IsSendEnabledCoins(ctx sdk.Context, coins ...sdk.Coin) error
	// BlockedAddr(addr sdk.AccAddress) bool
	// SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	// LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	// BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	// SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	// GetSupply(ctx sdk.Context, denom string) sdk.Coin
	// SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	// InputOutputCoins(ctx sdk.Context, inputs []banktypes.Input, outputs []banktypes.Output) error
	// MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	// SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
	// GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	// Methods imported from bank should be defined here
}
