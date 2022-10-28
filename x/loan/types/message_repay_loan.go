package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRepayLoan = "repay_loan"

var _ sdk.Msg = &MsgRepayLoan{}

func NewMsgRepayLoan(creator string, id uint64) *MsgRepayLoan {
	return &MsgRepayLoan{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRepayLoan) Route() string {
	return RouterKey
}

func (msg *MsgRepayLoan) Type() string {
	return TypeMsgRepayLoan
}

func (msg *MsgRepayLoan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRepayLoan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRepayLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
