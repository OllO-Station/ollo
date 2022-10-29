package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddThread = "add_thread"

var _ sdk.Msg = &MsgAddThread{}

func NewMsgAddThread(creator string, name string, thread string, addr string, offer string) *MsgAddThread {
	return &MsgAddThread{
		Creator: creator,
		Name:    name,
		Thread:  thread,
		Addr:    addr,
		Offer:   offer,
	}
}

func (msg *MsgAddThread) Route() string {
	return RouterKey
}

func (msg *MsgAddThread) Type() string {
	return TypeMsgAddThread
}

func (msg *MsgAddThread) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddThread) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddThread) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
