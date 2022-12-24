package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/mint/keeper"
	"ollo/x/mint/types"
)

// InitGenesis new mint genesis
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, ak types.AccountKeeper, data *types.GenesisState) {
	keeper.SetMinter(ctx, data.Minter)
	keeper.SetParams(ctx, data.Params)
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Minter = keeper.GetMinter(ctx)
	genesis.Params = keeper.GetParams(ctx)

	return genesis
}
