package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"ollo/x/dex/keeper"
	"ollo/x/dex/types"
)

func SimulateMsgCancelBuyOrder(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCancelBuyOrder{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CancelBuyOrder simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CancelBuyOrder simulation not implemented"), nil, nil
	}
}
