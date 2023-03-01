package simulation

import (
	"math/rand"

	"github.com/ollo-station/ollo/x/ons/keeper"
	"github.com/ollo-station/ollo/x/ons/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSetName(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSetName{
			CreatorAddr: simAccount.Address.String(),
		}

		// TODO: Handling the SetName simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SetName simulation not implemented"), nil, nil
	}
}
