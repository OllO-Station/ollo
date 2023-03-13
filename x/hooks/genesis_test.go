package hooks_test

import (
	"testing"

	keepertest "github.com/ollo-station/ollo/testutil/keeper"
	"github.com/ollo-station/ollo/testutil/nullify"
	"github.com/ollo-station/ollo/x/hooks"
	"github.com/ollo-station/ollo/x/hooks/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HooksKeeper(t)
	hooks.InitGenesis(ctx, *k, genesisState)
	got := hooks.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
