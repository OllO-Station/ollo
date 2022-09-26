package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetName = "set_name"

var _ sdk.Msg = &MsgSetName{}

func NewMsgSetName(creator string, name string, value string) *MsgSetName {
	return &MsgSetName{
		Creator: creator,
		Name:    name,
		Value:   value,
	}
}

func (msg *MsgSetName) Route() string {
	return RouterKey
}

func (msg *MsgSetName) Type() string {
	return TypeMsgSetName
}

func (msg *MsgSetName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
