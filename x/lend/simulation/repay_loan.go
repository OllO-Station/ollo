package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/ollo-station/ollo/x/lend/keeper"
	"github.com/ollo-station/ollo/x/lend/types"
)

func SimulateMsgRepayLoan(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRepayLoan{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RepayLoan simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RepayLoan simulation not implemented"), nil, nil
	}
}
