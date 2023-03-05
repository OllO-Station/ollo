package keeper

import (
    "fmt"
    "time"
    sdk "github.com/cosmos/cosmos-sdk/types"
    types "github.com/ollo-station/ollo/x/epoch/types"
    telemetry "github.com/cosmos/cosmos-sdk/telemetry"
)

func (k Keeper) BeginBlocker(c sdk.Context) {
    defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker) 
    k.IterateEpochs(c, func(num int64, epoch types.Epoch) (stop bool) {
        logger := k.Logger(c)
        if c.BlockTime().Before(epoch.Start) {
            return
        }
        shouldStartInit := !epoch.EpochStarted
        epochEnd := epoch.CurrentEpochStart.Add(epoch.Duration)
        shouldStart := (c.BlockTime().After(epochEnd)) || shouldStartInit
        if !shouldStart { return false }
        epoch.CurrentEpochStartHeight = uint64(c.BlockHeight())
        if shouldStartInit {
            epoch.EpochStarted = true
            epoch.CurrentEpochNumber = 1
            epoch.CurrentEpochStart = epoch.Start
            logger.Info(fmt.Sprintf("Epoch %d started at %s", epoch.CurrentEpochNumber, epoch.CurrentEpochStart))
        } else {
            c.EventManager().EmitEvent(
                sdk.NewEvent(
                    types.EventTypeEpochEnd,
                    sdk.NewAttribute(types.AttributeEpochNum, fmt.Sprintf("%d", epoch.CurrentEpochNumber)),
                ),
            )
            k.EpochEnd(c, epoch.Id, epoch.CurrentEpochNumber)
            epoch.CurrentEpochNumber += 1
            epoch.CurrentEpochStart = epoch.CurrentEpochStart.Add(epoch.Duration)
            logger.Info(fmt.Sprintf("Epoch %d started at %s", epoch.CurrentEpochNumber, epoch.CurrentEpochStart))
        }

        c.EventManager().EmitEvent(
            sdk.NewEvent(
                types.EventTypeEpochStart,
                sdk.NewAttribute(types.AttributeEpochNum, fmt.Sprintf("%d", epoch.CurrentEpochNumber)),
                sdk.NewAttribute(types.AttributeEpochStartBlockTime, fmt.Sprintf("%s", epoch.CurrentEpochStart.Unix())),
            ),
        )
        k.setEpoch(c, epoch)
        k.EpochStart(c, epoch.Id, epoch.CurrentEpochNumber)
        return false
    })
}

func (k Keeper) EndBlocker(c sdk.Context) {
}
