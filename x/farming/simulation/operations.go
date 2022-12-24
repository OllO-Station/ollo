package simulation

// Simulation operation weights constants.
// const (
// 	OpWeightMsgCreateFixedAmountPlan = "op_weight_msg_create_fixed_amount_plan"
// 	OpWeightMsgCreateRatioPlan       = "op_weight_msg_create_ratio_plan"
// 	OpWeightMsgStake                 = "op_weight_msg_stake"
// 	OpWeightMsgUnstake               = "op_weight_msg_unstake"
// 	OpWeightMsgHarvest               = "op_weight_msg_harvest"
// 	OpWeightMsgRemovePlan            = "op_weight_msg_remove_plan"
// )

// var (
// 	poolCoinDenoms = []string{
// 		"pool93E069B333B5ECEBFE24C6E1437E814003248E0DD7FF8B9F82119F4587449BA5",
// 		"pool3036F43CB8131A1A63D2B3D3B11E9CF6FA2A2B6FEC17D5AD283C25C939614A8C",
// 		"poolE4D2617BFE03E1146F6BBA1D9893F2B3D77BA29E7ED532BB721A39FF1ECC1B07",
// 	}

// 	testCoinDenoms = []string{
// 		"testa",
// 		"testb",
// 		"testc",
// 	}
// )

// func init() {
// 	keeper.EnableRatioPlan = true
// }

// // WeightedOperations returns all the operations from the module with their respective weights.
// func WeightedOperations(
// 	appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper,
// 	bk types.BankKeeper, k keeper.Keeper,
// ) simulation.WeightedOperations {

// 	var weightMsgCreateFixedAmountPlan int
// 	appParams.GetOrGenerate(cdc, OpWeightMsgCreateFixedAmountPlan, &weightMsgCreateFixedAmountPlan, nil,
// 		func(_ *rand.Rand) {
// 			weightMsgCreateFixedAmountPlan = params.DefaultWeightMsgCreateFixedAmountPlan
// 		},
// 	)

// 	var weightMsgCreateRatioPlan int
// 	appParams.GetOrGenerate(cdc, OpWeightMsgCreateRatioPlan, &weightMsgCreateRatioPlan, nil,
// 		func(_ *rand.Rand) {
// 			weightMsgCreateRatioPlan = params.DefaultWeightMsgCreateRatioPlan
// 		},
// 	)

// 	var weightMsgStake int
// 	appParams.GetOrGenerate(cdc, OpWeightMsgStake, &weightMsgStake, nil,
// 		func(_ *rand.Rand) {
// 			weightMsgStake = params.DefaultWeightMsgStake
// 		},
// 	)

// 	var weightMsgUnstake int
// 	appParams.GetOrGenerate(cdc, OpWeightMsgUnstake, &weightMsgUnstake, nil,
// 		func(_ *rand.Rand) {
// 			weightMsgUnstake = params.DefaultWeightMsgUnstake
// 		},
// 	)

// 	var weightMsgHarvest int
// 	appParams.GetOrGenerate(cdc, OpWeightMsgHarvest, &weightMsgHarvest, nil,
// 		func(_ *rand.Rand) {
// 			weightMsgHarvest = params.DefaultWeightMsgHarvest
// 		},
// 	)

// 	var weightMsgRemovePlan int
// 	appParams.GetOrGenerate(cdc, OpWeightMsgRemovePlan, &weightMsgRemovePlan, nil,
// 		func(r *rand.Rand) {
// 			weightMsgRemovePlan = params.DefaultWeightMsgRemovePlan
// 		},
// 	)

// 	return simulation.WeightedOperations{
// 		simulation.NewWeightedOperation(
// 			weightMsgCreateFixedAmountPlan,
// 			SimulateMsgCreateFixedAmountPlan(ak, bk, k),
// 		),
// 		simulation.NewWeightedOperation(
// 			weightMsgCreateRatioPlan,
// 			SimulateMsgCreateRatioPlan(ak, bk, k),
// 		),
// 		simulation.NewWeightedOperation(
// 			weightMsgStake,
// 			SimulateMsgStake(ak, bk, k),
// 		),
// 		simulation.NewWeightedOperation(
// 			weightMsgUnstake,
// 			SimulateMsgUnstake(ak, bk, k),
// 		),
// 		simulation.NewWeightedOperation(
// 			weightMsgHarvest,
// 			SimulateMsgHarvest(ak, bk, k),
// 		),
// 		simulation.NewWeightedOperation(
// 			weightMsgRemovePlan,
// 			SimulateMsgRemovePlan(ak, bk, k),
// 		),
// 	}
// }

// // SimulateMsgCreateFixedAmountPlan generates a MsgCreateFixedAmountPlan with random values
// // nolint: interfacer
// func SimulateMsgCreateFixedAmountPlan(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
// 	return func(
// 		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		simAccount, _ := simtypes.RandomAcc(r, accs)

// 		account := ak.GetAccount(ctx, simAccount.Address)
// 		spendable := bk.SpendableCoins(ctx, account.GetAddress())

// 		params := k.GetParams(ctx)
// 		if uint32(k.GetNumActivePrivatePlans(ctx)) > params.MaxNumPrivatePlans {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateFixedAmountPlan, "maximum number of private plans reached"), nil, nil
// 		}

// 		_, hasNeg := spendable.SafeSub(params.PrivatePlanCreationFee)
// 		if hasNeg {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateFixedAmountPlan, "insufficient balance for plan creation fee"), nil, nil
// 		}

// 		name := "simulation-test-" + simtypes.RandStringOfLength(r, 5) // name must be unique
// 		creatorAcc := account.GetAddress()
// 		// mint pool coins to simulate the real-world cases
// 		funds, err := fundBalances(ctx, r, bk, creatorAcc, testCoinDenoms)
// 		if err != nil {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateFixedAmountPlan, "unable to mint pool coins"), nil, nil
// 		}
// 		stakingCoinWeights := sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 1))
// 		startTime := ctx.BlockTime()
// 		endTime := startTime.AddDate(1, 0, 0)
// 		epochAmount := sdk.NewCoins(
// 			sdk.NewInt64Coin(funds[r.Intn(3)].Denom, int64(simtypes.RandIntBetween(r, 10_000_000, 1_000_000_000))),
// 		)

// 		msg := types.NewMsgCreateFixedAmountPlan(
// 			name,
// 			creatorAcc,
// 			stakingCoinWeights,
// 			startTime,
// 			endTime,
// 			epochAmount,
// 		)

// 		txCtx := simulation.OperationInput{
// 			R:               r,
// 			App:             app,
// 			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
// 			Cdc:             nil,
// 			Msg:             msg,
// 			MsgType:         msg.Type(),
// 			Context:         ctx,
// 			SimAccount:      simAccount,
// 			AccountKeeper:   ak,
// 			Bankkeeper:      bk,
// 			ModuleName:      types.ModuleName,
// 			CoinsSpentInMsg: spendable,
// 		}

// 		return simulation.GenAndDeliverTxWithRandFees(txCtx)
// 	}
// }

// // SimulateMsgCreateRatioPlan generates a MsgCreateRatioPlan with random values
// // nolint: interfacer
// func SimulateMsgCreateRatioPlan(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
// 	return func(
// 		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		simAccount, _ := simtypes.RandomAcc(r, accs)

// 		account := ak.GetAccount(ctx, simAccount.Address)
// 		spendable := bk.SpendableCoins(ctx, account.GetAddress())

// 		params := k.GetParams(ctx)
// 		if uint32(k.GetNumActivePrivatePlans(ctx)) > params.MaxNumPrivatePlans {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateRatioPlan, "maximum number of private plans reached"), nil, nil
// 		}

// 		_, hasNeg := spendable.SafeSub(params.PrivatePlanCreationFee)
// 		if hasNeg {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateRatioPlan, "insufficient balance for plan creation fee"), nil, nil
// 		}

// 		name := "simulation-test-" + simtypes.RandStringOfLength(r, 5) // name must be unique
// 		creatorAcc := account.GetAddress()
// 		// mint pool coins to simulate the real-world cases
// 		_, err := fundBalances(ctx, r, bk, account.GetAddress(), testCoinDenoms)
// 		if err != nil {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateRatioPlan, "unable to mint pool coins"), nil, nil
// 		}
// 		stakingCoinWeights := sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 1))
// 		startTime := ctx.BlockTime()
// 		endTime := startTime.AddDate(1, 0, 0)
// 		epochRatio := sdk.NewDecWithPrec(int64(simtypes.RandIntBetween(r, 1, 10)), 3)

// 		msg := types.NewMsgCreateRatioPlan(
// 			name,
// 			creatorAcc,
// 			stakingCoinWeights,
// 			startTime,
// 			endTime,
// 			epochRatio,
// 		)

// 		txCtx := simulation.OperationInput{
// 			R:               r,
// 			App:             app,
// 			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
// 			Cdc:             nil,
// 			Msg:             msg,
// 			MsgType:         msg.Type(),
// 			Context:         ctx,
// 			SimAccount:      simAccount,
// 			AccountKeeper:   ak,
// 			Bankkeeper:      bk,
// 			ModuleName:      types.ModuleName,
// 			CoinsSpentInMsg: spendable,
// 		}

// 		return simulation.GenAndDeliverTxWithRandFees(txCtx)
// 	}
// }

// // SimulateMsgStake generates a MsgStake with random values
// // nolint: interfacer
// func SimulateMsgStake(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
// 	return func(
// 		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		simAccount, _ := simtypes.RandomAcc(r, accs)

// 		account := ak.GetAccount(ctx, simAccount.Address)
// 		spendable := bk.SpendableCoins(ctx, account.GetAddress())

// 		farmer := account.GetAddress()
// 		stakingCoins := sdk.NewCoins(
// 			sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simtypes.RandIntBetween(r, 1_000_000, 1_000_000_000))),
// 		)
// 		if !spendable.IsAllGTE(stakingCoins) {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnstake, "insufficient funds"), nil, nil
// 		}

// 		msg := types.NewMsgStake(farmer, stakingCoins)
// 		txCtx := simulation.OperationInput{
// 			R:               r,
// 			App:             app,
// 			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
// 			Cdc:             nil,
// 			Msg:             msg,
// 			MsgType:         msg.Type(),
// 			Context:         ctx,
// 			SimAccount:      simAccount,
// 			AccountKeeper:   ak,
// 			Bankkeeper:      bk,
// 			ModuleName:      types.ModuleName,
// 			CoinsSpentInMsg: spendable,
// 		}

// 		return simulation.GenAndDeliverTxWithRandFees(txCtx)
// 	}
// }

// // SimulateMsgUnstake generates a SimulateMsgUnstake with random values
// // nolint: interfacer
// func SimulateMsgUnstake(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
// 	return func(
// 		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		simAccount, _ := simtypes.RandomAcc(r, accs)

// 		account := ak.GetAccount(ctx, simAccount.Address)
// 		spendable := bk.SpendableCoins(ctx, account.GetAddress())

// 		farmer := account.GetAddress()
// 		unstakingCoins := sdk.NewCoins(
// 			sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simtypes.RandIntBetween(r, 1_000_000, 100_000_000))),
// 		)

// 		// staking must exist in order to unstake
// 		staking, sf := k.GetStaking(ctx, sdk.DefaultBondDenom, farmer)
// 		if !sf {
// 			staking = types.Staking{
// 				Amount: sdk.ZeroInt(),
// 			}
// 		}
// 		queuedStaking, qsf := k.GetQueuedStaking(ctx, sdk.DefaultBondDenom, farmer)
// 		if !qsf {
// 			if !qsf {
// 				queuedStaking = types.QueuedStaking{
// 					Amount: sdk.ZeroInt(),
// 				}
// 			}
// 		}
// 		if !sf && !qsf {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnstake, "unable to find staking and queued staking"), nil, nil
// 		}
// 		// sum of staked and queued coins must be greater than unstaking coins
// 		if !staking.Amount.Add(queuedStaking.Amount).GTE(unstakingCoins[0].Amount) {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnstake, "insufficient funds"), nil, nil
// 		}

// 		// spendable must be greater than unstaking coins
// 		if !spendable.IsAllGT(unstakingCoins) {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnstake, "insufficient funds"), nil, nil
// 		}

// 		msg := types.NewMsgUnstake(farmer, unstakingCoins)
// 		txCtx := simulation.OperationInput{
// 			R:               r,
// 			App:             app,
// 			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
// 			Cdc:             nil,
// 			Msg:             msg,
// 			MsgType:         msg.Type(),
// 			Context:         ctx,
// 			SimAccount:      simAccount,
// 			AccountKeeper:   ak,
// 			Bankkeeper:      bk,
// 			ModuleName:      types.ModuleName,
// 			CoinsSpentInMsg: spendable,
// 		}
// 		return simulation.GenAndDeliverTxWithRandFees(txCtx)
// 	}
// }

// // SimulateMsgHarvest generates a MsgHarvest with random values
// // nolint: interfacer
// func SimulateMsgHarvest(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
// 	return func(
// 		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		var simAccount simtypes.Account
// 		var stakingCoinDenoms []string

// 		skip := true
// 		// find staking from the simulated accounts
// 		for _, acc := range accs {
// 			staked := k.GetAllStakedCoinsByFarmer(ctx, acc.Address)
// 			stakingCoinDenoms = nil
// 			for _, coin := range staked {
// 				rewards := k.Rewards(ctx, acc.Address, coin.Denom)
// 				if !rewards.IsZero() {
// 					stakingCoinDenoms = append(stakingCoinDenoms, coin.Denom)
// 				}
// 			}
// 			if len(stakingCoinDenoms) > 0 {
// 				simAccount = acc
// 				skip = false
// 				break
// 			}
// 		}
// 		if skip {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgHarvest, "no account to harvest rewards"), nil, nil
// 		}

// 		account := ak.GetAccount(ctx, simAccount.Address)
// 		spendable := bk.SpendableCoins(ctx, account.GetAddress())

// 		msg := types.NewMsgHarvest(simAccount.Address, stakingCoinDenoms)

// 		txCtx := simulation.OperationInput{
// 			R:               r,
// 			App:             app,
// 			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
// 			Cdc:             nil,
// 			Msg:             msg,
// 			MsgType:         msg.Type(),
// 			Context:         ctx,
// 			SimAccount:      simAccount,
// 			AccountKeeper:   ak,
// 			Bankkeeper:      bk,
// 			ModuleName:      types.ModuleName,
// 			CoinsSpentInMsg: spendable,
// 		}

// 		return simulation.GenAndDeliverTxWithRandFees(txCtx)
// 	}
// }

// // SimulateMsgRemovePlan generates a MsgRemovePlan with random values
// func SimulateMsgRemovePlan(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
// 	return func(
// 		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		simAccount, _ := simtypes.RandomAcc(r, accs)

// 		account := ak.GetAccount(ctx, simAccount.Address)
// 		spendable := bk.SpendableCoins(ctx, account.GetAddress())

// 		creator := account.GetAddress()

// 		var terminatedPlans []types.PlanI
// 		for _, plan := range k.GetPlans(ctx) {
// 			if plan.IsTerminated() {
// 				terminatedPlans = append(terminatedPlans, plan)
// 			}
// 		}
// 		if len(terminatedPlans) == 0 {
// 			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemovePlan, "no terminated plans to remove"), nil, nil
// 		}

// 		// Select a random terminated plan.
// 		plan := terminatedPlans[simtypes.RandIntBetween(r, 0, len(terminatedPlans))]

// 		msg := types.NewMsgRemovePlan(creator, plan.GetId())
// 		txCtx := simulation.OperationInput{
// 			R:               r,
// 			App:             app,
// 			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
// 			Cdc:             nil,
// 			Msg:             msg,
// 			MsgType:         msg.Type(),
// 			Context:         ctx,
// 			SimAccount:      simAccount,
// 			AccountKeeper:   ak,
// 			Bankkeeper:      bk,
// 			ModuleName:      types.ModuleName,
// 			CoinsSpentInMsg: spendable,
// 		}

// 		return simulation.GenAndDeliverTxWithRandFees(txCtx)
// 	}
// }

// // fundBalances mints random amount of coins with the provided coin denoms and
// // send them to the simulated account.
// func fundBalances(ctx sdk.Context, r *rand.Rand, bk types.BankKeeper, acc sdk.AccAddress, denoms []string) (mintCoins sdk.Coins, err error) {
// 	for _, denom := range denoms {
// 		mintCoins = mintCoins.Add(sdk.NewInt64Coin(denom, int64(simtypes.RandIntBetween(r, 1e14, 1e15))))
// 	}

// 	if err := bk.MintCoins(ctx, minttypes.ModuleName, mintCoins); err != nil {
// 		return nil, err
// 	}

// 	if err := bk.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, acc, mintCoins); err != nil {
// 		return nil, err
// 	}
// 	return mintCoins, nil
// }
