package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"ollo/x/liquidity/types"
)

func (k Keeper) ValidateMsgCreatePool(ctx sdk.Context, msg *types.MsgCreatePool) error {
	params := k.GetParams(ctx)
	var poolType types.PoolType

	// check poolType exist, get poolType from param
	if len(params.PoolTypes) >= int(msg.PoolTypeId) {
		poolType = params.PoolTypes[msg.PoolTypeId-1]
		if poolType.Id != msg.PoolTypeId {
			return types.ErrPoolTypeNotExists
		}
	} else {
		return types.ErrPoolTypeNotExists
	}

	reserveCoinNum := uint32(msg.DepositCoins.Len())
	if reserveCoinNum > poolType.MaxReserveCoinNum || poolType.MinReserveCoinNum > reserveCoinNum {
		return types.ErrNumOfReserveCoin
	}

	reserveDenoms := make([]string, reserveCoinNum)
	for i := 0; i < int(reserveCoinNum); i++ {
		reserveDenoms[i] = msg.DepositCoins.GetDenomByIndex(i)
	}

	denomA, denomB := types.AlphabeticalDenomPair(reserveDenoms[0], reserveDenoms[1])
	if denomA != msg.DepositCoins[0].Denom || denomB != msg.DepositCoins[1].Denom {
		return types.ErrBadOrderingReserveCoin
	}

	if denomA == denomB {
		return types.ErrEqualDenom
	}

	if err := types.ValidateReserveCoinLimit(params.MaxReserveCoinAmount, msg.DepositCoins); err != nil {
		return err
	}

	poolName := types.PoolName(reserveDenoms, msg.PoolTypeId)
	reserveAcc := types.GetPoolReserveAcc(poolName, false)
	_, found := k.GetPoolByReserveAccIndex(ctx, reserveAcc)
	if found {
		return types.ErrPoolAlreadyExists
	}
	return nil
}

func (k Keeper) MintAndSendPoolCoin(ctx sdk.Context, pool types.Pool, srcAddr, creatorAddr sdk.AccAddress, depositCoins sdk.Coins) (sdk.Coin, error) {
	cacheCtx, writeCache := ctx.CacheContext()

	params := k.GetParams(cacheCtx)

	mintingCoin := sdk.NewCoin(pool.PoolCoinDenom, params.InitPoolCoinMintAmount)
	mintingCoins := sdk.NewCoins(mintingCoin)
	if err := k.bankKeeper.MintCoins(cacheCtx, types.ModuleName, mintingCoins); err != nil {
		return sdk.Coin{}, err
	}

	reserveAcc := pool.GetReserveAccount()

	var inputs []banktypes.Input
	var outputs []banktypes.Output

	inputs = append(inputs, banktypes.NewInput(srcAddr, depositCoins))
	outputs = append(outputs, banktypes.NewOutput(reserveAcc, depositCoins))

	inputs = append(inputs, banktypes.NewInput(k.accountKeeper.GetModuleAddress(types.ModuleName), mintingCoins))
	outputs = append(outputs, banktypes.NewOutput(creatorAddr, mintingCoins))

	if err := k.bankKeeper.InputOutputCoins(cacheCtx, inputs, outputs); err != nil {
		return sdk.Coin{}, err
	}

	writeCache()

	return mintingCoin, nil
}

func (k Keeper) CreatePool(ctx sdk.Context, msg *types.MsgCreatePool) (types.Pool, error) {
	if err := k.ValidateMsgCreatePool(ctx, msg); err != nil {
		return types.Pool{}, err
	}

	params := k.GetParams(ctx)

	denom1, denom2 := types.AlphabeticalDenomPair(msg.DepositCoins[0].Denom, msg.DepositCoins[1].Denom)
	reserveDenoms := []string{denom1, denom2}

	poolName := types.PoolName(reserveDenoms, msg.PoolTypeId)

	pool := types.Pool{
		// Id: will set on SetPoolAtomic
		TypeId:                msg.PoolTypeId,
		ReserveCoinDenoms:     reserveDenoms,
		ReserveAccountAddress: types.GetPoolReserveAcc(poolName, false).String(),
		PoolCoinDenom:         types.GetPoolCoinDenom(poolName),
	}

	poolCreator := msg.GetPoolCreator()

	for _, coin := range msg.DepositCoins {
		if coin.Amount.LT(params.MinInitDepositAmount) {
			return types.Pool{}, sdkerrors.Wrapf(
				types.ErrLessThanMinInitDeposit, "deposit coin %s is smaller than %s", coin, params.MinInitDepositAmount)
		}
	}

	for _, coin := range msg.DepositCoins {
		balance := k.bankKeeper.GetBalance(ctx, poolCreator, coin.Denom)
		if balance.IsLT(coin) {
			return types.Pool{}, sdkerrors.Wrapf(
				types.ErrInsufficientBalance, "%s is smaller than %s", balance, coin)
		}
	}

	for _, coin := range params.PoolCreationFee {
		balance := k.bankKeeper.GetBalance(ctx, poolCreator, coin.Denom)
		neededAmt := coin.Amount.Add(msg.DepositCoins.AmountOf(coin.Denom))
		neededCoin := sdk.NewCoin(coin.Denom, neededAmt)
		if balance.IsLT(neededCoin) {
			return types.Pool{}, sdkerrors.Wrapf(
				types.ErrInsufficientPoolCreationFee, "%s is smaller than %s", balance, neededCoin)
		}
	}

	if _, err := k.MintAndSendPoolCoin(ctx, pool, poolCreator, poolCreator, msg.DepositCoins); err != nil {
		return types.Pool{}, err
	}

	// pool creation fees are collected in community pool
	if err := k.distrKeeper.FundCommunityPool(ctx, params.PoolCreationFee, poolCreator); err != nil {
		return types.Pool{}, err
	}

	pool = k.SetPoolAtomic(ctx, pool)
	batch := types.NewPoolBatch(pool.Id, 1)
	batch.BeginHeight = ctx.BlockHeight()

	k.SetPoolBatch(ctx, batch)

	reserveCoins := k.GetReserveCoins(ctx, pool)
	lastReserveRatio := sdk.NewDecFromInt(reserveCoins[0].Amount).Quo(sdk.NewDecFromInt(reserveCoins[1].Amount))
	logger := k.Logger(ctx)
	logger.Debug(
		"create liquidity pool",
		"msg", msg,
		"pool", pool,
		"reserveCoins", reserveCoins,
		"lastReserveRatio", lastReserveRatio,
	)

	return pool, nil
}

func (k Keeper) ExecuteDeposit(ctx sdk.Context, msg types.DepositMsgState, batch types.PoolBatch) error {
	if msg.Executed || msg.ToBeDeleted || msg.Succeeded {
		return fmt.Errorf("cannot process already executed batch msg")
	}
	msg.Executed = true
	k.SetPoolBatchDepositMsgState(ctx, msg.Msg.PoolId, msg)

	if err := k.ValidateMsgDepositWithinBatch(ctx, *msg.Msg); err != nil {
		return err
	}

	pool, found := k.GetPool(ctx, msg.Msg.PoolId)
	if !found {
		return types.ErrPoolNotExists
	}

	depositCoins := msg.Msg.DepositCoins.Sort()

	batchEscrowAcc := k.accountKeeper.GetModuleAddress(types.ModuleName)
	reserveAcc := pool.GetReserveAccount()
	depositor := msg.Msg.GetDepositor()

	params := k.GetParams(ctx)

	reserveCoins := k.GetReserveCoins(ctx, pool)

	// reinitialize pool if the pool is depleted
	if k.IsDepletedPool(ctx, pool) {
		for _, depositCoin := range msg.Msg.DepositCoins {
			if depositCoin.Amount.Add(reserveCoins.AmountOf(depositCoin.Denom)).LT(params.MinInitDepositAmount) {
				return types.ErrLessThanMinInitDeposit
			}
		}
		poolCoin, err := k.MintAndSendPoolCoin(ctx, pool, batchEscrowAcc, depositor, msg.Msg.DepositCoins)
		if err != nil {
			return err
		}

		// set deposit msg state of the pool batch complete
		msg.Succeeded = true
		msg.ToBeDeleted = true
		k.SetPoolBatchDepositMsgState(ctx, msg.Msg.PoolId, msg)

		reserveCoins = k.GetReserveCoins(ctx, pool)
		lastReserveCoinA := sdk.NewDecFromInt(reserveCoins[0].Amount)
		lastReserveCoinB := sdk.NewDecFromInt(reserveCoins[1].Amount)
		lastReserveRatio := lastReserveCoinA.Quo(lastReserveCoinB)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeDepositToPool,
				sdk.NewAttribute(types.AttributeValuePoolId, strconv.FormatUint(pool.Id, 10)),
				sdk.NewAttribute(types.AttributeValueBatchIndex, strconv.FormatUint(batch.Index, 10)),
				sdk.NewAttribute(types.AttributeValueMsgIndex, strconv.FormatUint(msg.MsgIndex, 10)),
				sdk.NewAttribute(types.AttributeValueDepositor, depositor.String()),
				sdk.NewAttribute(types.AttributeValueAcceptedCoins, msg.Msg.DepositCoins.String()),
				sdk.NewAttribute(types.AttributeValueRefundedCoins, ""),
				sdk.NewAttribute(types.AttributeValuePoolDenom, poolCoin.Denom),
				sdk.NewAttribute(types.AttributeValuePoolCoinAmount, poolCoin.Amount.String()),
				sdk.NewAttribute(types.AttributeValueSuccess, types.Success),
			),
		)
		logger := k.Logger(ctx)
		logger.Debug(
			"reinitialize pool",
			"msg", msg,
			"pool", pool,
			"reserveCoins", reserveCoins,
			"lastReserveRatio", lastReserveRatio,
		)

		return nil
	}

	reserveCoins.Sort()

	lastReserveCoinA := reserveCoins[0]
	lastReserveCoinB := reserveCoins[1]

	depositCoinA := depositCoins[0]
	depositCoinB := depositCoins[1]

	poolCoinTotalSupply := sdk.NewDecFromInt(k.GetPoolCoinTotalSupply(ctx, pool))
	if err := types.CheckOverflowWithDec(poolCoinTotalSupply, sdk.NewDecFromInt(depositCoinA.Amount)); err != nil {
		return err
	}
	if err := types.CheckOverflowWithDec(poolCoinTotalSupply, sdk.NewDecFromInt(depositCoinB.Amount)); err != nil {
		return err
	}
	poolCoinMintAmt := sdk.MinDec(
		poolCoinTotalSupply.MulTruncate(sdk.NewDecFromInt(depositCoinA.Amount)).QuoTruncate(sdk.NewDecFromInt(lastReserveCoinA.Amount)),
		poolCoinTotalSupply.MulTruncate(sdk.NewDecFromInt(depositCoinB.Amount)).QuoTruncate(sdk.NewDecFromInt(lastReserveCoinB.Amount)),
	)
	mintRate := poolCoinMintAmt.TruncateDec().QuoTruncate(poolCoinTotalSupply)
	acceptedCoins := sdk.NewCoins(
		sdk.NewCoin(depositCoins[0].Denom, sdk.NewDecFromInt(lastReserveCoinA.Amount).Mul(mintRate).TruncateInt()),
		sdk.NewCoin(depositCoins[1].Denom, sdk.NewDecFromInt(lastReserveCoinB.Amount).Mul(mintRate).TruncateInt()),
	)
	refundedCoins := depositCoins.Sub(acceptedCoins...)
	refundedCoinA := sdk.NewCoin(depositCoinA.Denom, refundedCoins.AmountOf(depositCoinA.Denom))
	refundedCoinB := sdk.NewCoin(depositCoinB.Denom, refundedCoins.AmountOf(depositCoinB.Denom))

	mintPoolCoin := sdk.NewCoin(pool.PoolCoinDenom, poolCoinMintAmt.TruncateInt())
	mintPoolCoins := sdk.NewCoins(mintPoolCoin)

	if mintPoolCoins.IsZero() || acceptedCoins.IsZero() {
		return fmt.Errorf("pool coin truncated, no accepted coin, refund")
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintPoolCoins); err != nil {
		return err
	}

	var inputs []banktypes.Input
	var outputs []banktypes.Output

	if !refundedCoins.IsZero() {
		// refund truncated deposit coins
		inputs = append(inputs, banktypes.NewInput(batchEscrowAcc, refundedCoins))
		outputs = append(outputs, banktypes.NewOutput(depositor, refundedCoins))
	}

	// send accepted deposit coins
	inputs = append(inputs, banktypes.NewInput(batchEscrowAcc, acceptedCoins))
	outputs = append(outputs, banktypes.NewOutput(reserveAcc, acceptedCoins))

	// send minted pool coins
	inputs = append(inputs, banktypes.NewInput(batchEscrowAcc, mintPoolCoins))
	outputs = append(outputs, banktypes.NewOutput(depositor, mintPoolCoins))

	// execute multi-send
	if err := k.bankKeeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
		return err
	}

	msg.Succeeded = true
	msg.ToBeDeleted = true
	k.SetPoolBatchDepositMsgState(ctx, msg.Msg.PoolId, msg)

	if BatchLogicInvariantCheckFlag {
		afterReserveCoins := k.GetReserveCoins(ctx, pool)
		afterReserveCoinA := afterReserveCoins[0].Amount
		afterReserveCoinB := afterReserveCoins[1].Amount

		MintingPoolCoinsInvariant(poolCoinTotalSupply.TruncateInt(), mintPoolCoin.Amount, depositCoinA.Amount, depositCoinB.Amount,
			lastReserveCoinA.Amount, lastReserveCoinB.Amount, refundedCoinA.Amount, refundedCoinB.Amount)
		DepositInvariant(lastReserveCoinA.Amount, lastReserveCoinB.Amount, depositCoinA.Amount, depositCoinB.Amount,
			afterReserveCoinA, afterReserveCoinB, refundedCoinA.Amount, refundedCoinB.Amount)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDepositToPool,
			sdk.NewAttribute(types.AttributeValuePoolId, strconv.FormatUint(pool.Id, 10)),
			sdk.NewAttribute(types.AttributeValueBatchIndex, strconv.FormatUint(batch.Index, 10)),
			sdk.NewAttribute(types.AttributeValueMsgIndex, strconv.FormatUint(msg.MsgIndex, 10)),
			sdk.NewAttribute(types.AttributeValueDepositor, depositor.String()),
			sdk.NewAttribute(types.AttributeValueAcceptedCoins, acceptedCoins.String()),
			sdk.NewAttribute(types.AttributeValueRefundedCoins, refundedCoins.String()),
			sdk.NewAttribute(types.AttributeValuePoolDenom, mintPoolCoin.Denom),
			sdk.NewAttribute(types.AttributeValuePoolCoinAmount, mintPoolCoin.Amount.String()),
			sdk.NewAttribute(types.AttributeValueSuccess, types.Success),
		),
	)

	reserveCoins = k.GetReserveCoins(ctx, pool)
	lastReserveRatio := sdk.NewDecFromInt(reserveCoins[0].Amount).Quo(sdk.NewDecFromInt(reserveCoins[1].Amount))

	logger := k.Logger(ctx)
	logger.Debug(
		"deposit coins to the pool",
		"msg", msg,
		"pool", pool,
		"inputs", inputs,
		"outputs", outputs,
		"reserveCoins", reserveCoins,
		"lastReserveRatio", lastReserveRatio,
	)

	return nil
}

// ExecuteWithdrawal withdraws pool coin from the liquidity pool
func (k Keeper) ExecuteWithdrawal(ctx sdk.Context, msg types.WithdrawMsgState, batch types.PoolBatch) error {
	if msg.Executed || msg.ToBeDeleted || msg.Succeeded {
		return fmt.Errorf("cannot process already executed batch msg")
	}
	msg.Executed = true
	k.SetPoolBatchWithdrawMsgState(ctx, msg.Msg.PoolId, msg)

	if err := k.ValidateMsgWithdrawWithinBatch(ctx, *msg.Msg); err != nil {
		return err
	}
	poolCoins := sdk.NewCoins(msg.Msg.PoolCoin)

	pool, found := k.GetPool(ctx, msg.Msg.PoolId)
	if !found {
		return types.ErrPoolNotExists
	}

	poolCoinTotalSupply := k.GetPoolCoinTotalSupply(ctx, pool)
	reserveCoins := k.GetReserveCoins(ctx, pool)
	reserveCoins.Sort()

	var inputs []banktypes.Input
	var outputs []banktypes.Output

	reserveAcc := pool.GetReserveAccount()
	withdrawer := msg.Msg.GetWithdrawer()

	params := k.GetParams(ctx)
	withdrawProportion := sdk.OneDec().Sub(params.WithdrawFeeRate)
	withdrawCoins := sdk.NewCoins()
	withdrawFeeCoins := sdk.NewCoins()

	// Case for withdrawing all reserve coins
	if msg.Msg.PoolCoin.Amount.Equal(poolCoinTotalSupply) {
		withdrawCoins = reserveCoins
	} else {
		// Calculate withdraw amount of respective reserve coin considering fees and pool coin's totally supply
		for _, reserveCoin := range reserveCoins {
			if err := types.CheckOverflow(reserveCoin.Amount, msg.Msg.PoolCoin.Amount); err != nil {
				return err
			}
			if err := types.CheckOverflow(sdk.NewDecFromInt(reserveCoin.Amount.Mul(msg.Msg.PoolCoin.Amount)).TruncateInt(), poolCoinTotalSupply); err != nil {
				return err
			}
			// WithdrawAmount = ReserveAmount * PoolCoinAmount * WithdrawFeeProportion / TotalSupply
			withdrawAmtWithFee := sdk.NewDecFromInt(reserveCoin.Amount.Mul(msg.Msg.PoolCoin.Amount)).TruncateInt().Quo(poolCoinTotalSupply)
			withdrawAmt := sdk.NewDecFromInt(reserveCoin.Amount.Mul(msg.Msg.PoolCoin.Amount)).MulTruncate(withdrawProportion).TruncateInt().Quo(poolCoinTotalSupply)
			withdrawCoins = append(withdrawCoins, sdk.NewCoin(reserveCoin.Denom, withdrawAmt))
			withdrawFeeCoins = append(withdrawFeeCoins, sdk.NewCoin(reserveCoin.Denom, withdrawAmtWithFee.Sub(withdrawAmt)))
		}
	}

	if withdrawCoins.IsValid() {
		inputs = append(inputs, banktypes.NewInput(reserveAcc, withdrawCoins))
		outputs = append(outputs, banktypes.NewOutput(withdrawer, withdrawCoins))
	} else {
		return types.ErrBadPoolCoinAmount
	}

	// send withdrawing coins to the withdrawer
	if err := k.bankKeeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
		return err
	}

	// burn the escrowed pool coins
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, poolCoins); err != nil {
		return err
	}

	msg.Succeeded = true
	msg.ToBeDeleted = true
	k.SetPoolBatchWithdrawMsgState(ctx, msg.Msg.PoolId, msg)

	if BatchLogicInvariantCheckFlag {
		afterPoolCoinTotalSupply := k.GetPoolCoinTotalSupply(ctx, pool)
		afterReserveCoins := k.GetReserveCoins(ctx, pool)
		afterReserveCoinA := sdk.ZeroInt()
		afterReserveCoinB := sdk.ZeroInt()
		if !afterReserveCoins.IsZero() {
			afterReserveCoinA = afterReserveCoins[0].Amount
			afterReserveCoinB = afterReserveCoins[1].Amount
		}
		burnedPoolCoin := poolCoins[0].Amount
		withdrawCoinA := withdrawCoins[0].Amount
		withdrawCoinB := withdrawCoins[1].Amount
		reserveCoinA := reserveCoins[0].Amount
		reserveCoinB := reserveCoins[1].Amount
		lastPoolCoinTotalSupply := poolCoinTotalSupply
		afterPoolTotalSupply := afterPoolCoinTotalSupply

		BurningPoolCoinsInvariant(burnedPoolCoin, withdrawCoinA, withdrawCoinB, reserveCoinA, reserveCoinB, lastPoolCoinTotalSupply, withdrawFeeCoins)
		WithdrawReserveCoinsInvariant(withdrawCoinA, withdrawCoinB, reserveCoinA, reserveCoinB,
			afterReserveCoinA, afterReserveCoinB, afterPoolTotalSupply, lastPoolCoinTotalSupply, burnedPoolCoin)
		WithdrawAmountInvariant(withdrawCoinA, withdrawCoinB, reserveCoinA, reserveCoinB, burnedPoolCoin, lastPoolCoinTotalSupply, params.WithdrawFeeRate)
		ImmutablePoolPriceAfterWithdrawInvariant(reserveCoinA, reserveCoinB, withdrawCoinA, withdrawCoinB, afterReserveCoinA, afterReserveCoinB)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawFromPool,
			sdk.NewAttribute(types.AttributeValuePoolId, strconv.FormatUint(pool.Id, 10)),
			sdk.NewAttribute(types.AttributeValueBatchIndex, strconv.FormatUint(batch.Index, 10)),
			sdk.NewAttribute(types.AttributeValueMsgIndex, strconv.FormatUint(msg.MsgIndex, 10)),
			sdk.NewAttribute(types.AttributeValueWithdrawer, withdrawer.String()),
			sdk.NewAttribute(types.AttributeValuePoolDenom, msg.Msg.PoolCoin.Denom),
			sdk.NewAttribute(types.AttributeValuePoolCoinAmount, msg.Msg.PoolCoin.Amount.String()),
			sdk.NewAttribute(types.AttributeValueWithdrawCoins, withdrawCoins.String()),
			sdk.NewAttribute(types.AttributeValueWithdrawFeeCoins, withdrawFeeCoins.String()),
			sdk.NewAttribute(types.AttributeValueSuccess, types.Success),
		),
	)

	reserveCoins = k.GetReserveCoins(ctx, pool)

	var lastReserveRatio sdk.Dec
	if reserveCoins.IsZero() {
		lastReserveRatio = sdk.ZeroDec()
	} else {
		lastReserveRatio = sdk.NewDecFromInt(reserveCoins[0].Amount).Quo(sdk.NewDecFromInt(reserveCoins[1].Amount))
	}

	logger := k.Logger(ctx)
	logger.Debug(
		"withdraw pool coin from the pool",
		"msg", msg,
		"pool", pool,
		"inputs", inputs,
		"outputs", outputs,
		"reserveCoins", reserveCoins,
		"lastReserveRatio", lastReserveRatio,
	)

	return nil
}

// GetPoolCoinTotalSupply returns total supply of pool coin of the pool in form of sdk.Int
//
//nolint:staticcheck
func (k Keeper) GetPoolCoinTotalSupply(ctx sdk.Context, pool types.Pool) sdk.Int {
	return k.bankKeeper.GetSupply(ctx, pool.PoolCoinDenom).Amount
}

// IsDepletedPool returns true if the pool is depleted.
func (k Keeper) IsDepletedPool(ctx sdk.Context, pool types.Pool) bool {
	reserveCoins := k.GetReserveCoins(ctx, pool)
	return !k.GetPoolCoinTotalSupply(ctx, pool).IsPositive() ||
		reserveCoins.AmountOf(pool.ReserveCoinDenoms[0]).IsZero() ||
		reserveCoins.AmountOf(pool.ReserveCoinDenoms[1]).IsZero()
}

// GetPoolCoinTotal returns total supply of pool coin of the pool in form of sdk.Coin
func (k Keeper) GetPoolCoinTotal(ctx sdk.Context, pool types.Pool) sdk.Coin {
	return sdk.NewCoin(pool.PoolCoinDenom, k.GetPoolCoinTotalSupply(ctx, pool))
}

// GetReserveCoins returns reserve coins from the liquidity pool
func (k Keeper) GetReserveCoins(ctx sdk.Context, pool types.Pool) (reserveCoins sdk.Coins) {
	reserveAcc := pool.GetReserveAccount()
	reserveCoins = sdk.NewCoins()
	for _, denom := range pool.ReserveCoinDenoms {
		reserveCoins = append(reserveCoins, k.bankKeeper.GetBalance(ctx, reserveAcc, denom))
	}
	return
}

// GetPoolMetaData returns metadata of the pool
func (k Keeper) GetPoolMetaData(ctx sdk.Context, pool types.Pool) types.PoolMetadata {
	return types.PoolMetadata{
		PoolId:              pool.Id,
		PoolCoinTotalSupply: k.GetPoolCoinTotal(ctx, pool),
		ReserveCoins:        k.GetReserveCoins(ctx, pool),
	}
}

// GetPoolRecord returns the liquidity pool record with the given pool information
func (k Keeper) GetPoolRecord(ctx sdk.Context, pool types.Pool) (types.PoolRecord, bool) {
	batch, found := k.GetPoolBatch(ctx, pool.Id)
	if !found {
		return types.PoolRecord{}, false
	}
	return types.PoolRecord{
		Pool:              pool,
		PoolMetadata:      k.GetPoolMetaData(ctx, pool),
		PoolBatch:         batch,
		DepositMsgStates:  k.GetAllPoolBatchDepositMsgs(ctx, batch),
		WithdrawMsgStates: k.GetAllPoolBatchWithdrawMsgStates(ctx, batch),
		SwapMsgStates:     k.GetAllPoolBatchSwapMsgStates(ctx, batch),
	}, true
}

// SetPoolRecord stores liquidity pool states
func (k Keeper) SetPoolRecord(ctx sdk.Context, record types.PoolRecord) types.PoolRecord {
	k.SetPoolAtomic(ctx, record.Pool)
	if record.PoolBatch.BeginHeight > ctx.BlockHeight() {
		record.PoolBatch.BeginHeight = 0
	}
	k.SetPoolBatch(ctx, record.PoolBatch)
	k.SetPoolBatchDepositMsgStates(ctx, record.Pool.Id, record.DepositMsgStates)
	k.SetPoolBatchWithdrawMsgStates(ctx, record.Pool.Id, record.WithdrawMsgStates)
	k.SetPoolBatchSwapMsgStates(ctx, record.Pool.Id, record.SwapMsgStates)
	return record
}

// RefundDeposit refunds deposit amounts to the depositor
func (k Keeper) RefundDeposit(ctx sdk.Context, batchMsg types.DepositMsgState, batch types.PoolBatch) error {
	batchMsg, _ = k.GetPoolBatchDepositMsgState(ctx, batchMsg.Msg.PoolId, batchMsg.MsgIndex)
	if !batchMsg.Executed || batchMsg.Succeeded {
		return fmt.Errorf("cannot refund not executed or already succeeded msg")
	}
	pool, _ := k.GetPool(ctx, batchMsg.Msg.PoolId)
	if err := k.ReleaseEscrow(ctx, batchMsg.Msg.GetDepositor(), batchMsg.Msg.DepositCoins); err != nil {
		return err
	}
	// not delete now, set ToBeDeleted true for delete on next block beginblock
	batchMsg.ToBeDeleted = true
	k.SetPoolBatchDepositMsgState(ctx, batchMsg.Msg.PoolId, batchMsg)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDepositToPool,
			sdk.NewAttribute(types.AttributeValuePoolId, strconv.FormatUint(pool.Id, 10)),
			sdk.NewAttribute(types.AttributeValueBatchIndex, strconv.FormatUint(batch.Index, 10)),
			sdk.NewAttribute(types.AttributeValueMsgIndex, strconv.FormatUint(batchMsg.MsgIndex, 10)),
			sdk.NewAttribute(types.AttributeValueDepositor, batchMsg.Msg.GetDepositor().String()),
			sdk.NewAttribute(types.AttributeValueAcceptedCoins, sdk.NewCoins().String()),
			sdk.NewAttribute(types.AttributeValueRefundedCoins, batchMsg.Msg.DepositCoins.String()),
			sdk.NewAttribute(types.AttributeValueSuccess, types.Failure),
		))
	return nil
}

// RefundWithdrawal refunds pool coin of the liquidity pool to the withdrawer
func (k Keeper) RefundWithdrawal(ctx sdk.Context, batchMsg types.WithdrawMsgState, batch types.PoolBatch) error {
	batchMsg, _ = k.GetPoolBatchWithdrawMsgState(ctx, batchMsg.Msg.PoolId, batchMsg.MsgIndex)
	if !batchMsg.Executed || batchMsg.Succeeded {
		return fmt.Errorf("cannot refund not executed or already succeeded msg")
	}
	pool, _ := k.GetPool(ctx, batchMsg.Msg.PoolId)
	if err := k.ReleaseEscrow(ctx, batchMsg.Msg.GetWithdrawer(), sdk.NewCoins(batchMsg.Msg.PoolCoin)); err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawFromPool,
			sdk.NewAttribute(types.AttributeValuePoolId, strconv.FormatUint(pool.Id, 10)),
			sdk.NewAttribute(types.AttributeValueBatchIndex, strconv.FormatUint(batch.Index, 10)),
			sdk.NewAttribute(types.AttributeValueMsgIndex, strconv.FormatUint(batchMsg.MsgIndex, 10)),
			sdk.NewAttribute(types.AttributeValueWithdrawer, batchMsg.Msg.GetWithdrawer().String()),
			sdk.NewAttribute(types.AttributeValuePoolDenom, batchMsg.Msg.PoolCoin.Denom),
			sdk.NewAttribute(types.AttributeValuePoolCoinAmount, batchMsg.Msg.PoolCoin.Amount.String()),
			sdk.NewAttribute(types.AttributeValueSuccess, types.Failure),
		))

	// not delete now, set ToBeDeleted true for delete on next block beginblock
	batchMsg.ToBeDeleted = true
	k.SetPoolBatchWithdrawMsgState(ctx, batchMsg.Msg.PoolId, batchMsg)
	return nil
}

// ValidateMsgDepositWithinBatch validates MsgDepositWithinBatch
func (k Keeper) ValidateMsgDepositWithinBatch(ctx sdk.Context, msg types.MsgDepositWithinBatch) error {
	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return types.ErrPoolNotExists
	}

	if msg.DepositCoins.Len() != len(pool.ReserveCoinDenoms) {
		return types.ErrNumOfReserveCoin
	}

	params := k.GetParams(ctx)
	reserveCoins := k.GetReserveCoins(ctx, pool)
	if err := types.ValidateReserveCoinLimit(params.MaxReserveCoinAmount, reserveCoins.Add(msg.DepositCoins...)); err != nil {
		return err
	}

	denomA, denomB := types.AlphabeticalDenomPair(msg.DepositCoins[0].Denom, msg.DepositCoins[1].Denom)
	if denomA != pool.ReserveCoinDenoms[0] || denomB != pool.ReserveCoinDenoms[1] {
		return types.ErrNotMatchedReserveCoin
	}
	return nil
}

// ValidateMsgWithdrawWithinBatch validates MsgWithdrawWithinBatch
func (k Keeper) ValidateMsgWithdrawWithinBatch(ctx sdk.Context, msg types.MsgWithdrawWithinBatch) error {
	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return types.ErrPoolNotExists
	}

	if msg.PoolCoin.Denom != pool.PoolCoinDenom {
		return types.ErrBadPoolCoinDenom
	}

	poolCoinTotalSupply := k.GetPoolCoinTotalSupply(ctx, pool)
	if k.IsDepletedPool(ctx, pool) {
		return types.ErrDepletedPool
	}

	if msg.PoolCoin.Amount.GT(poolCoinTotalSupply) {
		return types.ErrBadPoolCoinAmount
	}
	return nil
}

// ValidatePool validates logic for liquidity pool after set or before export
func (k Keeper) ValidatePool(ctx sdk.Context, pool *types.Pool) error {
	params := k.GetParams(ctx)
	var poolType types.PoolType

	// check poolType exist, get poolType from param
	if len(params.PoolTypes) >= int(pool.TypeId) {
		poolType = params.PoolTypes[pool.TypeId-1]
		if poolType.Id != pool.TypeId {
			return types.ErrPoolTypeNotExists
		}
	} else {
		return types.ErrPoolTypeNotExists
	}

	if poolType.MaxReserveCoinNum > types.MaxReserveCoinNum || types.MinReserveCoinNum > poolType.MinReserveCoinNum {
		return types.ErrNumOfReserveCoin
	}

	reserveCoins := k.GetReserveCoins(ctx, *pool)
	if uint32(reserveCoins.Len()) > poolType.MaxReserveCoinNum || poolType.MinReserveCoinNum > uint32(reserveCoins.Len()) {
		return types.ErrNumOfReserveCoin
	}

	if len(pool.ReserveCoinDenoms) != reserveCoins.Len() {
		return types.ErrNumOfReserveCoin
	}
	for i, denom := range pool.ReserveCoinDenoms {
		if denom != reserveCoins[i].Denom {
			return types.ErrInvalidDenom
		}
	}

	denomA, denomB := types.AlphabeticalDenomPair(pool.ReserveCoinDenoms[0], pool.ReserveCoinDenoms[1])
	if denomA != pool.ReserveCoinDenoms[0] || denomB != pool.ReserveCoinDenoms[1] {
		return types.ErrBadOrderingReserveCoin
	}

	poolName := types.PoolName(pool.ReserveCoinDenoms, pool.TypeId)
	poolCoin := k.GetPoolCoinTotal(ctx, *pool)
	if poolCoin.Denom != types.GetPoolCoinDenom(poolName) {
		return types.ErrBadPoolCoinDenom
	}

	_, found := k.GetPoolBatch(ctx, pool.Id)
	if !found {
		return types.ErrPoolBatchNotExists
	}

	return nil
}

// ValidatePoolMetadata validates logic for liquidity pool metadata
func (k Keeper) ValidatePoolMetadata(ctx sdk.Context, pool *types.Pool, metaData *types.PoolMetadata) error {
	if err := metaData.ReserveCoins.Validate(); err != nil {
		return err
	}
	if !metaData.ReserveCoins.IsEqual(k.GetReserveCoins(ctx, *pool)) {
		return types.ErrNumOfReserveCoin
	}
	if !metaData.PoolCoinTotalSupply.IsEqual(sdk.NewCoin(pool.PoolCoinDenom, k.GetPoolCoinTotalSupply(ctx, *pool))) {
		return types.ErrBadPoolCoinAmount
	}
	return nil
}

// ValidatePoolRecord validates liquidity pool record after init or after export
func (k Keeper) ValidatePoolRecord(ctx sdk.Context, record types.PoolRecord) error {
	if err := k.ValidatePool(ctx, &record.Pool); err != nil {
		return err
	}

	if err := k.ValidatePoolMetadata(ctx, &record.Pool, &record.PoolMetadata); err != nil {
		return err
	}

	if len(record.DepositMsgStates) != 0 && record.PoolBatch.DepositMsgIndex != record.DepositMsgStates[len(record.DepositMsgStates)-1].MsgIndex+1 {
		return types.ErrBadBatchMsgIndex
	}
	if len(record.WithdrawMsgStates) != 0 && record.PoolBatch.WithdrawMsgIndex != record.WithdrawMsgStates[len(record.WithdrawMsgStates)-1].MsgIndex+1 {
		return types.ErrBadBatchMsgIndex
	}
	if len(record.SwapMsgStates) != 0 && record.PoolBatch.SwapMsgIndex != record.SwapMsgStates[len(record.SwapMsgStates)-1].MsgIndex+1 {
		return types.ErrBadBatchMsgIndex
	}

	return nil
}

// IsPoolCoinDenom returns true if the denom is a valid pool coin denom.
func (k Keeper) IsPoolCoinDenom(ctx sdk.Context, denom string) bool {
	reserveAcc, err := types.GetReserveAcc(denom, false)
	if err != nil {
		return false
	}
	_, found := k.GetPoolByReserveAccIndex(ctx, reserveAcc)
	return found
}
