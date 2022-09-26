package loan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/loan/keeper"
	"ollo/x/loan/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the loan
	for _, elem := range genState.LoanList {
		k.SetLoan(ctx, elem)
	}

	// Set loan count
	k.SetLoanCount(ctx, genState.LoanCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LoanList = k.GetAllLoan(ctx)
	genesis.LoanCount = k.GetLoanCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
