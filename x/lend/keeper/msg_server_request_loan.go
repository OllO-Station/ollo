package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ollo-station/ollo/x/lend/types"
)

func (k msgServer) RequestLoan(goCtx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var loan = types.Loan{
		Borrower:   msg.Creator,
		Status:     "requested",
		Amount:     msg.Amount,
		Fee:        msg.Fee,
		Collateral: msg.Collateral,
		Deadline:   msg.Deadline,
	}
	borrower, e := sdk.AccAddressFromBech32(msg.Creator)
	if e != nil {
		panic(e)
	}
	collateral, e := sdk.ParseCoinNormalized(loan.Collateral)
	if e != nil {
		panic(e)
	}
	sdkerr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, sdk.Coins{collateral})
	if sdkerr != nil {
		return nil, sdkerr
	}
	k.AppendLoan(ctx, loan)
	return &types.MsgRequestLoanResponse{}, nil
}
