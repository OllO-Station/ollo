package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ollo-station/ollo/x/lend/types"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.Status != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.Status)
	}
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amount, err := sdk.ParseCoinsNormalized(loan.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongLoanState, "Cannot parse coins in loan amount")
	}
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}
	loan.Lender = msg.Creator
	loan.Status = "approved"
	k.SetLoan(ctx, loan)
	return &types.MsgApproveLoanResponse{}, nil
}
