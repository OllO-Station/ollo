package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ollotypes "github.com/ollo-station/ollo/x/ollo/types"
)

func (app *App) ScheduleFork(ctx sdk.Context) {
	if !ollotypes.IsMainnet(ctx.ChainID()) {
		return
	}
	upgradePlan := upgradetypes.Plan{
		Height: ctx.BlockHeight(),
	}

	switch ctx.BlockHeight() {
	case -1:
		upgradePlan.Name = "upgrade-1"
		upgradePlan.Info = "upgrade-1"
	default:
		return
	}
	if err := app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradePlan); err != nil {
		panic(
			fmt.Errorf(
				"failed to schedule upgrade %s at height %d: %w",
				upgradePlan.Name,
				ctx.BlockHeight(),
				err,
			),
		)
	}
}
