package market

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ollo-station/ollo/x/market/keeper"
	"github.com/ollo-station/ollo/x/market/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	// msgServer := keeper.NewMsgServerImpl(k)
	return func(c sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		// ctx = c.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		default:
			errMsg := fmt.Sprintf("Unrecognized %s market Msg type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
			// case *types.Msg

		}
	}
}
