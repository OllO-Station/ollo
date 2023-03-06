package ibc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
)

func GetTransferSenderRecipient(packet channeltypes.Packet) (string, string, error) {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return "", "", sdkerrors.ErrUnknownRequest.Wrap(
			"cannot unmarshal ICS-20 transfer packet data",
		)
	}
	if data.Sender == "" {
		return "", "", sdkerrors.ErrInvalidAddress.Wrap("missing sender address")
	}
	if data.Receiver == "" {
		return "", "", sdkerrors.ErrInvalidAddress.Wrap("missing receiver address")
	}
	return data.Sender, data.Receiver, nil
}

func GetTransferAmount(packet channeltypes.Packet) (string, error) {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return "", sdkerrors.ErrUnknownRequest.Wrap("cannot unmarshal ICS-20 transfer packet data")
	}
	if data.Amount == "" {
		return "", sdkerrors.ErrInvalidCoins.Wrap("cannot transfer zero coins")
	}
	if _, ok := sdk.NewIntFromString(data.Amount); !ok {
		return "", sdkerrors.ErrInvalidCoins.Wrap("cannot transfer invalid amount")
	}
	return data.Amount, nil
}
