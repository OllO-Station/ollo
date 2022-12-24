package keeper_test

import (
	_ "github.com/stretchr/testify/suite"
)

// func (s *KeeperTestSuite) TestPlaceBid_Validation() {
// 	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: 1,
// 		Bidder:    s.addr(2).String(),
// 		BidType:   types.BidTypeFixedPrice,
// 		Price:     parseDec("0.5"),
// 		Coin:      parseCoin("200_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, sdkerrors.ErrNotFound)

// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		parseDec("1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	auction.SetStatus(types.AuctionStatusCancelled)
// 	s.keeper.SetAuction(s.ctx, auction)

// 	_, err = s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(2).String(),
// 		BidType:   types.BidTypeFixedPrice,
// 		Price:     parseDec("0.5"),
// 		Coin:      parseCoin("200_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrInvalidAuctionStatus)
// }

// func (s *KeeperTestSuite) TestFixedPrice_InvalidStartPrice() {
// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		parseDec("1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	// Fund bidder
// 	s.fundAddr(s.addr(2), parseCoins("200_000_000denom2"))

// 	// Set allowed bidder
// 	s.addAllowedBidder(auction.Id, s.addr(2), bidSellingAmount(parseDec("1"), parseCoin("200_000_000denom2")))

// 	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(2).String(),
// 		BidType:   types.BidTypeFixedPrice,
// 		Price:     parseDec("0.5"),
// 		Coin:      parseCoin("200_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrInvalidStartPrice)
// }

// func (s *KeeperTestSuite) TestFixedPrice_InsufficientRemainingAmount() {
// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		parseDec("1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	s.placeBidFixedPrice(auction.Id, s.addr(1), parseDec("1"), parseCoin("200_000_000denom2"), true)
// 	s.placeBidFixedPrice(auction.Id, s.addr(2), parseDec("1"), parseCoin("200_000_000denom2"), true)
// 	s.placeBidFixedPrice(auction.Id, s.addr(3), parseDec("1"), parseCoin("250_000_000denom2"), true)
// 	s.placeBidFixedPrice(auction.Id, s.addr(4), parseDec("1"), parseCoin("250_000_000denom2"), true)

// 	// The remaining coin amount must be insufficient
// 	s.fundAddr(s.addr(5), parseCoins("300_000_000denom2"))
// 	s.addAllowedBidder(auction.Id, s.addr(5), bidSellingAmount(parseDec("1"), parseCoin("300_000_000denom2")))

// 	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(5).String(),
// 		BidType:   types.BidTypeFixedPrice,
// 		Price:     parseDec("1.0"),
// 		Coin:      parseCoin("300_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrInsufficientRemainingAmount)
// }

// func (s *KeeperTestSuite) TestFixedPrice_OverMaxBidAmountLimit() {
// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		parseDec("1"),
// 		parseCoin("1000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	s.placeBidFixedPrice(auction.Id, s.addr(1), parseDec("1"), parseCoin("100_000_000denom2"), true)
// 	s.placeBidFixedPrice(auction.Id, s.addr(1), parseDec("1"), parseCoin("100_000_000denom2"), true)

// 	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidType:   types.BidTypeFixedPrice,
// 		Price:     parseDec("1"),
// 		Coin:      parseCoin("100_000_001denom2"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrOverMaxBidAmountLimit)
// }

// func (s *KeeperTestSuite) TestFixedPrice_IncorrectCoinDenom() {
// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		parseDec("1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	s.placeBidFixedPrice(auction.Id, s.addr(1), parseDec("1"), parseCoin("100_000_000denom2"), true)
// 	s.placeBidFixedPrice(auction.Id, s.addr(1), parseDec("1"), parseCoin("100_000_000denom1"), true)

// 	// The remaining coin amount must be insufficient
// 	s.fundAddr(s.addr(1), parseCoins("100_000_000denom2"))
// 	s.addAllowedBidder(auction.Id, s.addr(1), bidSellingAmount(parseDec("1"), parseCoin("100_000_000denom2")))

// 	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidType:   types.BidTypeFixedPrice,
// 		Price:     parseDec("1"),
// 		Coin:      parseCoin("10_000_000denom3"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrIncorrectCoinDenom)
// }

// func (s *KeeperTestSuite) TestFixedPrice_IncorrectAuctionType() {
// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		parseDec("1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	s.fundAddr(s.addr(2), parseCoins("200_000_000denom2"))
// 	s.addAllowedBidder(auction.Id, s.addr(2), bidSellingAmount(parseDec("1"), parseCoin("200_000_000denom2")))

// 	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(2).String(),
// 		BidType:   types.BidTypeBatchWorth,
// 		Price:     parseDec("1.0"),
// 		Coin:      parseCoin("200_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrIncorrectAuctionType)
// }

// func (s *KeeperTestSuite) TestBatchAuction_IncorrectCoinDenom() {
// 	auction := s.createBatchAuction(
// 		s.addr(1),
// 		parseDec("0.5"),
// 		parseDec("0.1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		1,
// 		sdk.MustNewDecFromStr("0.2"),
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	s.fundAddr(s.addr(1), parseCoins("200_000_000denom1, 200_000_000denom2"))
// 	s.addAllowedBidder(auction.Id, s.addr(1), parseCoin("200_000_000denom1").Amount)
// 	s.addAllowedBidder(auction.Id, s.addr(1), parseCoin("200_000_000denom2").Amount)

// 	// Place a BidTypeBatchWorth bid with an incorrect denom (SellingCoinDenom)
// 	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidType:   types.BidTypeBatchWorth,
// 		Price:     parseDec("1"),
// 		Coin:      parseCoin("100_000_000denom1"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrIncorrectCoinDenom)

// 	// Place a BidTypeBatchMany bid with an incorrect denom (PayingCoinDenom)
// 	_, err = s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidType:   types.BidTypeBatchMany,
// 		Price:     parseDec("1"),
// 		Coin:      parseCoin("100_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrIncorrectCoinDenom)
// }

// func (s *KeeperTestSuite) TestBatchWorth_OverMaxBidAmountLimit() {
// 	auction := s.createBatchAuction(
// 		s.addr(1),
// 		parseDec("0.5"),
// 		parseDec("0.1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		1,
// 		sdk.MustNewDecFromStr("0.2"),
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	s.placeBidBatchWorth(auction.Id, s.addr(1), parseDec("0.5"), parseCoin("500_000_000denom2"), sdk.NewInt(1000_000_000), true)

// 	s.fundAddr(s.addr(2), parseCoins("1000_000_000denom2"))
// 	s.addAllowedBidder(auction.Id, s.addr(2), parseCoin("800_000_000denom1").Amount)

// 	// Place a BidTypeBatchWorth bid with more than maxBidAmount
// 	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(2).String(),
// 		BidType:   types.BidTypeBatchWorth,
// 		Price:     parseDec("0.5"),
// 		Coin:      parseCoin("500_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrOverMaxBidAmountLimit)
// }

// func (s *KeeperTestSuite) TestBatchMany_OverMaxBidAmountLimit() {
// 	auction := s.createBatchAuction(
// 		s.addr(1),
// 		parseDec("0.5"),
// 		parseDec("0.1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		1,
// 		sdk.MustNewDecFromStr("0.2"),
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	s.placeBidBatchMany(auction.Id, s.addr(1), parseDec("0.5"), parseCoin("500_000_000denom1"), sdk.NewInt(800_000_000), true)

// 	s.fundAddr(s.addr(2), parseCoins("1000_000_000denom2"))
// 	s.addAllowedBidder(auction.Id, s.addr(2), parseCoin("400_000_000denom1").Amount)

// 	// Place a BidTypeBatchMany bid with more than maxBidAmount
// 	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
// 		AuctionId: auction.Id,
// 		Bidder:    s.addr(2).String(),
// 		BidType:   types.BidTypeBatchMany,
// 		Price:     parseDec("0.5"),
// 		Coin:      parseCoin("500_000_000denom1"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrOverMaxBidAmountLimit)
// }

// func (s *KeeperTestSuite) TestModifyBid_Validation() {
// 	err := s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: 1,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     5,
// 		Price:     parseDec("0.8"),
// 		Coin:      parseCoin("100_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, sdkerrors.ErrNotFound)

// 	auction := s.createFixedPriceAuction(
// 		s.addr(0),
// 		parseDec("1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, auction.GetStatus())

// 	err = s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: 1,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     5,
// 		Price:     parseDec("0.8"),
// 		Coin:      parseCoin("100_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrIncorrectAuctionType)

// 	auction.SetStatus(types.AuctionStatusCancelled)
// 	s.keeper.SetAuction(s.ctx, auction)

// 	err = s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: 1,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     5,
// 		Price:     parseDec("0.8"),
// 		Coin:      parseCoin("100_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrInvalidAuctionStatus)

// }

// func (s *KeeperTestSuite) TestModifyBid_BidTypeWorth() {
// 	a := s.createBatchAuction(
// 		s.addr(0),
// 		parseDec("0.1"),
// 		parseDec("0.1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		1,
// 		sdk.MustNewDecFromStr("0.2"),
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, a.GetStatus())

// 	// Place a bid
// 	b := s.placeBidBatchWorth(a.Id, s.addr(1), parseDec("0.6"), parseCoin("100_000_000denom2"), sdk.NewInt(1_000_000_000), true)

// 	// Modify the bid with not existing bid
// 	err := s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: a.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     5,
// 		Price:     parseDec("0.8"),
// 		Coin:      parseCoin("100_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, sdkerrors.ErrNotFound)

// 	// Modify the bid with an incorrect owner
// 	err = s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: a.Id,
// 		Bidder:    s.addr(0).String(),
// 		BidId:     b.Id,
// 		Price:     parseDec("0.8"),
// 		Coin:      parseCoin("100_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, sdkerrors.ErrUnauthorized)

// 	// Modify the bid with an incorrect denom
// 	err = s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: a.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     b.Id,
// 		Price:     parseDec("0.8"),
// 		Coin:      parseCoin("100_000_000denom1"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrIncorrectCoinDenom)

// 	// Modify the bid with lower bid price
// 	err = s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: a.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     b.Id,
// 		Price:     parseDec("0.3"),
// 		Coin:      parseCoin("100_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)

// 	// Modify the bid with lower coin amount
// 	err = s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: a.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     b.Id,
// 		Price:     parseDec("0.8"),
// 		Coin:      parseCoin("100denom2"),
// 	})
// 	s.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)

// 	// No modification
// 	err = s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: a.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     b.Id,
// 		Price:     parseDec("0.6"),
// 		Coin:      parseCoin("100_000_000denom2"),
// 	})
// 	s.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)
// }

// func (s *KeeperTestSuite) TestModifyBid_BidTypeMany() {
// 	a := s.createBatchAuction(
// 		s.addr(0),
// 		parseDec("0.1"),
// 		parseDec("0.1"),
// 		parseCoin("1_000_000_000denom1"),
// 		"denom2",
// 		[]types.VestingSchedule{},
// 		1,
// 		sdk.MustNewDecFromStr("0.2"),
// 		time.Now().AddDate(0, 0, -1),
// 		time.Now().AddDate(0, 0, -1).AddDate(0, 2, 0),
// 		true,
// 	)
// 	s.Require().Equal(types.AuctionStatusStarted, a.GetStatus())

// 	// Place a bid
// 	b := s.placeBidBatchMany(a.Id, s.addr(1), parseDec("0.5"), parseCoin("100_000_000denom1"), sdk.NewInt(1_000_000_000), true)

// 	// Insufficient minimum price
// 	err := s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: a.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     b.Id,
// 		Price:     parseDec("0.01"),
// 		Coin:      parseCoin("100_000_000denom1"),
// 	})
// 	s.Require().ErrorIs(err, types.ErrInsufficientMinBidPrice)

// 	// Fund the bidder enough paying coin
// 	s.fundAddr(s.addr(1), parseCoins("1_000_000_000_000denom2"))

// 	// Modify the bid with not existing bid
// 	err = s.keeper.ModifyBid(s.ctx, &types.MsgModifyBid{
// 		AuctionId: a.Id,
// 		Bidder:    s.addr(1).String(),
// 		BidId:     b.Id,
// 		Price:     parseDec("0.8"),
// 		Coin:      parseCoin("100_000_000denom1"),
// 	})
// 	s.Require().NoError(err)
// }
