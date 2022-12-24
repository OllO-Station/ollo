package prices_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ollo/testutil/keeper"
	"ollo/testutil/nullify"
	"ollo/x/prices"
	"ollo/x/prices/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PricesKeeper(t)
	prices.InitGenesis(ctx, *k, genesisState)
	got := prices.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
