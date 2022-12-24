package grants

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/grants/keeper"
	"ollo/x/grants/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// Get all auctions from the store and execute operations depending on auction status.
	for _, auction := range k.GetAuctions(ctx) {
		switch auction.GetStatus() {
		case types.AuctionStatusStandBy:
			k.ExecuteStandByStatus(ctx, auction)
		case types.AuctionStatusStarted:
			k.ExecuteStartedStatus(ctx, auction)
		case types.AuctionStatusVesting:
			k.ExecuteVestingStatus(ctx, auction)
		default:
			continue
		}
	}
}
