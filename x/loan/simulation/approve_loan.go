package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"ollo/x/loan/keeper"
	"ollo/x/loan/types"
)

func SimulateMsgApproveLoan(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgApproveLoan{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ApproveLoan simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ApproveLoan simulation not implemented"), nil, nil
	}
}
