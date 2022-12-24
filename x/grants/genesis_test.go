package grants_test

import (
	"testing"

	keepertest "ollo/testutil/keeper"
	"ollo/testutil/nullify"
	"ollo/x/grants"
	"ollo/x/grants/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GrantsKeeper(t)
	grants.InitGenesis(ctx, *k, genesisState)
	got := grants.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
