package keeper

import (
  "context"
	"fmt"

	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/codec"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	// channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	// porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	// ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"
	"github.com/tendermint/tendermint/libs/log"
	"ollo/x/ons/types"
)
type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		groupKeeper   types.GroupKeeper
	ics4Wrapper types.ICS4Wrapper
		channelKeeper   types.ChannelKeeper
    portKeeper types.PortKeeper
    scopedKeeper types.ScopedKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper, accountKeeper types.AccountKeeper, groupKeeper types.GroupKeeper,
  ics4Wrapper types.ICS4Wrapper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper types.ScopedKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
    ics4Wrapper: ics4Wrapper,
		bankKeeper: bankKeeper, 
    accountKeeper: accountKeeper, 
    groupKeeper: groupKeeper,
    channelKeeper: channelKeeper,
    portKeeper: portKeeper,
    scopedKeeper: scopedKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// IsBound checks if the IBC query module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the ort Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the transfer module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the transfer module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}
// ClaimCapability allows the transfer module that can claim a capability that IBC module
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
    return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

//   func(k Keeper) BuyName(ctx context.Context, _ *types.MsgBuyName) (*types.MsgBuyNameResponse, error) { 
//     return &types.MsgBuyNameResponse{}, nil
// }
//   func(k Keeper) SellName(ctx context.Context, _ *types.MsgSellName) (*types.MsgSellNameResponse, error) { 
//     return &types.MsgSellNameResponse{}, nil 
// }
//   func(k Keeper) SetName(ctx context.Context, _ *types.MsgSetName) (*types.MsgSetNameResponse, error) {
//     return &types.MsgSetNameResponse{}, nil
// }
  func(k Keeper) TagName(ctx context.Context, _ *types.MsgTagName) (*types.MsgTagNameResponse, error) {
    return &types.MsgTagNameResponse{}, nil
}
  func(k Keeper) SetOwnedName(ctx context.Context, _ *types.MsgSetOwnedName) (*types.MsgSetOwnedNameResponse, error) {
    return &types.MsgSetOwnedNameResponse{}, nil
}
  func(k Keeper) EnableOwnedName(ctx context.Context, _ *types.MsgEnableOwnedName) (*types.MsgEnableOwnedNameResponse, error) {
    return &types.MsgEnableOwnedNameResponse{}, nil
}
  func(k Keeper) DisableOwnedName(ctx context.Context, _ *types.MsgDisableOwnedName) (*types.MsgDisableOwnedNameResponse, error) {
    return &types.MsgDisableOwnedNameResponse{}, nil
}
  // func(k Keeper) DeleteName(ctx context.Context, _ *types.MsgDeleteName) (*types.MsgDeleteNameResponse, error) {
  //   return &types.MsgDeleteNameResponse{}, nil
// }
//   func(k Keeper) AddThread(ctx context.Context, _ *types.MsgAddThread) (*types.MsgAddThreadResponse, error) {
//     return &types.MsgAddThreadResponse{}, nil
// }
  func(k Keeper) ReplyToThreadMessage(ctx context.Context, _ *types.MsgReplyToThreadMessage) (*types.MsgReplyToThreadMessageResponse, error) {
    return &types.MsgReplyToThreadMessageResponse{}, nil
}
  func(k Keeper) Follow(ctx context.Context, _ *types.MsgFollow) (*types.MsgFollowResponse, error) {
    return &types.MsgFollowResponse{}, nil
}
  func(k Keeper) FollowTrades(ctx context.Context, _ *types.MsgFollowTrades) (*types.MsgFollowTradesResponse, error) {
    return &types.MsgFollowTradesResponse{}, nil
}
  func(k Keeper) Mute(ctx context.Context, _ *types.MsgMute) (*types.MsgMuteResponse, error) {
    return &types.MsgMuteResponse{}, nil
}
  func(k Keeper) Message(ctx context.Context, _ *types.MsgMessage) (*types.MsgMessageResponse, error) {
    return &types.MsgMessageResponse{}, nil
}
//   func(k Keeper) DeleteThread(ctx context.Context, _ *types.MsgDeleteThread) (*types.MsgDeleteThreadResponse, error) {
//     return &types.MsgDeleteThreadResponse{}, nil
// }
