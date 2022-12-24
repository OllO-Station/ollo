package keeper_test

import (
	_ "github.com/stretchr/testify/suite"
)

// var _ types.GrantsHooks = &MockgrantsHooksReceiver{}

// // MockGrantsHooksReceiver event hooks for governance proposal object (noalias)
// type MockGrantsHooksReceiver struct {
// 	BeforeFixedPriceAuctionCreatedValid bool
// 	AfterFixedPriceAuctionCreatedValid  bool
// 	BeforeBatchAuctionCreatedValid      bool
// 	AfterBatchAuctionCreatedValid       bool
// 	BeforeAuctionCanceledValid          bool
// 	BeforeBidPlacedValid                bool
// 	BeforeBidModifiedValid              bool
// 	BeforeAllowedBiddersAddedValid      bool
// 	BeforeAllowedBidderUpdatedValid     bool
// 	BeforeSellingCoinsAllocatedValid    bool
// }

// func (h *MockGrantsHooksReceiver) BeforeFixedPriceAuctionCreated(
// 	ctx sdk.Context,
// 	auctioneer string,
// 	startPrice sdk.Dec,
// 	sellingCoin sdk.Coin,
// 	payingCoinDenom string,
// 	vestingSchedules []types.VestingSchedule,
// 	startTime time.Time,
// 	endTime time.Time,
// ) {
// 	h.BeforeFixedPriceAuctionCreatedValid = true
// }

// func (h *MockGrantsHooksReceiver) AfterFixedPriceAuctionCreated(
// 	ctx sdk.Context,
// 	auctionId uint64,
// 	auctioneer string,
// 	startPrice sdk.Dec,
// 	sellingCoin sdk.Coin,
// 	payingCoinDenom string,
// 	vestingSchedules []types.VestingSchedule,
// 	startTime time.Time,
// 	endTime time.Time,
// ) {
// 	h.AfterFixedPriceAuctionCreatedValid = true
// }

// func (h *MockGrantsHooksReceiver) BeforeBatchAuctionCreated(
// 	ctx sdk.Context,
// 	auctioneer string,
// 	startPrice sdk.Dec,
// 	minBidPrice sdk.Dec,
// 	sellingCoin sdk.Coin,
// 	payingCoinDenom string,
// 	vestingSchedules []types.VestingSchedule,
// 	maxExtendedRound uint32,
// 	extendedRoundRate sdk.Dec,
// 	startTime time.Time,
// 	endTime time.Time,
// ) {
// 	h.BeforeBatchAuctionCreatedValid = true
// }

// func (h *MockGrantsHooksReceiver) AfterBatchAuctionCreated(
// 	ctx sdk.Context,
// 	auctionId uint64,
// 	auctioneer string,
// 	startPrice sdk.Dec,
// 	minBidPrice sdk.Dec,
// 	sellingCoin sdk.Coin,
// 	payingCoinDenom string,
// 	vestingSchedules []types.VestingSchedule,
// 	maxExtendedRound uint32,
// 	extendedRoundRate sdk.Dec,
// 	startTime time.Time,
// 	endTime time.Time,
// ) {
// 	h.AfterBatchAuctionCreatedValid = true
// }

// func (h *MockGrantsHooksReceiver) BeforeAuctionCanceled(
// 	ctx sdk.Context,
// 	auctionId uint64,
// 	auctioneer string,
// ) {
// 	h.BeforeAuctionCanceledValid = true
// }

// func (h *MockGrantsHooksReceiver) BeforeBidPlaced(
// 	ctx sdk.Context,
// 	auctionId uint64,
// 	bidId uint64,
// 	bidder string,
// 	bidType types.BidType,
// 	price sdk.Dec,
// 	coin sdk.Coin,
// ) {
// 	h.BeforeBidPlacedValid = true
// }

// func (h *MockGrantsHooksReceiver) BeforeBidModified(
// 	ctx sdk.Context,
// 	auctionId uint64,
// 	bidId uint64,
// 	bidder string,
// 	bidType types.BidType,
// 	price sdk.Dec,
// 	coin sdk.Coin,
// ) {
// 	h.BeforeBidModifiedValid = true
// }

// func (h *MockGrantsHooksReceiver) BeforeAllowedBiddersAdded(
// 	ctx sdk.Context,
// 	allowedBidders []types.AllowedBidder,
// ) {
// 	h.BeforeAllowedBiddersAddedValid = true
// }

// func (h *MockGrantsHooksReceiver) BeforeAllowedBidderUpdated(
// 	ctx sdk.Context,
// 	auctionId uint64,
// 	bidder sdk.AccAddress,
// 	maxBidAmount sdk.Int,
// ) {
// 	h.BeforeAllowedBidderUpdatedValid = true
// }

// func (h *MockGrantsHooksReceiver) BeforeSellingCoinsAllocated(
// 	ctx sdk.Context,
// 	auctionId uint64,
// 	allocationMap map[string]sdk.Int,
// 	refundMap map[string]sdk.Int,
// ) {
// 	h.BeforeSellingCoinsAllocatedValid = true
// }

// func (s *KeeperTestSuite) TestHooks() {
// 	grantsHooksReceiver := MockGrantsHooksReceiver{}

// 	// Set hooks
// 	s.keeper.SetHooks(types.NewMultiGrantsHooks(&grantsHooksReceiver))

// 	s.Require().False(grantsHooksReceiver.BeforeFixedPriceAuctionCreatedValid)
// 	s.Require().False(grantsHooksReceiver.AfterFixedPriceAuctionCreatedValid)
// 	s.Require().False(grantsHooksReceiver.BeforeBatchAuctionCreatedValid)
// 	s.Require().False(grantsHooksReceiver.AfterBatchAuctionCreatedValid)
// 	s.Require().False(grantsHooksReceiver.BeforeAuctionCanceledValid)
// 	s.Require().False(grantsHooksReceiver.BeforeBidPlacedValid)
// 	s.Require().False(grantsHooksReceiver.BeforeBidModifiedValid)
// 	s.Require().False(grantsHooksReceiver.BeforeAllowedBiddersAddedValid)
// 	s.Require().False(grantsHooksReceiver.BeforeAllowedBidderUpdatedValid)
// 	s.Require().False(grantsHooksReceiver.BeforeSellingCoinsAllocatedValid)

// 	// Create a fixed price auction
// 	s.createFixedPriceAuction(
// 		s.addr(0),
// 		parseDec("2.0"),
// 		parseCoin("1_000_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().True(grantsHooksReceiver.BeforeFixedPriceAuctionCreatedValid)
// 	s.Require().True(grantsHooksReceiver.AfterFixedPriceAuctionCreatedValid)

// 	// Create a batch auction
// 	batchAuction := s.createBatchAuction(
// 		s.addr(1),
// 		parseDec("0.5"),
// 		parseDec("0.1"),
// 		parseCoin("1_000_000_000_000denom3"),
// 		"denom4",
// 		[]types.VestingSchedule{},
// 		1,
// 		sdk.MustNewDecFromStr("0.2"),
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().True(grantsHooksReceiver.BeforeBatchAuctionCreatedValid)
// 	s.Require().True(grantsHooksReceiver.AfterBatchAuctionCreatedValid)

// 	// Create auction that is stand by status
// 	standByAuction := s.createFixedPriceAuction(
// 		s.addr(2),
// 		parseDec("2.0"),
// 		parseCoin("1_000_000_000_000denom5"),
// 		"denom6",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 1, 0),
// 		time.Now().AddDate(0, 3, 0),
// 		true,
// 	)

// 	// Cancel the auction
// 	err := s.keeper.CancelAuction(s.ctx, &types.MsgCancelAuction{
// 		Auctioneer: standByAuction.Auctioneer,
// 		AuctionId:  standByAuction.Id,
// 	})
// 	s.Require().NoError(err)
// 	s.Require().True(grantsHooksReceiver.BeforeAuctionCanceledValid)

// 	// Get already started batch auction
// 	auction, found := s.keeper.GetAuction(s.ctx, batchAuction.Id)
// 	s.Require().True(found)

// 	// Add allowed bidder
// 	allowedBidders := []types.AllowedBidder{types.NewAllowedBidder(s.addr(3), parseInt("100_000_000_000"))}
// 	s.Require().NoError(s.keeper.AddAllowedBidders(s.ctx, auction.GetId(), allowedBidders))
// 	s.Require().True(grantsHooksReceiver.BeforeAllowedBiddersAddedValid)

// 	// Update the allowed bidder
// 	err = s.keeper.UpdateAllowedBidder(s.ctx, auction.GetId(), s.addr(3), parseInt("110_000_000_000"))
// 	s.Require().NoError(err)
// 	s.Require().True(grantsHooksReceiver.BeforeAllowedBidderUpdatedValid)

// 	// Place a bid
// 	bid := s.placeBidBatchWorth(auction.GetId(), s.addr(3), parseDec("0.55"), parseCoin("5_000_000denom4"), sdk.NewInt(10_000_000), true)
// 	s.Require().True(grantsHooksReceiver.BeforeBidPlacedValid)

// 	// Modify the bid
// 	s.fundAddr(bid.GetBidder(), sdk.NewCoins(parseCoin("1_000_000denom4")))
// 	err = s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: bid.AuctionId,
// 		BidId:     bid.Id,
// 		Bidder:    bid.Bidder,
// 		Price:     bid.Price,
// 		Coin:      parseCoin("6_000_000denom4"),
// 	})
// 	s.Require().NoError(err)
// 	s.Require().True(grantsHooksReceiver.BeforeBidModifiedValid)

// 	// Calculate fixed price allocation
// 	mInfo := s.keeper.CalculateFixedPriceAllocation(s.ctx, auction)

// 	// Allocate the selling coin
// 	err = s.keeper.AllocateSellingCoin(s.ctx, auction, mInfo)
// 	s.Require().NoError(err)
// 	s.Require().True(grantsHooksReceiver.BeforeSellingCoinsAllocatedValid)
// }
