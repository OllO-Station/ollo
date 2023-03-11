package lend

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/ollo-station/ollo/testutil/sample"
	lendsimulation "github.com/ollo-station/ollo/x/lend/simulation"
	"github.com/ollo-station/ollo/x/lend/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = lendsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRequestLoan = "op_weight_msg_request_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestLoan int = 100

	opWeightMsgApproveLoan = "op_weight_msg_approve_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveLoan int = 100

	opWeightMsgRepayLoan = "op_weight_msg_repay_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRepayLoan int = 100

	opWeightMsgLiquidateLoan = "op_weight_msg_liquidate_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLiquidateLoan int = 100

	opWeightMsgCancelLoan = "op_weight_msg_cancel_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelLoan int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	lendGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&lendGenesis)
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

	var weightMsgRequestLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestLoan, &weightMsgRequestLoan, nil,
		func(_ *rand.Rand) {
			weightMsgRequestLoan = defaultWeightMsgRequestLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestLoan,
		lendsimulation.SimulateMsgRequestLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveLoan, &weightMsgApproveLoan, nil,
		func(_ *rand.Rand) {
			weightMsgApproveLoan = defaultWeightMsgApproveLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveLoan,
		lendsimulation.SimulateMsgApproveLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRepayLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRepayLoan, &weightMsgRepayLoan, nil,
		func(_ *rand.Rand) {
			weightMsgRepayLoan = defaultWeightMsgRepayLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRepayLoan,
		lendsimulation.SimulateMsgRepayLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLiquidateLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLiquidateLoan, &weightMsgLiquidateLoan, nil,
		func(_ *rand.Rand) {
			weightMsgLiquidateLoan = defaultWeightMsgLiquidateLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLiquidateLoan,
		lendsimulation.SimulateMsgLiquidateLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelLoan, &weightMsgCancelLoan, nil,
		func(_ *rand.Rand) {
			weightMsgCancelLoan = defaultWeightMsgCancelLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelLoan,
		lendsimulation.SimulateMsgCancelLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
