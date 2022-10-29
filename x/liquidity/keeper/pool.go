package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"ollo/x/liquidity/amm"
	"ollo/x/liquidity/types"
)

// getNextPoolIdWithUpdate increments pool id by one and set it.
func (k Keeper) getNextPoolIdWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetLastPoolId(ctx) + 1
	k.SetLastPairId(ctx, id)
	return id
}

// getNextRequestDepositIdWithUpdate increments the pool's Prev deposit request
// id and returns it.
func (k Keeper) getNextRequestDepositIdWithUpdate(ctx sdk.Context, pool types.Pool) uint64 {
	id := pool.PrevDepositReqId + 1
	pool.PrevDepositReqId = id
	k.SetPool(ctx, pool)
	return id
}

// getNextRequestWithdrawIdWithUpdate increments the pool's Prev withdraw
// request id and returns it.
func (k Keeper) getNextRequestWithdrawIdWithUpdate(ctx sdk.Context, pool types.Pool) uint64 {
	id := pool.PrevWithdrawReqId + 1
	pool.PrevWithdrawReqId = id
	k.SetPool(ctx, pool)
	return id
}

// GetPoolBalances returns the balances of the pool.
func (k Keeper) GetPoolBalances(ctx sdk.Context, pool types.Pool) (rx sdk.Coin, ry sdk.Coin) {
	reserveAddr := pool.GetReserveAddress()
	pair, _ := k.GetPair(ctx, pool.PairId)
	spendable := k.bankKeeper.SpendableCoins(ctx, reserveAddr)
	rx = sdk.NewCoin(pair.QuoteDenom, spendable.AmountOf(pair.QuoteDenom))
	ry = sdk.NewCoin(pair.BaseDenom, spendable.AmountOf(pair.BaseDenom))
	return
}

// getPoolBalances returns the balances of the pool.
// It is used internally when caller already has types.Pair instance.
func (k Keeper) getPoolBalances(ctx sdk.Context, pool types.Pool, pair types.Pair) (rx sdk.Coin, ry sdk.Coin) {
	reserveAddr := pool.GetReserveAddress()
	spendable := k.bankKeeper.SpendableCoins(ctx, reserveAddr)
	rx = sdk.NewCoin(pair.QuoteDenom, spendable.AmountOf(pair.QuoteDenom))
	ry = sdk.NewCoin(pair.BaseDenom, spendable.AmountOf(pair.BaseDenom))
	return
}

// GetPoolCoinSupply returns total pool coin supply of the pool.
func (k Keeper) GetPoolCoinSupply(ctx sdk.Context, pool types.Pool) sdk.Int {
	return k.bankKeeper.GetSupply(ctx, pool.Reserve.Denom).Amount
}

// MarkPoolAsDisabled marks a pool as Disabled.
func (k Keeper) MarkPoolAsDisabled(ctx sdk.Context, pool types.Pool) {
	pool.Disabled = true
	k.SetPool(ctx, pool)
}

// ValidateMsgCreatePool validates types.MsgCreatePool.
func (k Keeper) ValidateMsgCreatePool(ctx sdk.Context, msg *types.MsgCreatePool) error {
	pair, found := k.GetPair(ctx, msg.PairId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pair %d not found", msg.PairId)
	}

	minInitDepositAmt := k.GetMinInitialDepositAmount(ctx)
	for _, coin := range msg.DepositCoins {
		if coin.Denom != pair.BaseDenom && coin.Denom != pair.QuoteDenom {
			return sdkerrors.Wrapf(types.ErrInvalidCoinDenom, "coin denom %s is not in the pair", coin.Denom)
		}
		minDepositCoin := sdk.NewCoin(coin.Denom, minInitDepositAmt)
		if coin.IsLT(minDepositCoin) {
			return sdkerrors.Wrapf(
				types.ErrInsufficientDepositAmount, "%s is smaller than %s", coin, minDepositCoin)
		}
	}

	// Check if there is a basic pool in the pair.
	// Creating multiple basic pools within the same pair is disallowed.
	duplicate := false
	numActivePools := 0
	_ = k.IteratePoolsByPair(ctx, pair.Id, func(pool types.Pool) (stop bool, err error) {
		if pool.TypeId == 1 && !pool.Disabled {
			duplicate = true
			return true, nil
		}
		if !pool.Disabled {
			numActivePools++
		}
		return false, nil
	})
	if duplicate {
		return types.ErrPoolAlreadyExists
	}
	if numActivePools >= types.MaxNumActivePoolsPerPair {
		return types.ErrTooManyPools
	}

	return nil
}

// CreatePool handles types.MsgCreatePool and creates a basic pool.
func (k Keeper) CreatePool(ctx sdk.Context, msg *types.MsgCreatePool) (types.Pool, error) {
	if err := k.ValidateMsgCreatePool(ctx, msg); err != nil {
		return types.Pool{}, err
	}

	pair, _ := k.GetPair(ctx, msg.PairId)

	x, y := msg.DepositCoins.AmountOf(pair.QuoteDenom), msg.DepositCoins.AmountOf(pair.BaseDenom)
	ammPool, err := amm.CreateBasicPool(x, y)
	if err != nil {
		return types.Pool{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Create and save the new pool object.
	poolId := k.getNextPoolIdWithUpdate(ctx)
	pool := types.NewBasicPool(poolId, pair.Id, msg.GetCreator())
	k.SetPool(ctx, pool)
	k.SetPoolByReserveIndex(ctx, pool)
	k.SetPoolsByPairIndex(ctx, pool)

	// Send deposit coins to the pool's reserve account.
	creator := msg.GetCreator()
	if err := k.bankKeeper.SendCoins(ctx, creator, pool.GetReserveAddress(), msg.DepositCoins); err != nil {
		return types.Pool{}, err
	}

	// Send the pool creation fee to the fee collector.
	if err := k.bankKeeper.SendCoins(ctx, creator, k.GetFeeCollector(ctx), k.GetPoolCreationFee(ctx)); err != nil {
		return types.Pool{}, sdkerrors.Wrap(err, "insufficient pool creation fee")
	}

	// Mint and send pool coin to the creator.
	// Minting pool coin amount is calculated based on two coins' amount.
	// Minimum minting amount is params.MinInitialPoolCoinSupply.
	ps := sdk.MaxInt(ammPool.PoolCoinSupply(), k.GetMinInitialPoolCoinSupply(ctx))
	poolCoin := sdk.NewCoin(pool.Reserve.Denom, ps)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(poolCoin)); err != nil {
		return types.Pool{}, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(poolCoin)); err != nil {
		return types.Pool{}, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyPairId, strconv.FormatUint(msg.PairId, 10)),
			sdk.NewAttribute(types.AttributeKeyDepositCoins, msg.DepositCoins.String()),
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(pool.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyReserveAddress, pool.Reserve.Addr),
			sdk.NewAttribute(types.AttributeKeyMintedPoolCoin, poolCoin.String()),
		),
	})

	return pool, nil
}

// ValidateMsgCreateIntelligentPool validates types.MsgCreateIntelligentPool.
func (k Keeper) ValidateMsgCreatePoolCapped(ctx sdk.Context, msg *types.MsgCreatePoolCapped) error {
	tickPrec := k.GetTickPrecision(ctx)
	if !amm.PriceToDownTick(msg.MinPrice, int(tickPrec)).Equal(msg.MinPrice) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "min price is not on ticks")
	}
	if !amm.PriceToDownTick(msg.MaxPrice, int(tickPrec)).Equal(msg.MaxPrice) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "max price is not on ticks")
	}
	if !amm.PriceToDownTick(msg.InitialPrice, int(tickPrec)).Equal(msg.InitialPrice) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "initial price is not on ticks")
	}

	lowestTick := amm.LowestTick(int(tickPrec))
	if msg.MinPrice.LT(lowestTick) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "min price must not be less than %s", lowestTick)
	}

	pair, found := k.GetPair(ctx, msg.PairId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pair %d not found", msg.PairId)
	}

	for _, coin := range msg.DepositCoins {
		if coin.Denom != pair.BaseDenom && coin.Denom != pair.QuoteDenom {
			return sdkerrors.Wrapf(types.ErrInvalidCoinDenom, "coin denom %s is not in the pair", coin.Denom)
		}
	}

	numActivePools := 0
	_ = k.IteratePoolsByPair(ctx, pair.Id, func(pool types.Pool) (stop bool, err error) {
		if !pool.Disabled {
			numActivePools++
		}
		return false, nil
	})
	if numActivePools >= types.MaxNumActivePoolsPerPair {
		return types.ErrTooManyPools
	}

	return nil
}

// CreateIntelligentPool handles types.MsgCreateIntelligentPool and creates a Intelligent pool.
func (k Keeper) CreateCappedPool(ctx sdk.Context, msg *types.MsgCreatePoolCapped) (types.Pool, error) {
	if err := k.ValidateMsgCreatePoolCapped(ctx, msg); err != nil {
		return types.Pool{}, err
	}

	pair, _ := k.GetPair(ctx, msg.PairId)

	x, y := msg.DepositCoins.AmountOf(pair.QuoteDenom), msg.DepositCoins.AmountOf(pair.BaseDenom)
	ammPool, err := amm.CreatePoolCapped(x, y, msg.MinPrice, msg.MaxPrice, msg.InitialPrice)
	if err != nil {
		return types.Pool{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	ax, ay := ammPool.Balances()

	minInitDepositAmt := k.GetMinInitialDepositAmount(ctx)
	if ax.LT(minInitDepositAmt) && ay.LT(minInitDepositAmt) {
		return types.Pool{}, types.ErrInsufficientDepositAmount
	}

	// Create and save the new pool object.
	poolId := k.getNextPoolIdWithUpdate(ctx)
	pool := types.NewIntelligentPool(poolId, pair.Id, msg.GetCreatorAddress(), msg.MinPrice, msg.MaxPrice)
	k.SetPool(ctx, pool)
	k.SetPoolByReserveIndex(ctx, pool)
	k.SetPoolsByPairIndex(ctx, pool)

	// Send deposit coins to the pool's reserve account.
	creator := msg.GetCreatorAddress()
	depositCoins := sdk.NewCoins(
		sdk.NewCoin(pair.QuoteDenom, ax), sdk.NewCoin(pair.BaseDenom, ay))
	if err := k.bankKeeper.SendCoins(ctx, creator, pool.GetReserveAddress(), depositCoins); err != nil {
		return types.Pool{}, err
	}

	// Send the pool creation fee to the fee collector.
	feeCollector := k.GetFeeCollector(ctx)
	poolCreationFee := k.GetPoolCreationFee(ctx)
	if err := k.bankKeeper.SendCoins(ctx, creator, feeCollector, poolCreationFee); err != nil {
		return types.Pool{}, sdkerrors.Wrap(err, "insufficient pool creation fee")
	}

	// Mint and send pool coin to the creator.
	// Minimum minting amount is params.MinInitialPoolCoinSupply.
	ps := sdk.MaxInt(ammPool.PoolCoinSupply(), k.GetMinInitialPoolCoinSupply(ctx))
	poolCoin := sdk.NewCoin(pool.Reserve.Denom, ps)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(poolCoin)); err != nil {
		return types.Pool{}, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(poolCoin)); err != nil {
		return types.Pool{}, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateIntelligentPool,
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyPairId, strconv.FormatUint(msg.PairId, 10)),
			sdk.NewAttribute(types.AttributeKeyDepositCoins, msg.DepositCoins.String()),
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(pool.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyReserveAddress, pool.Reserve.Addr),
			sdk.NewAttribute(types.AttributeKeyMintedPoolCoin, poolCoin.String()),
		),
	})

	return pool, nil
}

// ValidateMsgDeposit validates types.MsgDeposit.
func (k Keeper) ValidateMsgDeposit(ctx sdk.Context, msg *types.MsgDeposit) error {
	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", msg.PoolId)
	}
	if pool.Disabled {
		return types.ErrInactivePool
	}
	if pool.TypeId == 1 && len(msg.DepositCoins) != 2 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "wrong number of deposit coins: %d", len(msg.DepositCoins))
	}

	pair, _ := k.GetPair(ctx, pool.PairId)

	for _, coin := range msg.DepositCoins {
		if coin.Denom != pair.BaseDenom && coin.Denom != pair.QuoteDenom {
			return sdkerrors.Wrapf(types.ErrInvalidCoinDenom, "coin denom %s is not in the pair", coin.Denom)
		}
	}

	rx, ry := k.getPoolBalances(ctx, pool, pair)
	if rx.Amount.Add(msg.DepositCoins.AmountOf(rx.Denom)).GT(amm.MaxCoinAmount) {
		return types.ErrTooLargePool
	}
	if ry.Amount.Add(msg.DepositCoins.AmountOf(ry.Denom)).GT(amm.MaxCoinAmount) {
		return types.ErrTooLargePool
	}

	return nil
}

// Deposit handles types.MsgDeposit and stores the request.
func (k Keeper) Deposit(ctx sdk.Context, msg *types.MsgDeposit) (types.RequestDeposit, error) {
	if err := k.ValidateMsgDeposit(ctx, msg); err != nil {
		return types.RequestDeposit{}, err
	}

	if err := k.bankKeeper.SendCoins(ctx, msg.GetDepositorAddress(), types.GlobalEscrowAddress, msg.DepositCoins); err != nil {
		return types.RequestDeposit{}, err
	}

	pool, _ := k.GetPool(ctx, msg.PoolId)
	requestId := k.getNextRequestDepositIdWithUpdate(ctx, pool)
	req := types.NewRequestDeposit(msg, pool, requestId, ctx.BlockHeight())
	k.SetRequestDeposit(ctx, req)
	k.SetRequestDepositIndex(ctx, req)

	ctx.GasMeter().ConsumeGas(k.GetDepositExtraGas(ctx), "DepositExtraGas")

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDeposit,
			sdk.NewAttribute(types.AttributeKeyDepositor, msg.Depositor),
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(pool.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyDepositCoins, msg.DepositCoins.String()),
			sdk.NewAttribute(types.AttributeKeyRequestId, strconv.FormatUint(req.Id, 10)),
		),
	})

	return req, nil
}

// ValidateMsgWithdraw validates types.MsgWithdraw.
func (k Keeper) ValidateMsgWithdraw(ctx sdk.Context, msg *types.MsgWithdraw) error {
	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", msg.PoolId)
	}
	if pool.Disabled {
		return types.ErrInactivePool
	}

	if msg.PoolCoin.Denom != pool.Reserve.Denom {
		return types.ErrWrongPoolCoinDenom
	}

	return nil
}

// Withdraw handles types.MsgWithdraw and stores the request.
func (k Keeper) Withdraw(ctx sdk.Context, msg *types.MsgWithdraw) (types.RequestWithdraw, error) {
	if err := k.ValidateMsgWithdraw(ctx, msg); err != nil {
		return types.RequestWithdraw{}, err
	}

	pool, _ := k.GetPool(ctx, msg.PoolId)
	if err := k.bankKeeper.SendCoins(ctx, msg.GetWithdrawerAddress(), types.GlobalEscrowAddress, sdk.NewCoins(msg.PoolCoin)); err != nil {
		return types.RequestWithdraw{}, err
	}

	requestId := k.getNextRequestWithdrawIdWithUpdate(ctx, pool)
	req := types.NewRequestWithdraw(msg, requestId, ctx.BlockHeight())
	k.SetRequestWithdraw(ctx, req)
	k.SetRequestWithdrawIndex(ctx, req)

	ctx.GasMeter().ConsumeGas(k.GetWithdrawExtraGas(ctx), "WithdrawExtraGas")

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdraw,
			sdk.NewAttribute(types.AttributeKeyWithdrawer, msg.Withdrawer),
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(pool.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyPoolCoin, msg.PoolCoin.String()),
			sdk.NewAttribute(types.AttributeKeyRequestId, strconv.FormatUint(req.Id, 10)),
		),
	})

	return req, nil
}

// ExecuteRequestDeposit executes a deposit request.
func (k Keeper) ExecuteRequestDeposit(ctx sdk.Context, req types.RequestDeposit) error {
	pool, _ := k.GetPool(ctx, req.PoolId)
	if pool.Disabled {
		if err := k.FinishRequestDeposit(ctx, req, types.RequestStatusFail); err != nil {
			return fmt.Errorf("refund deposit request: %w", err)
		}
		return nil
	}

	pair, _ := k.GetPair(ctx, pool.PairId)
	rx, ry := k.getPoolBalances(ctx, pool, pair)
	ps := k.GetPoolCoinSupply(ctx, pool)
	ammPool := pool.AMMPool(rx.Amount, ry.Amount, ps)
	if ammPool.IsDepleted() {
		k.MarkPoolAsDisabled(ctx, pool)
		if err := k.FinishRequestDeposit(ctx, req, types.RequestStatusFail); err != nil {
			return err
		}
		return nil
	}

	ax, ay, pc := amm.Deposit(rx.Amount, ry.Amount, ps, req.DepositAmt.AmountOf(pair.QuoteDenom), req.DepositAmt.AmountOf(pair.BaseDenom))

	if pc.IsZero() {
		if err := k.FinishRequestDeposit(ctx, req, types.RequestStatusFail); err != nil {
			return err
		}
		return nil
	}

	mintedPoolCoin := sdk.NewCoin(pool.Reserve.Denom, pc)
	mintingCoins := sdk.NewCoins(mintedPoolCoin)

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintingCoins); err != nil {
		return err
	}

	acceptedCoins := sdk.NewCoins(sdk.NewCoin(pair.QuoteDenom, ax), sdk.NewCoin(pair.BaseDenom, ay))
	bulkOp := types.NewBulkSendCoinsOperation()
	bulkOp.QueueSendCoins(types.GlobalEscrowAddress, pool.GetReserveAddress(), acceptedCoins)
	bulkOp.QueueSendCoins(k.accountKeeper.GetModuleAddress(types.ModuleName), req.GetDepositor(), mintingCoins)
	if err := bulkOp.Run(ctx, k.bankKeeper); err != nil {
		return err
	}

	req.AcceptedAmt = acceptedCoins
	req.PoolCoin = mintedPoolCoin
	if err := k.FinishRequestDeposit(ctx, req, types.RequestStatusSuccess); err != nil {
		return err
	}
	return nil
}

// FinishRequestDeposit refunds unhandled deposit coins and set request status.
func (k Keeper) FinishRequestDeposit(ctx sdk.Context, req types.RequestDeposit, status types.RequestStatus) error {
	if req.Status != types.RequestStatusPending { // sanity check
		return nil
	}

	refundingCoins := req.DepositAmt.Sub(req.AcceptedAmt...)
	if !refundingCoins.IsZero() {
		if err := k.bankKeeper.SendCoins(ctx, types.GlobalEscrowAddress, req.GetDepositor(), refundingCoins); err != nil {
			return err
		}
	}
	req.SetStatus(status)
	k.SetRequestDeposit(ctx, req)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDepositResult,
			sdk.NewAttribute(types.AttributeKeyRequestId, strconv.FormatUint(req.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyDepositor, req.DepositorAddr),
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(req.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyDepositCoins, req.DepositAmt.String()),
			sdk.NewAttribute(types.AttributeKeyAcceptedCoins, req.AcceptedAmt.String()),
			sdk.NewAttribute(types.AttributeKeyRefundedCoins, refundingCoins.String()),
			sdk.NewAttribute(types.AttributeKeyMintedPoolCoin, req.PoolCoin.String()),
			sdk.NewAttribute(types.AttributeKeyStatus, req.Status.String()),
		),
	})

	return nil
}

// ExecuteRequestWithdraw executes a withdraw request.
func (k Keeper) ExecuteRequestWithdraw(ctx sdk.Context, req types.RequestWithdraw) error {
	pool, _ := k.GetPool(ctx, req.PoolId)
	if pool.Disabled {
		if err := k.FinishRequestWithdraw(ctx, req, types.RequestStatusFail); err != nil {
			return err
		}
		return nil
	}

	pair, _ := k.GetPair(ctx, pool.PairId)
	rx, ry := k.getPoolBalances(ctx, pool, pair)
	ps := k.GetPoolCoinSupply(ctx, pool)
	ammPool := pool.AMMPool(rx.Amount, ry.Amount, ps)
	if ammPool.IsDepleted() {
		k.MarkPoolAsDisabled(ctx, pool)
		if err := k.FinishRequestWithdraw(ctx, req, types.RequestStatusFail); err != nil {
			return err
		}
		return nil
	}

	x, y := amm.Withdraw(rx.Amount, ry.Amount, ps, req.PoolCoin.Amount, k.GetWithdrawFeeRate(ctx))
	if x.IsZero() && y.IsZero() {
		if err := k.FinishRequestWithdraw(ctx, req, types.RequestStatusFail); err != nil {
			return err
		}
		return nil
	}

	withdrawnCoins := sdk.NewCoins(sdk.NewCoin(pair.QuoteDenom, x), sdk.NewCoin(pair.BaseDenom, y))
	burningCoins := sdk.NewCoins(req.PoolCoin)

	bulkOp := types.NewBulkSendCoinsOperation()
	bulkOp.QueueSendCoins(types.GlobalEscrowAddress, k.accountKeeper.GetModuleAddress(types.ModuleName), burningCoins)
	bulkOp.QueueSendCoins(pool.GetReserveAddress(), req.GetWithdrawer(), withdrawnCoins)
	if err := bulkOp.Run(ctx, k.bankKeeper); err != nil {
		return err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burningCoins); err != nil {
		return err
	}

	// If the pool coin supply becomes 0, disable the pool.
	if req.PoolCoin.Amount.Equal(ps) {
		k.MarkPoolAsDisabled(ctx, pool)
	}

	req.WithdrawAmt = withdrawnCoins
	if err := k.FinishRequestWithdraw(ctx, req, types.RequestStatusSuccess); err != nil {
		return err
	}
	return nil
}

// FinishRequestWithdraw refunds unhandled pool coin and set request status.
func (k Keeper) FinishRequestWithdraw(ctx sdk.Context, req types.RequestWithdraw, status types.RequestStatus) error {
	if req.Status != types.RequestStatusPending { // sanity check
		return nil
	}

	var refundingCoins sdk.Coins
	if status == types.RequestStatusFail {
		refundingCoins = sdk.NewCoins(req.PoolCoin)
		if err := k.bankKeeper.SendCoins(ctx, types.GlobalEscrowAddress, req.GetWithdrawer(), refundingCoins); err != nil {
			return err
		}
	}
	req.SetStatus(status)
	k.SetRequestWithdraw(ctx, req)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawalResult,
			sdk.NewAttribute(types.AttributeKeyRequestId, strconv.FormatUint(req.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyWithdrawer, req.WithdrawAddr),
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(req.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyPoolCoin, req.PoolCoin.String()),
			sdk.NewAttribute(types.AttributeKeyRefundedCoins, refundingCoins.String()),
			sdk.NewAttribute(types.AttributeKeyWithdrawnCoins, req.WithdrawAmt.String()),
			sdk.NewAttribute(types.AttributeKeyStatus, req.Status.String()),
		),
	})

	return nil
}
