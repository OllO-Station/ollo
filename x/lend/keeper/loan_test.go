package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/ollo-station/ollo/testutil/keeper"
	"github.com/ollo-station/ollo/testutil/nullify"
	"github.com/ollo-station/ollo/x/lend/keeper"
	"github.com/ollo-station/ollo/x/lend/types"
	"github.com/stretchr/testify/require"
)

func createNLoan(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Loan {
	items := make([]types.Loan, n)
	for i := range items {
		items[i].Id = keeper.AppendLoan(ctx, items[i])
	}
	return items
}

func TestLoanGet(t *testing.T) {
	keeper, ctx := keepertest.LendKeeper(t)
	items := createNLoan(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLoan(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLoanRemove(t *testing.T) {
	keeper, ctx := keepertest.LendKeeper(t)
	items := createNLoan(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLoan(ctx, item.Id)
		_, found := keeper.GetLoan(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLoanGetAll(t *testing.T) {
	keeper, ctx := keepertest.LendKeeper(t)
	items := createNLoan(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLoan(ctx)),
	)
}

func TestLoanCount(t *testing.T) {
	keeper, ctx := keepertest.LendKeeper(t)
	items := createNLoan(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLoanCount(ctx))
}
