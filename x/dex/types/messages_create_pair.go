package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendCreatePair = "send_create_pair"

var _ sdk.Msg = &MsgSendCreatePair{}

func NewMsgSendCreatePair(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	sourceDenom string,
	targetDenom string,
) *MsgSendCreatePair {
	return &MsgSendCreatePair{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		SourceDenom:      sourceDenom,
		TargetDenom:      targetDenom,
	}
}

func (msg *MsgSendCreatePair) Route() string {
	return RouterKey
}

func (msg *MsgSendCreatePair) Type() string {
	return TypeMsgSendCreatePair
}

func (msg *MsgSendCreatePair) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendCreatePair) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendCreatePair) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
