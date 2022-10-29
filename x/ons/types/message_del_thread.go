package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteThread = "del_thread"

var _ sdk.Msg = &MsgDeleteThread{}

func NewMsgDeleteThread(creator string, name string, thread string, addr string) *MsgDeleteThread {
	return &MsgDeleteThread{
		Creator: creator,
		Name:    name,
		Thread:  thread,
		Addr:    addr,
	}
}

func (msg *MsgDeleteThread) Route() string {
	return RouterKey
}

func (msg *MsgDeleteThread) Type() string {
	return TypeMsgDeleteThread
}

func (msg *MsgDeleteThread) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteThread) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteThread) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
