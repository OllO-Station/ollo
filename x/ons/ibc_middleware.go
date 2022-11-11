package ons

import (
	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	// "github.com/ignite/cli/ignite/pkg/cosmosanalysis/app"
	// host "github.com/cosmos/ibc-go/v5/modules/core/24-host/types"
	"ollo/x/ons/keeper"
	"ollo/x/ons/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	// exported "github.com/cosmos/ibc-go/v5/modules/core/exported"
)

type Middleware interface {
    IBCModule
    types.ICS4Wrapper
}
type IBCMiddleware struct {
	app    porttypes.IBCModule
	keeper keeper.Keeper    //add a keeper for stateful middleware
}

// IBCMiddleware creates a new IBCMiddleware given the associated keeper and underlying application
func NewIBCMiddleware(app porttypes.IBCModule, k keeper.Keeper) IBCMiddleware {
	return IBCMiddleware{
		app:    app,
		keeper: k,
	}
}
func (im IBCMiddleware) OnChanOpenInit(
    ctx sdk.Context,
    order channeltypes.Order,
    connectionHops []string,
    portID string,
    channelID string,
    channelCap *capabilitytypes.Capability,
    counterparty channeltypes.Counterparty,
    version string,
) (string, error) {
    if version != "" {
        // try to unmarshal JSON-encoded version string and pass
        // the app-specific version to app callback.
        // otherwise, pass version directly to app callback.
        // metadata, err := Unmarshal(version)
        // if err != nil {
        //     // Since it is valid for the fee version to not be specified,
        //     // the above middleware version may be for another middleware.
        //     // Pass the entire version string onto the underlying application.
            return im.app.OnChanOpenInit(
                ctx,
                order,
                connectionHops,
                portID,
                channelID,
                channelCap,
                counterparty,
                version,
            )
        } 
    // if the version string is empty, OnChanOpenInit is expected to return
    // a default version string representing the version(s) it supports
    appVersion, err := im.app.OnChanOpenInit(
        ctx,
        order,
        connectionHops,
        portID,
        channelID,
        channelCap,
        counterparty,
        version, // note you only pass app version here
    )
    if err != nil {
        return "", err
    }
    return appVersion, nil
}
