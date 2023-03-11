package lend_test

import (
	"testing"

	keepertest "github.com/ollo-station/ollo/testutil/keeper"
	"github.com/ollo-station/ollo/testutil/nullify"
	"github.com/ollo-station/ollo/x/lend"
	"github.com/ollo-station/ollo/x/lend/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LoanList: []types.Loan{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LoanCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LendKeeper(t)
	lend.InitGenesis(ctx, *k, genesisState)
	got := lend.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LoanList, got.LoanList)
	require.Equal(t, genesisState.LoanCount, got.LoanCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
