package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestLoan = "request_loan"

var _ sdk.Msg = &MsgRequestLoan{}

func NewMsgRequestLoan(creator string, amount string, collateral string, fee string, deadline string) *MsgRequestLoan {
	return &MsgRequestLoan{
		Creator:    creator,
		Amount:     amount,
		Collateral: collateral,
		Fee:        fee,
		Deadline:   deadline,
	}
}

func (msg *MsgRequestLoan) Route() string {
	return RouterKey
}

func (msg *MsgRequestLoan) Type() string {
	return TypeMsgRequestLoan
}

func (msg *MsgRequestLoan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestLoan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	amount, _ := sdk.ParseCoinsNormalized(msg.Amount)
	if !amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coins object")
	}
	if amount.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is empty")
	}
	fee, _ := sdk.ParseCoinsNormalized(msg.Fee)
	if !fee.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee is not a valid Coins object")
	}
	deadline, err := strconv.ParseInt(msg.Deadline, 10, 64)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "deadline is not an integer")
	}
	if deadline <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "deadline should be a positive integer")
	}
	collateral, _ := sdk.ParseCoinsNormalized(msg.Collateral)
	if !collateral.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral is not a valid Coins object")
	}
	if collateral.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral is empty")
	}
	return nil
}
