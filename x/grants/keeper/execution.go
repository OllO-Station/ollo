package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/grants/types"
)

// ExecuteStandByStatus simply updates the auction status to AuctionStatusStarted
// if the auction is ready to get started.
func (k Keeper) ExecuteStandByStatus(ctx sdk.Context, auction types.AuctionI) {
	if auction.ShouldAuctionStarted(ctx.BlockTime()) { // BlockTime >= StartTime
		if err := auction.SetStatus(types.AuctionStatusStarted); err != nil {
			panic(err)
		}
		k.SetAuction(ctx, auction)
	}
}

// ExecuteStartedStatus executes operations depending on the auction type.
func (k Keeper) ExecuteStartedStatus(ctx sdk.Context, auction types.AuctionI) {
	if auction.ShouldAuctionClosed(ctx.BlockTime()) { // BlockTime >= EndTime
		switch auction.GetType() {
		case types.AuctionTypeFixedPrice:
			k.CloseFixedPriceAuction(ctx, auction)

		case types.AuctionTypeBatch:
			k.CloseBatchAuction(ctx, auction)
		}
	}
}

// ExecuteVestingStatus first gets all vesting queues in the store and
// look up the release time of each vesting queue to see if the module needs to
// distribute the paying coin to the auctioneer.
func (k Keeper) ExecuteVestingStatus(ctx sdk.Context, auction types.AuctionI) {
	if err := k.ReleaseVestingPayingCoin(ctx, auction); err != nil {
		panic(err)
	}
}
