// package keeper

// import (
// 	"fmt"

// 	"github.com/tendermint/tendermint/libs/log"
// 	storetypes "github.com/cosmos/cosmos-sdk/store/types"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
// 	"ollo/x/intertx/types"
// 	"github.com/ignite/cli/ignite/pkg/cosmosibckeeper"
// )

// type (
// 	Keeper struct {
// 		*cosmosibckeeper.Keeper
// 		cdc      	codec.BinaryCodec
// 		storeKey 	storetypes.StoreKey
// 		memKey   	storetypes.StoreKey
// 		paramstore	paramtypes.Subspace
// 		
// 	}
// )

// func NewKeeper(
//     cdc codec.BinaryCodec,
//     storeKey,
//     memKey storetypes.StoreKey,
// 	ps paramtypes.Subspace,
//     channelKeeper cosmosibckeeper.ChannelKeeper,
//     portKeeper cosmosibckeeper.PortKeeper,
//     scopedKeeper cosmosibckeeper.ScopedKeeper,
//     
// ) *Keeper {
// 	// set KeyTable if it has not already been set
// 	if !ps.HasKeyTable() {
// 		ps = ps.WithKeyTable(types.ParamKeyTable())
// 	}

// 	return &Keeper{
// 		Keeper: cosmosibckeeper.NewKeeper(
// 			types.PortKey,
// 			storeKey,
// 			channelKeeper,
// 			portKeeper,
// 			scopedKeeper,
// 		),
// 		cdc:      	cdc,
// 		storeKey: 	storeKey,
// 		memKey:   	memKey,
// 		paramstore:	ps,
// 		
// 	}
// }

package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/keeper"
	// icacontrollertypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"ollo/x/intertx/types"
)

type Keeper struct {
	cdc codec.Codec

	storeKey storetypes.StoreKey

	scopedKeeper        capabilitykeeper.ScopedKeeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, iaKeeper icacontrollerkeeper.Keeper, scopedKeeper capabilitykeeper.ScopedKeeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		scopedKeeper:        scopedKeeper,
		ICAControllerKeeper: iaKeeper,
	}
}

// Logger returns the applICAtion logger, scoped to the associated module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s-%s", host.ModuleName, types.ModuleName))
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}
