package market

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ollo-station/ollo/x/market/keeper"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	var log = k.Logger(ctx)
	// err := k.UpdateMarket(ctx)
	// if err != nil {
	//     log.Error("Error updating market", "error", err)
	// }
	log.Info("updated market")
	return []abcitypes.ValidatorUpdate{}
}
