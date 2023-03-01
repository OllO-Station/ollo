package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ollo-station/ollo/x/reserve/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (server msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denom, err := server.Keeper.CreateDenom(ctx, msg.Sender, msg.Subdenom)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeMsgCreateDenom,
			sdk.NewAttribute(types.AttributeCreator, msg.Sender),
			sdk.NewAttribute(types.AttributeNewTokenDenom, denom),
		),
	})

	return &types.MsgCreateDenomResponse{
		NewTokenDenom: denom,
	}, nil
}

func (server msgServer) MintDenom(goCtx context.Context, msg *types.MsgMintDenom) (*types.MsgMintDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// pay some extra gas cost to give a better error here.
	_, denomExists := server.bankKeeper.GetDenomMetaData(ctx, msg.Amount.Denom)
	if !denomExists {
		return nil, types.ErrDenomDoesNotExist.Wrapf("denom: %s", msg.Amount.Denom)
	}

	authorityMetadata, err := server.Keeper.GetWhitelist(ctx, msg.Amount.GetDenom())
	if err != nil {
		return nil, err
	}

	for _, a := range authorityMetadata.GetAddresses() {
		if msg.Sender == a {
			err = server.Keeper.mintTo(ctx, msg.Amount, msg.Sender)
			if err != nil {
				return nil, err
			}

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.TypeMsgMintDenom,
					sdk.NewAttribute(types.AttributeMintToAddress, msg.Sender),
					sdk.NewAttribute(types.AttributeAmount, msg.Amount.String()),
				),
			})

			return &types.MsgMintDenomResponse{}, nil
		}
	}
	return nil, types.ErrUnauthorized

}

func (server msgServer) BurnDenom(goCtx context.Context, msg *types.MsgBurnDenom) (*types.MsgBurnDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorityMetadata, err := server.Keeper.GetWhitelist(ctx, msg.Amount.GetDenom())
	if err != nil {
		return nil, err
	}

	for _, a := range authorityMetadata.GetAddresses() {
		if msg.Sender == a {
			err = server.Keeper.burnFrom(ctx, msg.Amount, msg.Sender)
			if err != nil {
				return nil, err
			}

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.TypeMsgBurnDenom,
					sdk.NewAttribute(types.AttributeBurnFromAddress, msg.Sender),
					sdk.NewAttribute(types.AttributeAmount, msg.Amount.String()),
				),
			})

			return &types.MsgBurnDenomResponse{}, nil

		}

	}
	return nil, types.ErrUnauthorized

}

func (server msgServer) ForceTransfer(goCtx context.Context, msg *types.MsgForceTransfer) (*types.MsgForceTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorityMetadata, err := server.Keeper.GetWhitelist(ctx, msg.Amount.GetDenom())
	if err != nil {
		return nil, err
	}

	for _, a := range authorityMetadata.GetAddresses() {
		if msg.Sender == a {
			err = server.Keeper.forceTransfer(ctx, msg.Amount, msg.TransferFromAddress, msg.TransferToAddress)
			if err != nil {
				return nil, err
			}

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.TypeMsgForceTransfer,
					sdk.NewAttribute(types.AttributeTransferFromAddress, msg.TransferFromAddress),
					sdk.NewAttribute(types.AttributeTransferToAddress, msg.TransferToAddress),
					sdk.NewAttribute(types.AttributeAmount, msg.Amount.String()),
				),
			})

			return &types.MsgForceTransferResponse{}, nil
		}
	}
	return nil, types.ErrUnauthorized

}

func (server msgServer) ChangeDenomAdmin(goCtx context.Context, msg *types.MsgChangeDenomAdmin) (*types.MsgChangeDenomAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorityMetadata, err := server.Keeper.GetWhitelist(ctx, msg.Denom)
	if err != nil {
		return nil, err
	}

	for _, a := range authorityMetadata.GetAddresses() {
		if msg.Sender == a {
			err = server.Keeper.setAddresses(ctx, msg.Denom, msg.NewAdmin)
			if err != nil {
				return nil, err
			}
			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.TypeMsgChangeDenomAdmin,
					sdk.NewAttribute(types.AttributeDenom, msg.GetDenom()),
					sdk.NewAttribute(types.AttributeNewAdmin, msg.NewAdmin),
				),
			})

			return &types.MsgChangeDenomAdminResponse{}, nil
		}
	}
	return nil, types.ErrUnauthorized
}

func (server msgServer) SetDenomMetadata(goCtx context.Context, msg *types.MsgSetDenomMetadata) (*types.MsgSetDenomMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Defense in depth validation of metadata
	// err := msg.Metadata.Validate()
	// if err != nil {
	// 	return nil, err
	// }

	authorityMetadata, err := server.Keeper.GetWhitelist(ctx, msg.Metadata.Base)
	if err != nil {
		return nil, err
	}

	for _, a := range authorityMetadata.GetAddresses() {
		if msg.Sender == a {
			server.Keeper.bankKeeper.SetDenomMetaData(ctx, banktypes.Metadata{
				Description: msg.Metadata.Description,

				// DenomUnits: msg.Metadata.DenomUnits,
				Base:    msg.Metadata.Base,
				Display: msg.Metadata.Display,
				Name:    msg.Metadata.Name,
				Symbol:  msg.Metadata.Symbol,
				URI:     msg.Metadata.URI,
				URIHash: msg.Metadata.URIHash,
			})

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.TypeMsgSetDenomMetadata,
					sdk.NewAttribute(types.AttributeDenom, msg.Metadata.Base),
					sdk.NewAttribute(types.AttributeDenomMetadata, msg.Metadata.String()),
				),
			})

			return &types.MsgSetDenomMetadataResponse{}, nil
		}
	}

	return nil, types.ErrUnauthorized
}
