package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/x/grants/types"
)

// Implements FundraisingHooks interface
var _ types.FundraisingHooks = Keeper{}

// BeforeFixedPriceAuctionCreated - call hook if registered
func (k Keeper) BeforeFixedPriceAuctionCreated(
	ctx sdk.Context,
	auctioneer string,
	startPrice sdk.Dec,
	sellingCoin sdk.Coin,
	payingCoinDenom string,
	vestingSchedules []types.VestingSchedule,
	startTime time.Time,
	endTime time.Time,
) {
	if k.hooks != nil {
		k.hooks.BeforeFixedPriceAuctionCreated(
			ctx,
			auctioneer,
			startPrice,
			sellingCoin,
			payingCoinDenom,
			vestingSchedules,
			startTime,
			endTime,
		)
	}
}

// AfterFixedPriceAuctionCreated - call hook if registered
func (k Keeper) AfterFixedPriceAuctionCreated(
	ctx sdk.Context,
	auctionId uint64,
	auctioneer string,
	startPrice sdk.Dec,
	sellingCoin sdk.Coin,
	payingCoinDenom string,
	vestingSchedules []types.VestingSchedule,
	startTime time.Time,
	endTime time.Time,
) {
	if k.hooks != nil {
		k.hooks.AfterFixedPriceAuctionCreated(
			ctx,
			auctionId,
			auctioneer,
			startPrice,
			sellingCoin,
			payingCoinDenom,
			vestingSchedules,
			startTime,
			endTime,
		)
	}
}

// BeforeBatchAuctionCreated - call hook if registered
func (k Keeper) BeforeBatchAuctionCreated(
	ctx sdk.Context,
	auctioneer string,
	startPrice sdk.Dec,
	minBidPrice sdk.Dec,
	sellingCoin sdk.Coin,
	payingCoinDenom string,
	vestingSchedules []types.VestingSchedule,
	maxExtendedRound uint32,
	extendedRoundRate sdk.Dec,
	startTime time.Time,
	endTime time.Time,
) {
	if k.hooks != nil {
		k.hooks.BeforeBatchAuctionCreated(
			ctx,
			auctioneer,
			startPrice,
			minBidPrice,
			sellingCoin,
			payingCoinDenom,
			vestingSchedules,
			maxExtendedRound,
			extendedRoundRate,
			startTime,
			endTime,
		)
	}
}

// AfterBatchAuctionCreated - call hook if registered
func (k Keeper) AfterBatchAuctionCreated(
	ctx sdk.Context,
	auctionId uint64,
	auctioneer string,
	startPrice sdk.Dec,
	minBidPrice sdk.Dec,
	sellingCoin sdk.Coin,
	payingCoinDenom string,
	vestingSchedules []types.VestingSchedule,
	maxExtendedRound uint32,
	extendedRoundRate sdk.Dec,
	startTime time.Time,
	endTime time.Time,
) {
	if k.hooks != nil {
		k.hooks.AfterBatchAuctionCreated(
			ctx,
			auctionId,
			auctioneer,
			startPrice,
			minBidPrice,
			sellingCoin,
			payingCoinDenom,
			vestingSchedules,
			maxExtendedRound,
			extendedRoundRate,
			startTime,
			endTime,
		)
	}
}

// BeforeAuctionCanceled - call hook if registered
func (k Keeper) BeforeAuctionCanceled(
	ctx sdk.Context,
	auctionId uint64,
	auctioneer string,
) {
	if k.hooks != nil {
		k.hooks.BeforeAuctionCanceled(ctx, auctionId, auctioneer)
	}
}

// BeforeBidPlaced - call hook if registered
func (k Keeper) BeforeBidPlaced(
	ctx sdk.Context,
	auctionId uint64,
	bidId uint64,
	bidder string,
	bidType types.BidType,
	price sdk.Dec,
	coin sdk.Coin,
) {
	if k.hooks != nil {
		k.hooks.BeforeBidPlaced(ctx, auctionId, bidId, bidder, bidType, price, coin)
	}
}

// BeforeBidModified - call hook if registered
func (k Keeper) BeforeBidModified(
	ctx sdk.Context,
	auctionId uint64,
	bidId uint64,
	bidder string,
	bidType types.BidType,
	price sdk.Dec,
	coin sdk.Coin,
) {
	if k.hooks != nil {
		k.hooks.BeforeBidModified(ctx, auctionId, bidId, bidder, bidType, price, coin)
	}
}

// BeforeAllowedBiddersAdded - call hook if registered
func (k Keeper) BeforeAllowedBiddersAdded(
	ctx sdk.Context,
	allowedBidders []types.AllowedBidder,
) {
	if k.hooks != nil {
		k.hooks.BeforeAllowedBiddersAdded(ctx, allowedBidders)
	}
}

// BeforeAllowedBidderUpdated - call hook if registered
func (k Keeper) BeforeAllowedBidderUpdated(
	ctx sdk.Context,
	auctionId uint64,
	bidder sdk.AccAddress,
	maxBidAmount sdk.Int,
) {
	if k.hooks != nil {
		k.hooks.BeforeAllowedBidderUpdated(ctx, auctionId, bidder, maxBidAmount)
	}
}

// BeforeSellingCoinsAllocated - call hook if registered
func (k Keeper) BeforeSellingCoinsAllocated(
	ctx sdk.Context,
	auctionId uint64,
	allocationMap map[string]sdk.Int,
	refundMap map[string]sdk.Int,
) {
	if k.hooks != nil {
		k.hooks.BeforeSellingCoinsAllocated(ctx, auctionId, allocationMap, refundMap)
	}
}
