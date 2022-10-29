package simulation

// "math/rand"

// "ollo/x/loan/keeper"
// "ollo/x/loan/types"

// "github.com/cosmos/cosmos-sdk/baseapp"
// sdk "github.com/cosmos/cosmos-sdk/types"
// simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

// func SimulateMsgRequestLoan(
// 	ak types.AccountKeeper,
// 	bk types.BankKeeper,
// 	k keeper.Keeper,
// ) simtypes.Operation {
// 	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		simAccount, _ := simtypes.RandomAcc(r, accs)
// 		msg := &types.MsgRequestLoan{
// 			Creator: simAccount.Address.String(),
// 		}

// 		// TODO: Handling the RequestLoan simulation

// 		//return simtypes.NoOpMsg(types.ModuleName, msg.Type), "RequestLoan simulation not implemented"), nil, nil
// 	}
// }
