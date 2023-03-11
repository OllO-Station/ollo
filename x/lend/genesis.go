package lend

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ollo-station/ollo/x/lend/keeper"
	"github.com/ollo-station/ollo/x/lend/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
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

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LoanList = k.GetAllLoan(ctx)
	genesis.LoanCount = k.GetLoanCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
