package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	simapp "ollo/app"
	"ollo/x/farming"
	"ollo/x/farming/types"

	_ "github.com/stretchr/testify/suite"
)

func (suite *KeeperTestSuite) TestAllocationInfos() {
	fixedAmountPlans := []types.PlanI{
		types.NewFixedAmountPlan(
			types.NewBasePlan(
				1,
				"",
				types.PlanTypePublic,
				suite.addrs[0].String(),
				suite.addrs[0].String(),
				sdk.NewDecCoins(sdk.NewDecCoinFromDec(denom1, sdk.NewDec(1))),
				types.ParseTime("2021-07-27T00:00:00Z"),
				types.ParseTime("2021-07-28T00:00:00Z"),
			),
			sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000))),
		types.NewFixedAmountPlan(
			types.NewBasePlan(
				2,
				"",
				types.PlanTypePublic,
				suite.addrs[0].String(),
				suite.addrs[0].String(),
				sdk.NewDecCoins(sdk.NewDecCoinFromDec(denom1, sdk.NewDec(1))),
				types.ParseTime("2021-07-27T12:00:00Z"),
				types.ParseTime("2021-07-28T12:00:00Z"),
			),
			sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000))),
	}

	ratioPlans := []types.PlanI{
		types.NewRatioPlan(
			types.NewBasePlan(
				3,
				"",
				types.PlanTypePublic,
				suite.addrs[0].String(),
				suite.addrs[0].String(),
				sdk.NewDecCoins(sdk.NewDecCoinFromDec(denom1, sdk.NewDec(1))),
				types.ParseTime("2021-07-27T00:00:00Z"),
				types.ParseTime("2021-07-28T00:00:00Z"),
			),
			sdk.MustNewDecFromStr("0.5")),
		types.NewRatioPlan(
			types.NewBasePlan(
				4,
				"",
				types.PlanTypePublic,
				suite.addrs[0].String(),
				suite.addrs[0].String(),
				sdk.NewDecCoins(sdk.NewDecCoinFromDec(denom1, sdk.NewDec(1))),
				types.ParseTime("2021-07-27T12:00:00Z"),
				types.ParseTime("2021-07-28T12:00:00Z"),
			),
			sdk.MustNewDecFromStr("0.6")),
	}

	hugeRatioPlan := types.NewRatioPlan(
		types.NewBasePlan(
			5,
			"",
			types.PlanTypePrivate,
			suite.addrs[0].String(),
			suite.addrs[0].String(),
			sdk.NewDecCoins(sdk.NewDecCoinFromDec(denom1, sdk.NewDec(1))),
			types.ParseTime("2021-07-27T12:00:00Z"),
			types.ParseTime("2021-07-28T12:00:00Z"),
		),
		sdk.MustNewDecFromStr("0.999999"))

	for _, tc := range []struct {
		name      string
		plans     []types.PlanI
		t         time.Time
		distrAmts map[uint64]sdk.Coins // planID => sdk.Coins
	}{
		{
			"insufficient farming pool balances",
			[]types.PlanI{
				types.NewFixedAmountPlan(
					types.NewBasePlan(
						1,
						"",
						types.PlanTypePrivate,
						suite.addrs[0].String(),
						suite.addrs[0].String(),
						sdk.NewDecCoins(sdk.NewDecCoinFromDec(denom1, sdk.NewDec(1))),
						types.ParseTime("2021-07-27T00:00:00Z"),
						types.ParseTime("2021-07-30T00:00:00Z"),
					),
					sdk.NewCoins(sdk.NewInt64Coin(denom3, 10_000_000_000))),
			},
			types.ParseTime("2021-07-28T00:00:00Z"),
			nil,
		},
		{
			"start time & end time edgecase #1",
			fixedAmountPlans,
			types.ParseTime("2021-07-26T23:59:59Z"),
			nil,
		},
		{
			"start time & end time edgecase #2",
			fixedAmountPlans,
			types.ParseTime("2021-07-27T00:00:00Z"),
			map[uint64]sdk.Coins{1: sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000))},
		},
		{
			"start time & end time edgecase #3",
			fixedAmountPlans,
			types.ParseTime("2021-07-27T11:59:59Z"),
			map[uint64]sdk.Coins{1: sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000))},
		},
		{
			"start time & end time edgecase #4",
			fixedAmountPlans,
			types.ParseTime("2021-07-27T12:00:00Z"),
			map[uint64]sdk.Coins{
				1: sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000)),
				2: sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000))},
		},
		{
			"start time & end time edgecase #5",
			fixedAmountPlans,
			types.ParseTime("2021-07-27T23:59:59Z"),
			map[uint64]sdk.Coins{
				1: sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000)),
				2: sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000))},
		},
		{
			"start time & end time edgecase #6",
			fixedAmountPlans,
			types.ParseTime("2021-07-28T00:00:00Z"),
			map[uint64]sdk.Coins{2: sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000))},
		},
		{
			"start time & end time edgecase #7",
			fixedAmountPlans,
			types.ParseTime("2021-07-28T11:59:59Z"),
			map[uint64]sdk.Coins{2: sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000))},
		},
		{
			"start time & end time edgecase #8",
			fixedAmountPlans,
			types.ParseTime("2021-07-28T12:00:00Z"),
			nil,
		},
		{
			"test case for ratio plans #1",
			ratioPlans,
			types.ParseTime("2021-07-27T00:00:00Z"),
			map[uint64]sdk.Coins{
				3: sdk.NewCoins(sdk.NewInt64Coin(denom1, 500000000), sdk.NewInt64Coin(denom2, 500000000),
					sdk.NewInt64Coin(denom3, 500000000), sdk.NewInt64Coin(sdk.DefaultBondDenom, 500000000000))},
		},
		{
			"test case for ratio plans #2",
			ratioPlans,
			types.ParseTime("2021-07-27T12:00:00Z"),
			nil,
		},
		{
			"test case for ratio plans #3",
			ratioPlans,
			types.ParseTime("2021-07-28T11:00:00Z"),
			map[uint64]sdk.Coins{
				4: sdk.NewCoins(sdk.NewInt64Coin(denom1, 600000000), sdk.NewInt64Coin(denom2, 600000000),
					sdk.NewInt64Coin(denom3, 600000000), sdk.NewInt64Coin(sdk.DefaultBondDenom, 600000000000))},
		},
		{
			"test case for fixed plans with a ratio plan over balance #1",
			append(fixedAmountPlans, hugeRatioPlan),
			types.ParseTime("2021-07-27T12:00:00Z"),
			nil,
		},
		{
			"test case for fixed plans with a ratio plan over balance #2",
			append([]types.PlanI{hugeRatioPlan}, fixedAmountPlans...),
			types.ParseTime("2021-07-27T12:00:00Z"),
			nil,
		},
	} {
		suite.Run(tc.name, func() {
			for _, plan := range suite.keeper.GetPlans(suite.ctx) {
				suite.keeper.DeletePlan(suite.ctx, plan)
			}
			for _, plan := range tc.plans {
				suite.keeper.SetPlan(suite.ctx, plan)
			}

			suite.ctx = suite.ctx.WithBlockTime(tc.t)
			distrInfos := suite.keeper.AllocationInfos(suite.ctx)
			if suite.Len(distrInfos, len(tc.distrAmts)) {
				for _, distrInfo := range distrInfos {
					distrAmt, ok := tc.distrAmts[distrInfo.Plan.GetId()]
					if suite.True(ok) {
						suite.True(coinsEq(distrAmt, distrInfo.Amount))
					}
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestAllocateRewards() {
	for _, plan := range suite.sampleFixedAmtPlans {
		_ = plan.SetStartTime(types.ParseTime("0001-01-01T00:00:00Z"))
		_ = plan.SetEndTime(types.ParseTime("9999-12-31T00:00:00Z"))
		suite.keeper.SetPlan(suite.ctx, plan)
	}

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))
	suite.Stake(suite.addrs[1], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000), sdk.NewInt64Coin(denom2, 1000000)))
	suite.keeper.ProcessQueuedCoins(suite.ctx)

	prevDistrCoins := map[uint64]sdk.Coins{}

	t := types.ParseTime("2021-09-01T00:00:00Z")
	for i := 0; i < 365; i++ {
		suite.ctx = suite.ctx.WithBlockTime(t)

		err := suite.keeper.AllocateRewards(suite.ctx)
		suite.Require().NoError(err)

		for _, plan := range suite.sampleFixedAmtPlans {
			plan, _ := suite.keeper.GetPlan(suite.ctx, plan.GetId())
			fixedAmtPlan := plan.(*types.FixedAmountPlan)

			dist := plan.GetDistributedCoins()
			suite.Require().True(coinsEq(prevDistrCoins[plan.GetId()].Add(fixedAmtPlan.EpochAmount...), dist))
			prevDistrCoins[plan.GetId()] = dist

			t2 := plan.GetLastDistributionTime()
			suite.Require().NotNil(t2)
			suite.Require().Equal(t, *t2)
		}

		t = t.AddDate(0, 0, 1)
	}
}

func (suite *KeeperTestSuite) TestAllocateRewards_FixedAmountPlanAllBalances() {
	farmingPoolAcc := simapp.AddTestAddrs(suite.app, suite.ctx, 1, sdk.ZeroInt())[0]
	err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, farmingPoolAcc, sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000000)))
	suite.Require().NoError(err)

	// The sum of epoch ratios is exactly 1.
	suite.CreateFixedAmountPlan(farmingPoolAcc, map[string]string{denom1: "1"}, map[string]int64{denom3: 600000})
	suite.CreateFixedAmountPlan(farmingPoolAcc, map[string]string{denom2: "1"}, map[string]int64{denom3: 400000})

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000), sdk.NewInt64Coin(denom2, 1000000)))

	suite.AdvanceEpoch()
	suite.AdvanceEpoch()

	rewards := suite.keeper.AllRewards(suite.ctx, suite.addrs[0])
	suite.Require().True(coinsEq(sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000000)), rewards))
}

func (suite *KeeperTestSuite) TestAllocateRewards_RatioPlanAllBalances() {
	farmingPoolAcc := simapp.AddTestAddrs(suite.app, suite.ctx, 1, sdk.ZeroInt())[0]
	err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, farmingPoolAcc, sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000000)))
	suite.Require().NoError(err)

	// The sum of epoch ratios is exactly 1.
	suite.CreateRatioPlan(farmingPoolAcc, map[string]string{denom1: "1"}, "0.5")
	suite.CreateRatioPlan(farmingPoolAcc, map[string]string{denom2: "1"}, "0.5")

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000), sdk.NewInt64Coin(denom2, 1000000)))

	suite.AdvanceEpoch()
	suite.AdvanceEpoch()

	rewards := suite.keeper.AllRewards(suite.ctx, suite.addrs[0])
	suite.Require().True(coinsEq(sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000000)), rewards))
}

func (suite *KeeperTestSuite) TestAllocateRewards_FixedAmountPlanOverBalances() {
	farmingPoolAcc := simapp.AddTestAddrs(suite.app, suite.ctx, 1, sdk.ZeroInt())[0]
	err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, farmingPoolAcc, sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000000)))
	suite.Require().NoError(err)

	// The sum of epoch amounts is over the balances the farming pool has,
	// so the reward allocation should never happen.
	suite.CreateFixedAmountPlan(farmingPoolAcc, map[string]string{denom1: "1"}, map[string]int64{denom3: 700000})
	suite.CreateFixedAmountPlan(farmingPoolAcc, map[string]string{denom2: "1"}, map[string]int64{denom3: 400000})

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000), sdk.NewInt64Coin(denom2, 1000000)))

	suite.AdvanceEpoch()
	suite.AdvanceEpoch()

	rewards := suite.keeper.AllRewards(suite.ctx, suite.addrs[0])
	suite.Require().True(rewards.IsZero())
}

func (suite *KeeperTestSuite) TestAllocateRewards_RatioPlanOverBalances() {
	farmingPoolAcc := simapp.AddTestAddrs(suite.app, suite.ctx, 1, sdk.ZeroInt())[0]
	err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, farmingPoolAcc, sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000000)))
	suite.Require().NoError(err)

	// The sum of epoch ratios is over 1, so the reward allocation should never happen.
	suite.CreateRatioPlan(farmingPoolAcc, map[string]string{denom1: "1"}, "0.8")
	suite.CreateRatioPlan(farmingPoolAcc, map[string]string{denom2: "1"}, "0.5")

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000), sdk.NewInt64Coin(denom2, 1000000)))

	suite.AdvanceEpoch()
	suite.AdvanceEpoch()

	rewards := suite.keeper.AllRewards(suite.ctx, suite.addrs[0])
	suite.Require().True(rewards.IsZero())
}

func (suite *KeeperTestSuite) TestOutstandingRewards() {
	// The block time here is not important, and has chosen randomly.
	suite.ctx = suite.ctx.WithBlockTime(types.ParseTime("2021-09-01T00:00:00Z"))

	suite.CreateFixedAmountPlan(suite.addrs[4], map[string]string{denom1: "1"}, map[string]int64{denom3: 1000})

	// Three farmers stake same amount of coins.
	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))
	suite.Stake(suite.addrs[1], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))
	suite.Stake(suite.addrs[2], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))

	// At first, the outstanding rewards shouldn't exist.
	_, found := suite.keeper.GetOutstandingRewards(suite.ctx, denom1)
	suite.Require().False(found)

	suite.AdvanceEpoch() // Queued staking coins have now staked.
	suite.AdvanceEpoch() // Allocate rewards for staked coins.

	// After the first allocation of rewards, the outstanding rewards should be 1000denom3.
	outstanding, found := suite.keeper.GetOutstandingRewards(suite.ctx, denom1)
	suite.Require().True(found)
	suite.Require().True(decCoinsEq(sdk.NewDecCoins(sdk.NewInt64DecCoin(denom3, 1000)), outstanding.Rewards))

	// All farmers harvest rewards, so the outstanding rewards should be (approximately)0.
	suite.Harvest(suite.addrs[0], []string{denom1})
	suite.Harvest(suite.addrs[1], []string{denom1})
	suite.Harvest(suite.addrs[2], []string{denom1})

	outstanding, _ = suite.keeper.GetOutstandingRewards(suite.ctx, denom1)
	truncatedOutstanding, _ := outstanding.Rewards.TruncateDecimal()
	suite.Require().True(truncatedOutstanding.IsZero())
}

func (suite *KeeperTestSuite) TestHarvest() {
	for _, plan := range suite.samplePlans {
		suite.keeper.SetPlan(suite.ctx, plan)
	}

	err := suite.keeper.Harvest(suite.ctx, suite.addrs[0], []string{denom1})
	suite.Require().EqualError(types.ErrStakingNotExists, err.Error())

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1_000_000)))
	suite.keeper.ProcessQueuedCoins(suite.ctx)

	balancesBefore := suite.app.BankKeeper.GetAllBalances(suite.ctx, suite.addrs[0])
	suite.ctx = suite.ctx.WithBlockTime(types.ParseTime("2021-08-05T00:00:00Z"))
	err = suite.keeper.AllocateRewards(suite.ctx)
	suite.Require().NoError(err)

	rewards := suite.keeper.AllRewards(suite.ctx, suite.addrs[0])

	err = suite.keeper.Harvest(suite.ctx, suite.addrs[0], []string{denom1})
	suite.Require().NoError(err)

	balancesAfter := suite.app.BankKeeper.GetAllBalances(suite.ctx, suite.addrs[0])
	suite.Require().True(coinsEq(balancesBefore.Add(rewards...), balancesAfter))
	suite.Require().True(suite.app.BankKeeper.GetAllBalances(suite.ctx, types.RewardsReserveAcc).IsZero())
	suite.Require().True(suite.keeper.AllRewards(suite.ctx, suite.addrs[0]).IsZero())
}

func (suite *KeeperTestSuite) TestMultipleHarvest() {
	suite.CreateFixedAmountPlan(suite.addrs[4], map[string]string{denom1: "1"}, map[string]int64{denom3: 1000000})

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))

	suite.AdvanceEpoch()
	suite.AdvanceEpoch()

	balancesBefore := suite.app.BankKeeper.GetAllBalances(suite.ctx, suite.addrs[0])
	suite.Harvest(suite.addrs[0], []string{denom1})
	balancesAfter := suite.app.BankKeeper.GetAllBalances(suite.ctx, suite.addrs[0])
	delta := balancesAfter.Sub(balancesBefore)
	suite.Require().True(coinsEq(sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000000)), delta))

	balancesBefore = balancesAfter
	suite.Harvest(suite.addrs[0], []string{denom1})
	balancesAfter = suite.app.BankKeeper.GetAllBalances(suite.ctx, suite.addrs[0])
	suite.Require().True(coinsEq(balancesBefore, balancesAfter))
}

func (suite *KeeperTestSuite) TestHistoricalRewards() {
	suite.ctx = suite.ctx.WithBlockTime(types.ParseTime("2021-08-06T00:00:00Z"))

	// Create two plans that share same staking coin denom in their staking coin weights.
	suite.CreateFixedAmountPlan(suite.addrs[4], map[string]string{denom1: "1"}, map[string]int64{denom3: 1000000})
	suite.CreateFixedAmountPlan(suite.addrs[4], map[string]string{denom1: "1"}, map[string]int64{denom3: 1000000})

	// Advancing epoch(s) before any staking is made doesn't create any historical rewards records.
	suite.AdvanceEpoch()
	suite.AdvanceEpoch()
	count := 0
	suite.keeper.IterateHistoricalRewards(suite.ctx, func(stakingCoinDenom string, epoch uint64, rewards types.HistoricalRewards) (stop bool) {
		count++
		return false
	})
	suite.Equal(count, 0)

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))
	// Advancing epoch here marks queued staking coins as staked.
	suite.AdvanceEpoch()

	// After a farmer has staked(not queued) coins, historical rewards records will be created for each epoch.
	// Here we advance epoch three times, and this will create 3 historical rewards records.
	suite.AdvanceEpoch()
	suite.AdvanceEpoch()
	suite.AdvanceEpoch()

	// First, ensure that we have only 3 entries in the store.
	count = 0
	suite.keeper.IterateHistoricalRewards(suite.ctx, func(stakingCoinDenom string, epoch uint64, rewards types.HistoricalRewards) (stop bool) {
		count++
		return false
	})
	suite.Require().Equal(4, count)

	// Next, check if cumulative unit rewards is correct.
	for i := uint64(1); i <= 3; i++ {
		historical, found := suite.keeper.GetHistoricalRewards(suite.ctx, denom1, i)
		suite.Require().True(found)
		suite.Require().True(decCoinsEq(sdk.NewDecCoins(sdk.NewInt64DecCoin(denom3, int64(i*2))), historical.CumulativeUnitRewards))
	}
}

// Test if initialization and pruning of staking coin info work properly.
func (suite *KeeperTestSuite) TestInitializeAndPruneStakingCoinInfo() {
	suite.CreateFixedAmountPlan(suite.addrs[4], map[string]string{denom1: "1"}, map[string]int64{denom3: 1000000})

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))

	suite.Require().Equal(uint64(0), suite.keeper.GetCurrentEpoch(suite.ctx, denom1))
	_, found := suite.keeper.GetHistoricalRewards(suite.ctx, denom1, 0)
	suite.Require().False(found)
	_, found = suite.keeper.GetHistoricalRewards(suite.ctx, denom1, 1)
	suite.Require().False(found)
	_, found = suite.keeper.GetOutstandingRewards(suite.ctx, denom1)
	suite.Require().False(found)

	suite.AdvanceEpoch()

	suite.Require().Equal(uint64(1), suite.keeper.GetCurrentEpoch(suite.ctx, denom1))
	historical, found := suite.keeper.GetHistoricalRewards(suite.ctx, denom1, 0)
	suite.Require().True(found)
	suite.Require().True(decCoinsEq(sdk.DecCoins{}, historical.CumulativeUnitRewards))
	outstanding, found := suite.keeper.GetOutstandingRewards(suite.ctx, denom1)
	suite.Require().True(found)
	suite.Require().True(decCoinsEq(sdk.DecCoins{}, outstanding.Rewards))

	suite.AdvanceEpoch()

	suite.Require().Equal(uint64(2), suite.keeper.GetCurrentEpoch(suite.ctx, denom1))
	historical, found = suite.keeper.GetHistoricalRewards(suite.ctx, denom1, 1)
	suite.Require().True(found)
	suite.Require().True(decCoinsEq(sdk.NewDecCoins(sdk.NewInt64DecCoin(denom3, 1)), historical.CumulativeUnitRewards))
	outstanding, found = suite.keeper.GetOutstandingRewards(suite.ctx, denom1)
	suite.Require().True(found)
	suite.Require().True(decCoinsEq(sdk.NewDecCoins(sdk.NewInt64DecCoin(denom3, 1000000)), outstanding.Rewards))
	// Historical rewards for epoch 2 must not be present at this point,
	// since current epoch is 2, and it has not ended yet.
	_, found = suite.keeper.GetHistoricalRewards(suite.ctx, denom1, 2)
	suite.Require().False(found)

	// Unstake most of the coins. This should not delete any info
	// about the staking coin yet.
	suite.Unstake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 999999)))
	suite.Require().Equal(uint64(2), suite.keeper.GetCurrentEpoch(suite.ctx, denom1))
	_, found = suite.keeper.GetHistoricalRewards(suite.ctx, denom1, 1)
	suite.Require().True(found)
	_, found = suite.keeper.GetOutstandingRewards(suite.ctx, denom1)
	suite.Require().True(found)

	// Now unstake the rest of the coins. This should delete info
	// about the staking coin.
	suite.Unstake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1)))
	farming.EndBlocker(suite.ctx, suite.keeper)
	_, found = suite.keeper.GetHistoricalRewards(suite.ctx, denom1, 1)
	suite.Require().False(found)
	_, found = suite.keeper.GetOutstandingRewards(suite.ctx, denom1)
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestAllocateRewardsZeroTotalStakings() {
	suite.CreateFixedAmountPlan(suite.addrs[4], map[string]string{denom1: "1"}, map[string]int64{denom3: 1000000})

	suite.Stake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))
	suite.AdvanceEpoch()
	suite.AdvanceEpoch()

	suite.Unstake(suite.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))
	suite.AdvanceEpoch() // This should not cause a panic.
	totalStakings, found := suite.keeper.GetTotalStakings(suite.ctx, denom1)
	suite.Require().True(found)
	suite.Require().True(totalStakings.Amount.IsZero())
	farming.EndBlocker(suite.ctx, suite.keeper)

	_, found = suite.keeper.GetTotalStakings(suite.ctx, denom1)
	suite.Require().False(found)
}
