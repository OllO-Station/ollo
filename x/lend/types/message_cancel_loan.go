package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelLoan = "cancel_loan"

var _ sdk.Msg = &MsgCancelLoan{}

func NewMsgCancelLoan(creator string, id uint64) *MsgCancelLoan {
	return &MsgCancelLoan{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCancelLoan) Route() string {
	return RouterKey
}

func (msg *MsgCancelLoan) Type() string {
	return TypeMsgCancelLoan
}

func (msg *MsgCancelLoan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelLoan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
