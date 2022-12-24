package keeper_test

// import (
// 	"time"

// 	"ollo/x/grants"

// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"ollo/x/grants/keeper"
// 	"ollo/x/grants/types"

// 	_ "github.com/stretchr/testify/suite"
// )

// func (s *KeeperTestSuite) TestSellingPoolReserveAmountInvariant() {
// 	logger := s.keeper.Logger(s.ctx)
// 	logger.Info("TestSellingPoolReserveAmountInvariant")

// 	// Create a fixed price auction that has started status
// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		sdk.MustNewDecFromStr("0.5"),
// 		sdk.NewInt64Coin("denom1", 500_000_000_000),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 3, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	_, broken := keeper.SellingPoolReserveAmountInvariant(s.keeper)(s.ctx)
// 	s.Require().False(broken)

// 	// Although it is not possible for an exploiter to have the same token denom in reality,
// 	// it is safe to test the case anyway
// 	exploiterAddr := s.addr(1)
// 	sellingReserveAddr := auction.GetSellingReserveAddress()
// 	s.sendCoins(exploiterAddr, sellingReserveAddr, sdk.NewCoins(
// 		sdk.NewInt64Coin("denom1", 500_000_000),
// 		sdk.NewInt64Coin("denom2", 500_000_000),
// 		sdk.NewInt64Coin("denom3", 500_000_000),
// 		sdk.NewInt64Coin("denom4", 500_000_000),
// 	), true)

// 	_, broken = keeper.SellingPoolReserveAmountInvariant(s.keeper)(s.ctx)
// 	s.Require().False(broken)
// }

// func (s *KeeperTestSuite) TestPayingPoolReserveAmountInvariant() {
// 	logger := s.keeper.Logger(s.ctx)
// 	logger.Info("TestPayingPoolReserveAmountInvariant")

// 	k, ctx := s.keeper, s.ctx

// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		sdk.OneDec(),
// 		sdk.NewInt64Coin("denom3", 500_000_000_000),
// 		"denom4",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 3, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	s.placeBidFixedPrice(auction.GetId(), s.addr(1), sdk.OneDec(), parseCoin("20000000denom4"), true)
// 	s.placeBidFixedPrice(auction.GetId(), s.addr(2), sdk.OneDec(), parseCoin("20000000denom4"), true)
// 	s.placeBidFixedPrice(auction.GetId(), s.addr(2), sdk.OneDec(), parseCoin("15000000denom4"), true)
// 	s.placeBidFixedPrice(auction.GetId(), s.addr(3), sdk.OneDec(), parseCoin("35000000denom4"), true)
// 	s.placeBidFixedPrice(auction.GetId(), s.addr(4), sdk.OneDec(), parseCoin("15000000denom3"), true)
// 	s.placeBidFixedPrice(auction.GetId(), s.addr(5), sdk.OneDec(), parseCoin("20000000denom3"), true)

// 	_, broken := keeper.PayingPoolReserveAmountInvariant(k)(ctx)
// 	s.Require().False(broken)

// 	// Although it is not possible for an exploiter to have the same token denom in reality,
// 	// it is safe to test the case anyway
// 	exploiterAddr := s.addr(1)
// 	payingReserveAddr := auction.GetPayingReserveAddress()
// 	s.sendCoins(exploiterAddr, payingReserveAddr, sdk.NewCoins(
// 		sdk.NewInt64Coin("denom1", 500_000_000),
// 		sdk.NewInt64Coin("denom2", 500_000_000),
// 		sdk.NewInt64Coin("denom3", 500_000_000),
// 		sdk.NewInt64Coin("denom4", 500_000_000),
// 	), true)

// 	_, broken = keeper.PayingPoolReserveAmountInvariant(k)(ctx)
// 	s.Require().False(broken)
// }

// func (s *KeeperTestSuite) TestVestingPoolReserveAmountInvariant() {
// 	logger := s.keeper.Logger(s.ctx)
// 	logger.Info("TestVestingPoolReserveAmountInvariant")

// 	k, ctx := s.keeper, s.ctx

// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		sdk.OneDec(),
// 		sdk.NewInt64Coin("denom3", 500_000_000_000),
// 		"denom4",
// 		[]types.VestingSchedule{
// 			{
// 				ReleaseTime: time.Now().AddDate(1, 0, 0),
// 				Weight:      sdk.MustNewDecFromStr("0.25"),
// 			},
// 			{
// 				ReleaseTime: time.Now().AddDate(1, 3, 0),
// 				Weight:      sdk.MustNewDecFromStr("0.25"),
// 			},
// 			{
// 				ReleaseTime: time.Now().AddDate(1, 6, 0),
// 				Weight:      sdk.MustNewDecFromStr("0.25"),
// 			},
// 			{
// 				ReleaseTime: time.Now().AddDate(1, 9, 0),
// 				Weight:      sdk.MustNewDecFromStr("0.25"),
// 			},
// 		},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 3, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	s.placeBidFixedPrice(auction.GetId(), s.addr(1), sdk.OneDec(), parseCoin("20000000denom4"), true)
// 	s.placeBidFixedPrice(auction.GetId(), s.addr(2), sdk.OneDec(), parseCoin("20000000denom4"), true)
// 	s.placeBidFixedPrice(auction.GetId(), s.addr(2), sdk.OneDec(), parseCoin("15000000denom4"), true)
// 	s.placeBidFixedPrice(auction.GetId(), s.addr(3), sdk.OneDec(), parseCoin("35000000denom4"), true)

// 	// Make the auction ended
// 	ctx = ctx.WithBlockTime(auction.GetEndTimes()[0].AddDate(0, 0, 1))
// 	grants.BeginBlocker(ctx, k)

// 	// Make first and second vesting queues over
// 	ctx = ctx.WithBlockTime(auction.GetVestingSchedules()[0].GetReleaseTime().AddDate(0, 0, 1))
// 	grants.BeginBlocker(ctx, k)

// 	_, broken := keeper.VestingPoolReserveAmountInvariant(k)(ctx)
// 	s.Require().False(broken)

// 	// Although it is not possible for an exploiter to have the same token denom in reality,
// 	// it is safe to test the case anyway
// 	exploiterAddr := s.addr(1)
// 	vestingReserveAddr := auction.GetVestingReserveAddress()
// 	s.sendCoins(exploiterAddr, vestingReserveAddr, sdk.NewCoins(
// 		sdk.NewInt64Coin("denom1", 500_000_000),
// 		sdk.NewInt64Coin("denom2", 500_000_000),
// 		sdk.NewInt64Coin("denom3", 500_000_000),
// 		sdk.NewInt64Coin("denom4", 500_000_000),
// 	), true)

// 	_, broken = keeper.VestingPoolReserveAmountInvariant(k)(ctx)
// 	s.Require().False(broken)
// }
