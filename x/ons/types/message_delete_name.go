package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteName = "delete_name"

var _ sdk.Msg = &MsgDeleteName{}

func NewMsgDeleteName(creator string, name string) *MsgDeleteName {
	return &MsgDeleteName{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgDeleteName) Route() string {
	return RouterKey
}

func (msg *MsgDeleteName) Type() string {
	return TypeMsgDeleteName
}

func (msg *MsgDeleteName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
