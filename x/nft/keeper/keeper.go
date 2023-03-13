package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"

	// paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ollo-station/ollo/x/nft/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey storetypes.StoreKey // Unexposed key to access store from sdk.Context
	cdc      codec.Codec
	// paramstore paramtypes.Subspace
	nk nftkeeper.Keeper
}

// NewKeeper creates a new instance of the NFT Keeper
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	// ps paramtypes.Subspace,
	// set KeyTable if it has not already been set

	ak nft.AccountKeeper,
	bk nft.BankKeeper,
) Keeper {
	// if !ps.HasKeyTable() {
	// 	ps = ps.WithKeyTable(types.ParamKeyTable())
	// }
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
		// paramstore: ps,
		nk: nftkeeper.NewKeeper(storeKey, cdc, ak, bk),
	}
}

// NFTkeeper returns a cosmos-sdk nftkeeper.Keeper.
func (k Keeper) NFTkeeper() nftkeeper.Keeper {
	return k.nk
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("irismod/%s", types.ModuleName))
}

// func (k Keeper) GetDenom(ctx sdk.Context, id uint64) string {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := store.Get(types.GetDenomKey(id))
// 	if bz == nil {
// 		return ""
// 	}
// 	return string(bz)
// }
