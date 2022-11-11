package oracle
import (
	"fmt"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"ollo/x/oracle/keeper"
	"ollo/x/oracle/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func NewFetchPriceHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.PricesProposal:
			return handlePriceProposal(ctx, k, c)

		default:
			return errors.New("")
		}
	}
}

func handlePriceProposal(ctx sdk.Context, k keeper.Keeper, p *types.PricesProposal) error {
	return k.HandleProposalPrice(ctx, p)
}
