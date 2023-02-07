package keeper

// DONTCOVER

// Although written in msg_server_test.go, it is approached at the keeper level rather than at the msgServer level
// so is not included in the coverage.

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/liquidity/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the distribution MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Message server, handler for CreatePool msg
func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.GetCircuitBreakerEnabled(ctx) {
		return nil, types.ErrCircuitBreakerEnabled
	}

	pool, err := k.Keeper.CreatePool(ctx, msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeValuePoolId, strconv.FormatUint(pool.Id, 10)),
			sdk.NewAttribute(types.AttributeValuePoolTypeId, fmt.Sprintf("%d", msg.PoolTypeId)),
			sdk.NewAttribute(types.AttributeValuePoolName, pool.Name()),
			sdk.NewAttribute(types.AttributeValueReserveAccount, pool.ReserveAccountAddress),
			sdk.NewAttribute(types.AttributeValueDepositCoins, msg.DepositCoins.String()),
			sdk.NewAttribute(types.AttributeValuePoolDenom, pool.PoolCoinDenom),
		),
	})

	return &types.MsgCreatePoolResponse{}, nil
}

// func (m msgServer) CreateRangedPool(goCtx context.Context, msg *types.MsgCreateRangedPool) (*types.MsgCreateRangedPoolResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)
//
// 	if _, err := m.Keeper.CreateRangedPool(ctx, msg); err != nil {
// 		return nil, err
// 	}
//
// 	return &types.MsgCreateRangedPoolResponse{}, nil
// }

// Message server, handler for MsgDepositWithinBatch
func (k msgServer) DepositWithinBatch(goCtx context.Context, msg *types.MsgDepositWithinBatch) (*types.MsgDepositWithinBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.GetCircuitBreakerEnabled(ctx) {
		return nil, types.ErrCircuitBreakerEnabled
	}

	poolBatch, found := k.GetPoolBatch(ctx, msg.PoolId)
	if !found {
		return nil, types.ErrPoolBatchNotExists
	}

	batchMsg, err := k.Keeper.DepositWithinBatch(ctx, msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
		sdk.NewEvent(
			types.EventTypeDepositWithinBatch,
			sdk.NewAttribute(types.AttributeValuePoolId, strconv.FormatUint(batchMsg.Msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeValueBatchIndex, strconv.FormatUint(poolBatch.Index, 10)),
			sdk.NewAttribute(types.AttributeValueMsgIndex, strconv.FormatUint(batchMsg.MsgIndex, 10)),
			sdk.NewAttribute(types.AttributeValueDepositCoins, batchMsg.Msg.DepositCoins.String()),
		),
	})

	return &types.MsgDepositWithinBatchResponse{}, nil
}

// Message server, handler for MsgWithdrawWithinBatch
func (k msgServer) WithdrawWithinBatch(goCtx context.Context, msg *types.MsgWithdrawWithinBatch) (*types.MsgWithdrawWithinBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	poolBatch, found := k.GetPoolBatch(ctx, msg.PoolId)
	if !found {
		return nil, types.ErrPoolBatchNotExists
	}

	batchMsg, err := k.Keeper.WithdrawWithinBatch(ctx, msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
		sdk.NewEvent(
			types.EventTypeWithdrawWithinBatch,
			sdk.NewAttribute(types.AttributeValuePoolId, strconv.FormatUint(batchMsg.Msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeValueBatchIndex, strconv.FormatUint(poolBatch.Index, 10)),
			sdk.NewAttribute(types.AttributeValueMsgIndex, strconv.FormatUint(batchMsg.MsgIndex, 10)),
			sdk.NewAttribute(types.AttributeValuePoolDenom, batchMsg.Msg.PoolCoin.Denom),
			sdk.NewAttribute(types.AttributeValuePoolCoinAmount, batchMsg.Msg.PoolCoin.Amount.String()),
		),
	})

	return &types.MsgWithdrawWithinBatchResponse{}, nil
}

// Deprecated: Message server, handler for MsgSwapWithinBatch
func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwapWithinBatch) (*types.MsgSwapWithinBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.GetCircuitBreakerEnabled(ctx) {
		return nil, types.ErrCircuitBreakerEnabled
	}

	return &types.MsgSwapWithinBatchResponse{}, nil
}

// CreatePair defines a method to create a pair.
func (m msgServer) CreatePair(goCtx context.Context, msg *types.MsgCreatePair) (*types.MsgCreatePairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := m.Keeper.CreatePair(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgCreatePairResponse{}, nil
}
