package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/ons/types"
)

func (k msgServer) AddThread(goCtx context.Context, msg *types.MsgAddThread) (*types.MsgAddThreadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddThreadResponse{}, nil
}
