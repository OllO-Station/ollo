package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLiquidateLoan = "liquidate_loan"

var _ sdk.Msg = &MsgLiquidateLoan{}

func NewMsgLiquidateLoan(creator string, id uint64) *MsgLiquidateLoan {
	return &MsgLiquidateLoan{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgLiquidateLoan) Route() string {
	return RouterKey
}

func (msg *MsgLiquidateLoan) Type() string {
	return TypeMsgLiquidateLoan
}

func (msg *MsgLiquidateLoan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLiquidateLoan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLiquidateLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
