package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/ons/types"
)

func (k msgServer) SellName(goCtx context.Context, msg *types.MsgSellName) (*types.MsgSellNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSellNameResponse{}, nil
}
