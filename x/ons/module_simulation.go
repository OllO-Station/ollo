package ons

import (
	"math/rand"

	"github.com/ollo-station/ollo/testutil/sample"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	onssimulation "github.com/ollo-station/ollo/x/ons/simulation"
	"github.com/ollo-station/ollo/x/ons/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = onssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgBuyName = "op_weight_msg_buy_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyName int = 100

	opWeightMsgSellName = "op_weight_msg_sell_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSellName int = 100

	opWeightMsgSetName = "op_weight_msg_set_name"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetName int = 100

	opWeightMsgAddThread = "op_weight_msg_add_thread"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddThread int = 100

	opWeightMsgDelThread = "op_weight_msg_del_thread"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDelThread int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	onsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&onsGenesis)
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

	var weightMsgBuyName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyName, &weightMsgBuyName, nil,
		func(_ *rand.Rand) {
			weightMsgBuyName = defaultWeightMsgBuyName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyName,
		onssimulation.SimulateMsgBuyName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSellName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSellName, &weightMsgSellName, nil,
		func(_ *rand.Rand) {
			weightMsgSellName = defaultWeightMsgSellName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSellName,
		onssimulation.SimulateMsgSellName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetName, &weightMsgSetName, nil,
		func(_ *rand.Rand) {
			weightMsgSetName = defaultWeightMsgSetName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetName,
		onssimulation.SimulateMsgSetName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddThread int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddThread, &weightMsgAddThread, nil,
		func(_ *rand.Rand) {
			weightMsgAddThread = defaultWeightMsgAddThread
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddThread,
		onssimulation.SimulateMsgAddThread(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDelThread int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDelThread, &weightMsgDelThread, nil,
		func(_ *rand.Rand) {
			weightMsgDelThread = defaultWeightMsgDelThread
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDelThread,
		onssimulation.SimulateMsgDelThread(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
