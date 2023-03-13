package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type DistrKeeper interface {
	// Methods imported from distr should be defined here
}

type StakingKeeper interface {
	// Methods imported from staking should be defined here
}

type EpochingKeeper interface {
	// Methods imported from epoching should be defined here
}

type MintKeeper interface {
	// Methods imported from mint should be defined here
}

type GovKeeper interface {
	// Methods imported from gov should be defined here
}

type LiquidityKeeper interface {
	// Methods imported from liquidity should be defined here
}

type LendKeeper interface {
	// Methods imported from lend should be defined here
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
