package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/x/exchange/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// func (k msgServer) ExecuteAtomicSwap(goCtx context.Context, msg *types.MsgExecuteAtomicSwap) (*types.MsgExecuteAtomicSwapResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)
// 	err := k.Keeper.ExecuteAtomicSwap(ctx, msg.From, msg.To, msg.Amount, msg.Secret, msg.SecretHash, msg.ExpireHeight)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &types.MsgExecuteAtomicSwapResponse{}, nil
// }

// func (k msgServer) CreateAtomicSwap(goCtx context.Context, msg *types.MsgCreateAtomicSwap) (*types.MsgCreateAtomicSwapResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)
// 	err := k.Keeper.CreateAtomicSwap(ctx, msg.From, msg.To, msg.Amount, msg.SecretHash, msg.ExpireHeight)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &types.MsgCreateAtomicSwapResponse{}, nil
// }

// func (k msgServer) RefundAtomicSwap(goCtx context.Context, msg *types.MsgRefundAtomicSwap) (*types.MsgRefundAtomicSwapResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)
// 	err := k.Keeper.RefundAtomicSwap(ctx, msg.From, msg.To, msg.SecretHash, msg.ExpireHeight)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &types.MsgRefundAtomicSwapResponse{}, nil
// }

func (k msgServer) CreateMarket(goCtx context.Context, msg *types.MsgCreateMarket) (*types.MsgCreateMarketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	market, err := k.Keeper.CreateMarket(
		ctx, sdk.MustAccAddressFromBech32(msg.Sender), msg.BaseDenom, msg.QuoteDenom)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateMarketResponse{MarketId: market.Id}, nil
}

func (k msgServer) PlaceLimitOrder(goCtx context.Context, msg *types.MsgPlaceLimitOrder) (*types.MsgPlaceLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	orderId, _, res, err := k.Keeper.PlaceLimitOrder(
		ctx, msg.MarketId, sdk.MustAccAddressFromBech32(msg.Sender),
		msg.IsBuy, msg.Price, msg.Quantity, msg.Lifespan)
	if err != nil {
		return nil, err
	}
	return &types.MsgPlaceLimitOrderResponse{
		OrderId:          orderId,
		ExecutedQuantity: res.ExecutedQuantity,
		Paid:             res.Paid,
		Received:         res.Received,
	}, nil
}

func (k msgServer) PlaceBatchLimitOrder(goCtx context.Context, msg *types.MsgPlaceBatchLimitOrder) (*types.MsgPlaceBatchLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	order, err := k.Keeper.PlaceBatchLimitOrder(
		ctx, msg.MarketId, sdk.MustAccAddressFromBech32(msg.Sender),
		msg.IsBuy, msg.Price, msg.Quantity, msg.Lifespan)
	if err != nil {
		return nil, err
	}
	return &types.MsgPlaceBatchLimitOrderResponse{
		OrderId: order.Id,
	}, nil
}

func (k msgServer) PlaceMMLimitOrder(goCtx context.Context, msg *types.MsgPlaceMMLimitOrder) (*types.MsgPlaceMMLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	orderId, _, res, err := k.Keeper.PlaceMMLimitOrder(
		ctx, msg.MarketId, sdk.MustAccAddressFromBech32(msg.Sender),
		msg.IsBuy, msg.Price, msg.Quantity, msg.Lifespan)
	if err != nil {
		return nil, err
	}
	return &types.MsgPlaceMMLimitOrderResponse{
		OrderId:          orderId,
		ExecutedQuantity: res.ExecutedQuantity,
		Paid:             res.Paid,
		Received:         res.Received,
	}, nil
}

func (k msgServer) PlaceMMBatchLimitOrder(goCtx context.Context, msg *types.MsgPlaceMMBatchLimitOrder) (*types.MsgPlaceMMBatchLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	order, err := k.Keeper.PlaceMMBatchLimitOrder(
		ctx, msg.MarketId, sdk.MustAccAddressFromBech32(msg.Sender),
		msg.IsBuy, msg.Price, msg.Quantity, msg.Lifespan)
	if err != nil {
		return nil, err
	}
	return &types.MsgPlaceMMBatchLimitOrderResponse{
		OrderId: order.Id,
	}, nil
}

func (k msgServer) PlaceMarketOrder(goCtx context.Context, msg *types.MsgPlaceMarketOrder) (*types.MsgPlaceMarketOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	orderId, res, err := k.Keeper.PlaceMarketOrder(
		ctx, msg.MarketId, sdk.MustAccAddressFromBech32(msg.Sender),
		msg.IsBuy, msg.Quantity)
	if err != nil {
		return nil, err
	}
	return &types.MsgPlaceMarketOrderResponse{
		OrderId:          orderId,
		ExecutedQuantity: res.ExecutedQuantity,
		Paid:             res.Paid,
		Received:         res.Received,
	}, nil
}

func (k msgServer) CancelOrder(goCtx context.Context, msg *types.MsgCancelOrder) (*types.MsgCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := k.Keeper.CancelOrder(
		ctx, sdk.MustAccAddressFromBech32(msg.Sender), msg.OrderId)
	if err != nil {
		return nil, err
	}
	return &types.MsgCancelOrderResponse{}, nil
}

func (k msgServer) CancelAllOrders(goCtx context.Context, msg *types.MsgCancelAllOrders) (*types.MsgCancelAllOrdersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	orders, err := k.Keeper.CancelAllOrders(
		ctx, sdk.MustAccAddressFromBech32(msg.Sender), msg.MarketId)
	if err != nil {
		return nil, err
	}
	var cancelledOrderIds []uint64
	for _, order := range orders {
		cancelledOrderIds = append(cancelledOrderIds, order.Id)
	}
	return &types.MsgCancelAllOrdersResponse{CancelledOrderIds: cancelledOrderIds}, nil
}

func (k msgServer) SwapExactAmountIn(goCtx context.Context, msg *types.MsgSwapExactAmountIn) (*types.MsgSwapExactAmountInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	output, results, err := k.Keeper.SwapExactAmountIn(
		ctx, sdk.MustAccAddressFromBech32(msg.Sender), msg.Routes, msg.Input, msg.MinOutput, false)
	if err != nil {
		return nil, err
	}
	return &types.MsgSwapExactAmountInResponse{
		Output:  output,
		Results: results,
	}, nil
}
