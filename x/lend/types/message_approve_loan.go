package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveLoan = "approve_loan"

var _ sdk.Msg = &MsgApproveLoan{}

func NewMsgApproveLoan(creator string, id uint64) *MsgApproveLoan {
	return &MsgApproveLoan{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgApproveLoan) Route() string {
	return RouterKey
}

func (msg *MsgApproveLoan) Type() string {
	return TypeMsgApproveLoan
}

func (msg *MsgApproveLoan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveLoan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
