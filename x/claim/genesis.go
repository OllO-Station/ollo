package claim

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/x/claim/keeper"
	"github.com/ollo-station/ollo/x/claim/types"
)

// InitGenesis initializes the claim module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the claimRecord
	for _, elem := range genState.ClaimRecords {
		k.SetClaimRecord(ctx, elem)
	}
	// Set all the mission
	for _, elem := range genState.Goals {
		k.SetGoal(ctx, elem)
	}

	if err := k.InitializeAirdropSupply(ctx, genState.AirdropSupply); err != nil {
		panic("airdrop supply failed to initialize: " + err.Error())
	}

	k.SetInitialClaim(ctx, genState.InitialClaim)

	k.SetParams(ctx, genState.Params)

	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the claim module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ClaimRecords = k.GetAllClaimRecord(ctx)
	genesis.Goals = k.GetAllGoal(ctx)
	airdropSupply, found := k.GetAirdropSupply(ctx)
	if found {
		genesis.AirdropSupply = airdropSupply
	} else {
		// set to zero coin otherwise
		genesis.AirdropSupply = types.DefaultGenesis().AirdropSupply
	}
	// Get all initialClaim
	initialClaim, found := k.GetInitialClaim(ctx)
	if found {
		genesis.InitialClaim = initialClaim
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
