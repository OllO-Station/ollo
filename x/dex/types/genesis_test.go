package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"ollo/x/dex/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PortId: types.PortID,
				SellOrderBookList: []types.SellOrderBook{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				BuyOrderBookList: []types.BuyOrderBook{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				DenomTraceList: []types.DenomTrace{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated sellOrderBook",
			genState: &types.GenesisState{
				SellOrderBookList: []types.SellOrderBook{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated buyOrderBook",
			genState: &types.GenesisState{
				BuyOrderBookList: []types.BuyOrderBook{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated denomTrace",
			genState: &types.GenesisState{
				DenomTraceList: []types.DenomTrace{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
