package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "ollo/testutil/keeper"
	"ollo/testutil/nullify"
	"ollo/x/loan/keeper"
	"ollo/x/loan/types"
)

func createNLoans(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Loans {
	items := make([]types.Loans, n)
	for i := range items {
		items[i].Id = keeper.AppendLoans(ctx, items[i])
	}
	return items
}

func TestLoansGet(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNLoans(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLoans(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLoansRemove(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNLoans(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLoans(ctx, item.Id)
		_, found := keeper.GetLoans(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLoansGetAll(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNLoans(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLoans(ctx)),
	)
}

func TestLoansCount(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNLoans(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLoansCount(ctx))
}
