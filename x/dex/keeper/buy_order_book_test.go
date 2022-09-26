package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "ollo/testutil/keeper"
	"ollo/testutil/nullify"
	"ollo/x/dex/keeper"
	"ollo/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBuyOrderBook(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.BuyOrderBook {
	items := make([]types.BuyOrderBook, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetBuyOrderBook(ctx, items[i])
	}
	return items
}

func TestBuyOrderBookGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBuyOrderBook(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBuyOrderBook(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBuyOrderBookRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBuyOrderBook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBuyOrderBook(ctx,
			item.Index,
		)
		_, found := keeper.GetBuyOrderBook(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestBuyOrderBookGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBuyOrderBook(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBuyOrderBook(ctx)),
	)
}
