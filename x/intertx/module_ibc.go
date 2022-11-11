// package intertx

// import (
// 	"fmt"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
// 	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
// 	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
// 	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
// 	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
// 	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"
// 	"ollo/x/intertx/keeper"
// 	"ollo/x/intertx/types"
// )

// type IBCModule struct {
// 	keeper keeper.Keeper
// }

// func NewIBCModule(k keeper.Keeper) IBCModule {
// 	return IBCModule{
// 		keeper: k,
// 	}
// }

// // OnChanOpenInit implements the IBCModule interface
// func (im IBCModule) OnChanOpenInit(
// 	ctx sdk.Context,
// 	order channeltypes.Order,
// 	connectionHops []string,
// 	portID string,
// 	channelID string,
// 	chanCap *capabilitytypes.Capability,
// 	counterparty channeltypes.Counterparty,
// 	version string,
// ) (string, error) {
//

// 	// Require portID is the portID module is bound to
// 	boundPort := im.keeper.GetPort(ctx)
// 	if boundPort != portID {
// 		return "", sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
// 	}

// 	if version != types.Version {
// 		return "", sdkerrors.Wrapf(types.ErrInvalidVersion, "got %s, expected %s", version, types.Version)
// 	}

// 	// Claim channel capability passed back by IBC module
// 	if err := im.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
// 		return "", err
// 	}

// 	return version, nil
// }

// // OnChanOpenTry implements the IBCModule interface
// func (im IBCModule) OnChanOpenTry(
// 	ctx sdk.Context,
// 	order channeltypes.Order,
// 	connectionHops []string,
// 	portID,
// 	channelID string,
// 	chanCap *capabilitytypes.Capability,
// 	counterparty channeltypes.Counterparty,
// 	counterpartyVersion string,
// ) (string, error) {
//

// 	// Require portID is the portID module is bound to
// 	boundPort := im.keeper.GetPort(ctx)
// 	if boundPort != portID {
// 		return "", sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
// 	}

// 	if counterpartyVersion != types.Version {
// 		return "", sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: got: %s, expected %s", counterpartyVersion, types.Version)
// 	}

// 	// Module may have already claimed capability in OnChanOpenInit in the case of crossing hellos
// 	// (ie chainA and chainB both call ChanOpenInit before one of them calls ChanOpenTry)
// 	// If module can already authenticate the capability then module already owns it so we don't need to claim
// 	// Otherwise, module does not have channel capability and we must claim it from IBC
// 	if !im.keeper.AuthenticateCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)) {
// 		// Only claim channel capability passed back by IBC module if we do not already own it
// 		if err := im.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
// 			return "", err
// 		}
// 	}

// 	return types.Version, nil
// }

// // OnChanOpenAck implements the IBCModule interface
// func (im IBCModule) OnChanOpenAck(
// 	ctx sdk.Context,
// 	portID,
// 	channelID string,
// 	_,
// 	counterpartyVersion string,
// ) error {
// 	if counterpartyVersion != types.Version {
// 		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: %s, expected %s", counterpartyVersion, types.Version)
// 	}
// 	return nil
// }

// // OnChanOpenConfirm implements the IBCModule interface
// func (im IBCModule) OnChanOpenConfirm(
// 	ctx sdk.Context,
// 	portID,
// 	channelID string,
// ) error {
// 	return nil
// }

// // OnChanCloseInit implements the IBCModule interface
// func (im IBCModule) OnChanCloseInit(
// 	ctx sdk.Context,
// 	portID,
// 	channelID string,
// ) error {
// 	// Disallow user-initiated channel closing for channels
// 	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
// }

// // OnChanCloseConfirm implements the IBCModule interface
// func (im IBCModule) OnChanCloseConfirm(
// 	ctx sdk.Context,
// 	portID,
// 	channelID string,
// ) error {
// 	return nil
// }

// // OnRecvPacket implements the IBCModule interface
// func (im IBCModule) OnRecvPacket(
// 	ctx sdk.Context,
// 	modulePacket channeltypes.Packet,
// 	relayer sdk.AccAddress,
// ) ibcexported.Acknowledgement {
// 	var ack channeltypes.Acknowledgement

// 	// this line is used by starport scaffolding # oracle/packet/module/recv

// 	var modulePacketData types.IntertxPacketData
// 	if err := modulePacketData.Unmarshal(modulePacket.GetData()); err != nil {
// 		return channeltypes.NewErrorAcknowledgement(sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error()))
// 	}

// 	// Dispatch packet
// 	switch packet := modulePacketData.Packet.(type) {
// 	// this line is used by starport scaffolding # ibc/packet/module/recv
// 	default:
// 		err := fmt.Errorf("unrecognized %s packet type: %T", types.ModuleName, packet)
// 		return channeltypes.NewErrorAcknowledgement(err)
// 	}

// 	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
//     return ack
// }

// // OnAcknowledgementPacket implements the IBCModule interface
// func (im IBCModule) OnAcknowledgementPacket(
// 	ctx sdk.Context,
// 	modulePacket channeltypes.Packet,
// 	acknowledgement []byte,
// 	relayer sdk.AccAddress,
// ) error {
// 	var ack channeltypes.Acknowledgement
// 	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet acknowledgement: %v", err)
// 	}

// 	// this line is used by starport scaffolding # oracle/packet/module/ack

// 	var modulePacketData types.IntertxPacketData
// 	if err := modulePacketData.Unmarshal(modulePacket.GetData()); err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
// 	}

// 	var eventType string

// 	// Dispatch packet
// 	switch packet := modulePacketData.Packet.(type) {
// 	// this line is used by starport scaffolding # ibc/packet/module/ack
// 	default:
// 		errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
// 		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
// 	}

//     ctx.EventManager().EmitEvent(
//         sdk.NewEvent(
//             eventType,
//             sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
//             sdk.NewAttribute(types.AttributeKeyAck, fmt.Sprintf("%v", ack)),
//         ),
//     )

// 	switch resp := ack.Response.(type) {
// 	case *channeltypes.Acknowledgement_Result:
// 		ctx.EventManager().EmitEvent(
// 			sdk.NewEvent(
// 				eventType,
// 				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
// 			),
// 		)
// 	case *channeltypes.Acknowledgement_Error:
// 		ctx.EventManager().EmitEvent(
// 			sdk.NewEvent(
// 				eventType,
// 				sdk.NewAttribute(types.AttributeKeyAckError, resp.Error),
// 			),
// 		)
// 	}

// 	return nil
// }

// // OnTimeoutPacket implements the IBCModule interface
// func (im IBCModule) OnTimeoutPacket(
// 	ctx sdk.Context,
// 	modulePacket channeltypes.Packet,
//     relayer sdk.AccAddress,
// ) error {
// 	var modulePacketData types.IntertxPacketData
// 	if err := modulePacketData.Unmarshal(modulePacket.GetData()); err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
// 	}

// 	// Dispatch packet
//     switch packet := modulePacketData.Packet.(type) {
//         // this line is used by starport scaffolding # ibc/packet/module/timeout
//         default:
//             errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
//             return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
//     }

//	    return nil
//	}
package intertx

import (
	"errors"

	proto "github.com/gogo/protobuf/proto"

	"ollo/x/intertx/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"
)

var _ porttypes.IBCModule = IBCModule{}

// IBCModule implements the ICS26 interface for interchain accounts controller chains
type IBCModule struct {
	keeper keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	return version, im.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID))
}

// OnChanOpenTry implements the IBCModule interface
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	return "", nil
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface. A successful acknowledgement
// is returned if the packet data is succesfully decoded and the receive application
// logic returns without error.
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	return channeltypes.NewErrorAcknowledgement(errors.New("cannot receive packet via interchain accounts authentication module"))
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-27 packet acknowledgement: %v", err)
	}

	txMsgData := &sdk.TxMsgData{}
	if err := proto.Unmarshal(ack.GetResult(), txMsgData); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-27 tx message data: %v", err)
	}

	switch len(txMsgData.Data) {
	case 0:
		// TODO: handle for sdk 0.46.x
		return nil
	default:
		for _, msgData := range txMsgData.Data {
			response, err := handleMsgData(ctx, msgData)
			if err != nil {
				return err
			}

			im.keeper.Logger(ctx).Info("message response in ICS-27 packet response", "response", response)
		}
		return nil
	}
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return nil
}

// NegotiateAppVersion implements the IBCModule interface
func (im IBCModule) NegotiateAppVersion(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionID string,
	portID string,
	counterparty channeltypes.Counterparty,
	proposedVersion string,
) (string, error) {
	return "", nil
}

func handleMsgData(ctx sdk.Context, msgData *sdk.MsgData) (string, error) {
	switch msgData.MsgType {
	case sdk.MsgTypeURL(&banktypes.MsgSend{}):
		msgResponse := &banktypes.MsgSendResponse{}
		if err := proto.Unmarshal(msgData.Data, msgResponse); err != nil {
			return "", sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal send response message: %s", err.Error())
		}

		return msgResponse.String(), nil

	// TODO: handle other messages

	default:
		return "", nil
	}
}
