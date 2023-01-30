package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"ollo/app"
	"ollo/x/liquidity"
	"ollo/x/liquidity/types"
)

func TestIterateAllBatchMsgs(t *testing.T) {
	simapp, ctx := createTestInput()
	simapp.LiquidityKeeper.SetParams(ctx, types.DefaultParams())

	// define test denom X, Y for Liquidity Pool
	denomX, denomY := types.AlphabeticalDenomPair(DenomX, DenomY)
	denomA, denomB := types.AlphabeticalDenomPair("denomA", "denomB")

	X := sdk.NewInt(1000000000)
	Y := sdk.NewInt(500000000)
	A := sdk.NewInt(500000000)
	B := sdk.NewInt(1000000000)

	addrs := app.AddTestAddrsIncremental(simapp, ctx, 20, sdk.NewInt(10000))
	poolID := app.TestCreatePool(t, simapp, ctx, X, Y, denomX, denomY, addrs[0])
	poolId2 := app.TestCreatePool(t, simapp, ctx, A, B, denomA, denomB, addrs[4])
	batch, found := simapp.LiquidityKeeper.GetPoolBatch(ctx, poolID)
	require.True(t, found)

	app.TestDepositPool(t, simapp, ctx, X.QuoRaw(10), Y, addrs[1:2], poolID, false)
	app.TestDepositPool(t, simapp, ctx, X.QuoRaw(10), Y, addrs[1:2], poolID, false)
	app.TestDepositPool(t, simapp, ctx, X.QuoRaw(10), Y, addrs[1:2], poolID, false)
	app.TestDepositPool(t, simapp, ctx, X, Y.QuoRaw(10), addrs[2:3], poolID, false)
	app.TestDepositPool(t, simapp, ctx, X, Y.QuoRaw(10), addrs[2:3], poolID, false)
	app.TestDepositPool(t, simapp, ctx, X, Y.QuoRaw(10), addrs[2:3], poolID, false)

	// next block
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	liquidity.BeginBlocker(ctx, simapp.LiquidityKeeper)

	app.TestDepositPool(t, simapp, ctx, A, B.QuoRaw(10), addrs[4:5], poolId2, false)
	app.TestWithdrawPool(t, simapp, ctx, sdk.NewInt(1000), addrs[4:5], poolId2, false)
	liquidity.EndBlocker(ctx, simapp.LiquidityKeeper)

	// next block
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	liquidity.BeginBlocker(ctx, simapp.LiquidityKeeper)

	app.TestDepositPool(t, simapp, ctx, A, B.QuoRaw(10), addrs[4:5], poolId2, false)
	app.TestWithdrawPool(t, simapp, ctx, sdk.NewInt(1000), addrs[4:5], poolId2, false)
	liquidity.EndBlocker(ctx, simapp.LiquidityKeeper)

	// next block,
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	// Reinitialize batch messages that were not executed in the previous batch and delete batch messages that were executed or ready to delete.
	liquidity.BeginBlocker(ctx, simapp.LiquidityKeeper)

	app.TestDepositPool(t, simapp, ctx, X.QuoRaw(10), Y, addrs[1:2], poolID, false)
	app.TestDepositPool(t, simapp, ctx, X.QuoRaw(10), Y, addrs[1:2], poolID, false)
	app.TestDepositPool(t, simapp, ctx, X, Y.QuoRaw(10), addrs[2:3], poolID, false)
	app.TestDepositPool(t, simapp, ctx, X, Y.QuoRaw(10), addrs[2:3], poolID, false)
	app.TestWithdrawPool(t, simapp, ctx, sdk.NewInt(50), addrs[1:2], poolID, false)
	app.TestWithdrawPool(t, simapp, ctx, sdk.NewInt(500), addrs[1:2], poolID, false)
	app.TestWithdrawPool(t, simapp, ctx, sdk.NewInt(50), addrs[2:3], poolID, false)
	app.TestWithdrawPool(t, simapp, ctx, sdk.NewInt(500), addrs[2:3], poolID, false)

	depositMsgsRemaining := simapp.LiquidityKeeper.GetAllRemainingPoolBatchDepositMsgStates(ctx, batch)
	require.Equal(t, 0, len(depositMsgsRemaining))

	var depositMsgs []types.DepositMsgState
	simapp.LiquidityKeeper.IterateAllDepositMsgStates(ctx, func(msg types.DepositMsgState) bool {
		depositMsgs = append(depositMsgs, msg)
		return false
	})
	require.Equal(t, 4, len(depositMsgs))

	depositMsgs[0].ToBeDeleted = true
	simapp.LiquidityKeeper.SetPoolBatchDepositMsgStates(ctx, poolID, []types.DepositMsgState{depositMsgs[0]})
	depositMsgsNotToDelete := simapp.LiquidityKeeper.GetAllPoolBatchDepositMsgStatesNotToBeDeleted(ctx, batch)
	require.Equal(t, 3, len(depositMsgsNotToDelete))

	var withdrawMsgs []types.WithdrawMsgState
	simapp.LiquidityKeeper.IterateAllWithdrawMsgStates(ctx, func(msg types.WithdrawMsgState) bool {
		withdrawMsgs = append(withdrawMsgs, msg)
		return false
	})
	withdrawMsgs[0].ToBeDeleted = true
	simapp.LiquidityKeeper.SetPoolBatchWithdrawMsgStates(ctx, poolID, withdrawMsgs[0:1])

	withdrawMsgsNotToDelete := simapp.LiquidityKeeper.GetAllPoolBatchWithdrawMsgStatesNotToBeDeleted(ctx, batch)
	require.Equal(t, 4, len(withdrawMsgs))
	require.Equal(t, 3, len(withdrawMsgsNotToDelete))
	require.NotEqual(t, withdrawMsgsNotToDelete, withdrawMsgs)

	app.TestDepositPool(t, simapp, ctx, A, B.QuoRaw(10), addrs[4:5], poolId2, false)
	app.TestWithdrawPool(t, simapp, ctx, sdk.NewInt(1000), addrs[4:5], poolId2, false)

	depositMsgs = simapp.LiquidityKeeper.GetAllDepositMsgStates(ctx)
	require.Equal(t, 5, len(depositMsgs))
	withdrawMsgs = simapp.LiquidityKeeper.GetAllWithdrawMsgStates(ctx)
	require.Equal(t, 5, len(depositMsgs))

	var depositMsgs2 []types.DepositMsgState
	simapp.LiquidityKeeper.IterateAllDepositMsgStates(ctx, func(msg types.DepositMsgState) bool {
		depositMsgs2 = append(depositMsgs2, msg)
		return false
	})

	var withdrawMsgs2 []types.WithdrawMsgState
	simapp.LiquidityKeeper.IterateAllWithdrawMsgStates(ctx, func(msg types.WithdrawMsgState) bool {
		withdrawMsgs2 = append(withdrawMsgs2, msg)
		return false
	})

	require.Equal(t, 5, len(depositMsgs2))

	require.Equal(t, 5, len(withdrawMsgs2))

	liquidity.EndBlocker(ctx, simapp.LiquidityKeeper)

	depositMsgsRemaining = simapp.LiquidityKeeper.GetAllRemainingPoolBatchDepositMsgStates(ctx, batch)
	require.Equal(t, 0, len(depositMsgsRemaining))

	// next block,
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	// Reinitialize batch messages that were not executed in the previous batch and delete batch messages that were executed or ready to delete.
	liquidity.BeginBlocker(ctx, simapp.LiquidityKeeper)

	var depositMsgs3 []types.DepositMsgState
	simapp.LiquidityKeeper.IterateAllDepositMsgStates(ctx, func(msg types.DepositMsgState) bool {
		depositMsgs3 = append(depositMsgs3, msg)
		return false
	})
	require.Equal(t, 0, len(depositMsgs3))

	var withdrawMsgs3 []types.WithdrawMsgState
	simapp.LiquidityKeeper.IterateAllWithdrawMsgStates(ctx, func(msg types.WithdrawMsgState) bool {
		withdrawMsgs3 = append(withdrawMsgs3, msg)
		return false
	})
	require.Equal(t, 0, len(withdrawMsgs3))

	liquidity.EndBlocker(ctx, simapp.LiquidityKeeper)

	genesis := simapp.LiquidityKeeper.ExportGenesis(ctx)
	simapp.LiquidityKeeper.InitGenesis(ctx, *genesis)
	genesisNew := simapp.LiquidityKeeper.ExportGenesis(ctx)
	require.Equal(t, genesis, genesisNew)

	simapp.LiquidityKeeper.DeletePoolBatch(ctx, batch)
	batch, found = simapp.LiquidityKeeper.GetPoolBatch(ctx, batch.PoolId)
	require.Equal(t, types.PoolBatch{}, batch)
	require.False(t, found)
}
