package v6

// import (
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	storetypes "github.com/cosmos/cosmos-sdk/store/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/types/module"
// 	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
// 	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

// 	v6 "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/migrations/v6"
// )

// const (
// 	UpgradeName = "v6"
// )

// func CreateUpgradeHandler(
// 	mm *module.Manager,
// 	configurator module.Configurator,
// 	cdc codec.BinaryCodec,
// 	capabilityStoreKey *storetypes.KVStoreKey,
// 	capabilityKeeper *capabilitykeeper.Keeper,
// 	moduleName string,
// ) upgradetypes.UpgradeHandler {
// 	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
// 		if err := v6.MigrateICS27ChannelCapability(ctx, cdc, capabilityStoreKey, capabilityKeeper, moduleName); err != nil {
// 			return nil, err
// 		}

// 		return mm.RunMigrations(ctx, configurator, vm)
// 	}
// }
