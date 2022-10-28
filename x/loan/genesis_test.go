package loan_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ollo/testutil/keeper"
	"ollo/testutil/nullify"
	"ollo/x/loan"
	"ollo/x/loan/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		LoansList: []types.Loans{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LoansCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoanKeeper(t)
	loan.InitGenesis(ctx, *k, genesisState)
	got := loan.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.LoansList, got.LoansList)
	require.Equal(t, genesisState.LoansCount, got.LoansCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
