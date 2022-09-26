package dex

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"ollo/testutil/sample"
	dexsimulation "ollo/x/dex/simulation"
	"ollo/x/dex/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = dexsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCancelSellOrder = "op_weight_msg_cancel_sell_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelSellOrder int = 100

	opWeightMsgCancelBuyOrder = "op_weight_msg_cancel_buy_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelBuyOrder int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dexGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dexGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCancelSellOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelSellOrder, &weightMsgCancelSellOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelSellOrder = defaultWeightMsgCancelSellOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelSellOrder,
		dexsimulation.SimulateMsgCancelSellOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelBuyOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelBuyOrder, &weightMsgCancelBuyOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelBuyOrder = defaultWeightMsgCancelBuyOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelBuyOrder,
		dexsimulation.SimulateMsgCancelBuyOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
