package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelSellOrder = "cancel_sell_order"

var _ sdk.Msg = &MsgCancelSellOrder{}

func NewMsgCancelSellOrder(creator string, port string, channel string, amountDenom string, priceDenom string, orderID int32) *MsgCancelSellOrder {
	return &MsgCancelSellOrder{
		Creator:     creator,
		Port:        port,
		Channel:     channel,
		AmountDenom: amountDenom,
		PriceDenom:  priceDenom,
		OrderID:     orderID,
	}
}

func (msg *MsgCancelSellOrder) Route() string {
	return RouterKey
}

func (msg *MsgCancelSellOrder) Type() string {
	return TypeMsgCancelSellOrder
}

func (msg *MsgCancelSellOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelSellOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelSellOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
