package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"ollo/x/ons/types"
)

func (k msgServer) SetName(goCtx context.Context, msg *types.MsgSetName) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting name information from the store
	name, _ := k.GetName(ctx, &types.QueryGetNameRequest{Name: msg.Name})

	// If the message sender address doesn't match the name owner, throw an error
	if !(msg.CreatorAddr == name.Name.OwnerAddr) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	// Otherwise, create an updated whois record
	newname := types.Name{
		OwnerAddr: msg.CreatorAddr,
		Id:     uint64(0),
		Name:      msg.Name,
		Value:     msg.Value,
		PricePaid:     name.Name.PricePaid,
	}

  fmt.Print(newname)
	// rite whois information to the store
	// k.SetName(ctx, newname)
	return &types.MsgSetNameResponse{}, nil
}
