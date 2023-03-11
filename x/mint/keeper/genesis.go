package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/ollo-station/ollo/x/mint/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	if data == nil {
		panic("invalid genesis state for mint module")
	}
	data.Minter.AnnualProvisions = *data.Params.GenesisEpochProvisions
	k.SetMinter(ctx, data.Minter)
	k.SetParams(ctx, data.Params)
	k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	// if !k.accountKeeper.HasAccount(ctx, k.accountKeeper.GetModuleAddress(types.FundedAccountName)) {
	//     totalVestingFund := sdk.NewCoin(data.Params.MintDenom, sdk.NewInt(0))
	//     k.bankKeeper.AddSupplyOffset(ctx, data.Params.MintDenom, sdk.NewInt(0))
	// }
	k.SetEpochLastReduction(ctx, uint64(data.LastEpochReduction))
}

func (k Keeper) ExportGenesis(c sdk.Context) *types.GenesisState {
	m := k.GetMinter(c)
	p := k.GetParams(c)
	if p.FundedAddresses == nil {
		p.FundedAddresses = make([]types.WeightedAddress, 0)
	}
	lastEpochHalving := k.GetEpochLastReduction(c)
	return types.NewGenesisState(m, p, int64(lastEpochHalving))
}
