package dex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/dex/keeper"
	"ollo/x/dex/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the sellOrderBook
	for _, elem := range genState.SellOrderBookList {
		k.SetSellOrderBook(ctx, elem)
	}
	// Set all the buyOrderBook
	for _, elem := range genState.BuyOrderBookList {
		k.SetBuyOrderBook(ctx, elem)
	}
	// Set all the denomTrace
	for _, elem := range genState.DenomTraceList {
		k.SetDenomTrace(ctx, elem)
	}
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

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.SellOrderBookList = k.GetAllSellOrderBook(ctx)
	genesis.BuyOrderBookList = k.GetAllBuyOrderBook(ctx)
	genesis.DenomTraceList = k.GetAllDenomTrace(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
