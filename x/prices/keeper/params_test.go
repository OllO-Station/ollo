package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ollo/testutil/keeper"
	"ollo/x/prices/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PricesKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
