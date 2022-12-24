package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaim = "claim"

var _ sdk.Msg = &MsgClaim{}

func NewMsgClaim(creator string, goalId uint64) *MsgClaim {
	return &MsgClaim{
		Claimer: creator,
		GoalId:  goalId,
	}
}

func (msg *MsgClaim) Route() string {
	return RouterKey
}

func (msg *MsgClaim) Type() string {
	return TypeMsgClaim
}

func (msg *MsgClaim) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaim) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Claimer); err != nil {
		return errors.Wrapf(errors.ErrPanic, "invalid creator address (%s)", err)
	}
	return nil
}
