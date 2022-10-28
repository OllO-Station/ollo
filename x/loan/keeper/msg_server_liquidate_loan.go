package keeper

import (
	"context"

	"ollo/x/loan/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LiquidateLoan(goCtx context.Context, msg *types.MsgLiquidateLoan) (*types.MsgLiquidateLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgLiquidateLoanResponse{}, nil
}
