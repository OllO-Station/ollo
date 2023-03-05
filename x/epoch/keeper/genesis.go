package keeper
import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/ollo-station/ollo/x/epoch/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
    for _, epoch := range data.Epochs {
        err := k.AddEpoch(ctx, epoch)
        if err != nil {
            panic(err)
        }
    }
}

func (k Keeper) ExportGenesis(ctx sdk.Context) (data types.GenesisState) {
    return types.GenesisState{
        Epochs: k.AllEpochs(ctx),
    }
}
