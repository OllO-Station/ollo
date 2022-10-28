package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/claim/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	for _, a := range genState.Params.Airdrops {
		_, found := k.GetAirdrop(ctx, a.Id)
		if found {
			panic("airdrop already exists")
		}
		k.SetAirdrop(ctx, a)
	}

	for _, r := range genState.Params.ClaimRecords {
		k.SetClaimRecord(ctx, r)
	}
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	airdrops := k.GetAllAirdrops(ctx)

	records := []types.ClaimRecord{}
	for _, a := range airdrops {
		records = append(records, k.GetAllClaimRecordsByAirdropId(ctx, a.Id)...)
	}

	return &types.GenesisState{
		Params: types.Params{
			Airdrops:     airdrops,
			ClaimRecords: records,
		},
	}
}
