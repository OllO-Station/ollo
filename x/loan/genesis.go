package loan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/loan/keeper"
	"ollo/x/loan/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the loans
	for _, elem := range genState.LoansList {
		k.SetLoans(ctx, elem)
	}

	// Set loans count
	// k.SetLoansCount(ctx, genState.LoansCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.LoansList = k.GetAllLoans(ctx)
	// genesis.LoansCount = 0
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
