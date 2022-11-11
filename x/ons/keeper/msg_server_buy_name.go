package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"ollo/x/ons/types"
)

func (k msgServer) BuyName(goCtx context.Context, msg *types.MsgBuyName) (*types.MsgBuyNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting a name from the store
	name, err := k.GetName(ctx, &types.QueryGetNameRequest{msg.Name})

	// Set the price at which the name has to be bought if it didn't have an owner before
	minPrice := sdk.Coins{sdk.NewInt64Coin("token", 10)}

	// Convert price and bid strings to sdk.Coins
	price, _ := sdk.ParseCoinsNormalized(name.Name.PricePaid)
	bid, _ := sdk.ParseCoinsNormalized(msg.Bid)

	// Convert owner and buyer address strings to sdk.AccAddress
	owner, _ := sdk.AccAddressFromBech32(name.Name.OwnerAddr)
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)

	// If a name is found in store
	if err != nil {
		// If the current price is higher than the bid
		if price.IsAllGT(bid) {
			// Throw an error
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is not high enough")
		}

		// Otherwise (when the bid is higher), send tokens from the buyer to the owner
		err := k.bankKeeper.SendCoins(ctx, buyer, owner, bid)
		if err != nil {
			return nil, err
		}
	} else { // If the name is not found in the store
		// If the minimum price is higher than the bid
		if minPrice.IsAllGT(bid) {
			// Throw an error
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is less than min amount")
		}

		// Otherwise (when the bid is higher), send tokens from the buyer's account to the module's account (as a payment for the name)
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, bid)
		if err != nil {
			return nil, err
		}
	}

	// Create an updated name record
	// newname := types.Name{
	// 	Id:     uint64(0),
	// 	Name:      msg.Name,
	// 	Value:     "",
	// 	PricePaid:     bid.String(),
	// 	OwnerAddr: msg.BidderAddr,
	// }

	// Write name information to the store
	// k.SetName(ctx, newname)
	return &types.MsgBuyNameResponse{}, nil
}
