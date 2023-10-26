package exchange

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ollo-station/ollo/x/exchange/keeper"
	"github.com/ollo-station/ollo/x/exchange/types"
)

// NewHandler returns a new msg handler.
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateMarket:
			res, err := msgServer.CreateMarket(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgPlaceLimitOrder:
			res, err := msgServer.PlaceLimitOrder(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgPlaceBatchLimitOrder:
			res, err := msgServer.PlaceBatchLimitOrder(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgPlaceMMLimitOrder:
			res, err := msgServer.PlaceMMLimitOrder(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgPlaceMMBatchLimitOrder:
			res, err := msgServer.PlaceMMBatchLimitOrder(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgPlaceMarketOrder:
			res, err := msgServer.PlaceMarketOrder(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCancelOrder:
			res, err := msgServer.CancelOrder(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSwapExactAmountIn:
			res, err := msgServer.SwapExactAmountIn(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}

// func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
// 	return func(ctx sdk.Context, content govtypes.Content) error {
// 		switch c := content.(type) {
// 		case *types.MarketParameterChangeProposal:
// 			return keeper.HandleMarketParameterChangeProposal(ctx, k, c)
// 		default:
// 			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized exchange proposal content type: %T", c)
// 		}
// 	}
// }
