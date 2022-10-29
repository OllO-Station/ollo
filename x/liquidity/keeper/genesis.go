package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/liquidity/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	if err := gs.Validate(); err != nil {
		panic(err)
	}
	k.SetParams(ctx, gs.Params)
	k.SetLastPairId(ctx, gs.PrevPairId)
	k.SetLastPoolId(ctx, gs.PrevPoolId)
	for _, pair := range gs.Pairs {
		k.SetPair(ctx, pair)
		k.SetPairIndex(ctx, pair.BaseDenom, pair.QuoteDenom, pair.Id)
		k.SetPairLookupIndex(ctx, pair.BaseDenom, pair.QuoteDenom, pair.Id)
		k.SetPairLookupIndex(ctx, pair.QuoteDenom, pair.BaseDenom, pair.Id)
	}
	for _, pool := range gs.Pools {
		k.SetPool(ctx, pool)
		k.SetPoolByReserveIndex(ctx, pool)
		k.SetPoolsByPairIndex(ctx, pool)
	}
	for _, req := range gs.Requests.Deposits {
		k.SetRequestDeposit(ctx, req)
		k.SetRequestDepositIndex(ctx, req)
	}
	for _, req := range gs.Requests.Withdrawals {
		k.SetRequestWithdraw(ctx, req)
		k.SetRequestWithdrawIndex(ctx, req)
	}
	for _, order := range gs.Requests.Orders {
		k.SetOrder(ctx, order)
		k.SetOrderIndex(ctx, order)
	}
	for _, index := range gs.Requests.MarketMakingOrderIds {
		k.SetMarketMakingOrderId(ctx, index)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:     k.GetParams(ctx),
		PrevPairId: k.GetLastPairId(ctx),
		PrevPoolId: k.GetLastPoolId(ctx),
		Pairs:      k.GetAllPairs(ctx),
		Pools:      k.GetAllPools(ctx),
		Requests: &types.GenesisRequestsState{
			Orders:               k.GetAllOrders(ctx),
			Withdrawals:          k.GetAllRequestWithdraws(ctx),
			Deposits:             k.GetAllRequestDeposits(ctx),
			MarketMakingOrderIds: k.GetAllMarketMakingOrderIdes(ctx),
		},
	}
}
