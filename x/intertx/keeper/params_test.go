package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ollo/testutil/keeper"
	"ollo/x/intertx/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IntertxKeeper(t)
	params := types.DefaultParams()

	// k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
