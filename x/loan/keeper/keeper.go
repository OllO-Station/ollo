package keeper

import (
	"fmt"

	"ollo/x/loan/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	// host "github.com/cosmos/ibc-go/modules/core/24-host"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	// exported "github.com/cosmos/ibc-go/v6/modules/core/exported"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc             codec.BinaryCodec
		storeKey        storetypes.StoreKey
		memKey          storetypes.StoreKey
		paramstore      paramtypes.Subspace
		channelKeeper   types.ChannelKeeper
		portKeeper      types.PortKeeper
		scopedKeeper    capabilitykeeper.ScopedKeeper
		bankKeeper      types.BankKeeper
		accountKeeper   types.AccountKeeper
		stakingKeeper   types.StakingKeeper
		liquidityKeeper types.LiquidityKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	bankKeeper types.BankKeeper, accountKeeper types.AccountKeeper, stakingKeeper types.StakingKeeper, liquidityKeeper types.LiquidityKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		// Keeper: cosmosibckeeper.NewKeeper(
		// 	types.PortKey,
		// 	storeKey,
		// 	channelKeeper,
		// 	portKeeper,
		// ),
		scopedKeeper:  scopedKeeper,
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		paramstore:    ps,
		bankKeeper:    bankKeeper, accountKeeper: accountKeeper, stakingKeeper: stakingKeeper, liquidityKeeper: liquidityKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// IsBound checks if the transfer module is already bound to the desired port
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

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the transfer module that can claim a capability that IBC module
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}
