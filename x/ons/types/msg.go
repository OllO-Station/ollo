package types

import (
	// "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)


var (
	_ sdk.Msg = &MsgTagName{}
	_ sdk.Msg = &MsgSetName{}
	_ sdk.Msg = &MsgSetOwnedName{}
	_ sdk.Msg = &MsgMute{}
	_ sdk.Msg = &MsgMessage{}
	_ sdk.Msg = &MsgReplyToThreadMessage{}
	_ sdk.Msg = &MsgBuyName{}
	_ sdk.Msg = &MsgDisableOwnedName{}
	_ sdk.Msg = &MsgEditNameParams{}
	_ sdk.Msg = &MsgEnableOwnedName{}
	_ sdk.Msg = &MsgFollow{}
	_ sdk.Msg = &MsgFollowTrades{}
	_ sdk.Msg = &MsgLoanOwnedName{}
	_ sdk.Msg = &MsgGiveOwnedName{}
)

// Message types for the liquidity module
const (
	TypeMsgTagName="tag_name"
	TypeMsgSetName="setname"
	TypeMsgSeOwnedtName="set"
	TypeMsgMute="mute"
	TypeMsgMessage="message"
	TypeMsgReplyToThreadMessage="reply"
	TypeMsgBuyName="buy"
	TypeMsgDisableOwnedName="disable"
	TypeMsgEditNameParams="edit"
	TypeMsgEnableOwnedName="enable"
	TypeMsgFollow="follow"
	TypeMsgFollowTrades="follow_Trades"
	TypeMsgLoanOwnedName="loan"
	TypeMsgGiveOwnedName="give"
)

func (msg *MsgReplyToThreadMessage) Route() string {
	return RouterKey
}

func (msg *MsgReplyToThreadMessage) Type() string {
	return TypeMsgSellName
}

func (msg *MsgReplyToThreadMessage) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgReplyToThreadMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReplyToThreadMessage) ValidateBasic() error {
	return nil
}

func (msg *MsgMessage) Route() string {
	return RouterKey
}

func (msg *MsgMessage) Type() string {
	return TypeMsgSellName
}

func (msg *MsgMessage) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMessage) ValidateBasic() error {
	return nil
}

func (msg *MsgEditNameParams) Route() string {
	return RouterKey
}

func (msg *MsgEditNameParams) Type() string {
	return TypeMsgSellName
}

func (msg *MsgEditNameParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgEditNameParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditNameParams) ValidateBasic() error {
	return nil
}


func (msg *MsgMute) Route() string {
	return RouterKey
}

func (msg *MsgMute) Type() string {
	return TypeMsgSellName
}

func (msg *MsgMute) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgMute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMute) ValidateBasic() error {
	return nil
}

func (msg *MsgFollowTrades) Route() string {
	return RouterKey
}

func (msg *MsgFollowTrades) Type() string {
	return TypeMsgSellName
}

func (msg *MsgFollowTrades) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgFollowTrades) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFollowTrades) ValidateBasic() error {
	return nil
}

func (msg *MsgFollow) Route() string {
	return RouterKey
}

func (msg *MsgFollow) Type() string {
	return TypeMsgSellName
}

func (msg *MsgFollow) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgFollow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFollow) ValidateBasic() error {
	return nil
}

func (msg *MsgGiveOwnedName) Route() string {
	return RouterKey
}

func (msg *MsgGiveOwnedName) Type() string {
	return TypeMsgSellName
}

func (msg *MsgGiveOwnedName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgGiveOwnedName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGiveOwnedName) ValidateBasic() error {
	return nil
}
func (msg *MsgLoanOwnedName) Route() string {
	return RouterKey
}

func (msg *MsgLoanOwnedName) Type() string {
	return TypeMsgSellName
}

func (msg *MsgLoanOwnedName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgLoanOwnedName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLoanOwnedName) ValidateBasic() error {
	return nil
}
func (msg *MsgDisableOwnedName) Route() string {
	return RouterKey
}

func (msg *MsgDisableOwnedName) Type() string {
	return TypeMsgSellName
}

func (msg *MsgDisableOwnedName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgDisableOwnedName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisableOwnedName) ValidateBasic() error {
	return nil
}
func (msg *MsgEnableOwnedName) Route() string {
	return RouterKey
}

func (msg *MsgEnableOwnedName) Type() string {
	return TypeMsgSellName
}

func (msg *MsgEnableOwnedName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgEnableOwnedName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableOwnedName) ValidateBasic() error {
	return nil
}
func (msg *MsgTagName) Route() string {
	return RouterKey
}

func (msg *MsgTagName) Type() string {
	return TypeMsgSellName
}

func (msg *MsgTagName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgTagName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTagName) ValidateBasic() error {
	return nil
}
func (msg *MsgSetOwnedName) Route() string {
	return RouterKey
}

func (msg *MsgSetOwnedName) Type() string {
	return TypeMsgSellName
}

func (msg *MsgSetOwnedName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgSetOwnedName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetOwnedName) ValidateBasic() error {
	return nil
}
