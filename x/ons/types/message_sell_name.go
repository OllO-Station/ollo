package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSellName = "sell_name"

var _ sdk.Msg = &MsgSellName{}

func NewMsgSellName(creator string, name string, sellerAddr string, offer string) *MsgSellName {
	return &MsgSellName{
		Creator:    creator,
		Name:       name,
		SellerAddr: sellerAddr,
		Offer:      offer,
	}
}

func (msg *MsgSellName) Route() string {
	return RouterKey
}

func (msg *MsgSellName) Type() string {
	return TypeMsgSellName
}

func (msg *MsgSellName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSellName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSellName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
