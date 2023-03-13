package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/ollo-station/ollo/x/market/types"
)

// GetNftListingCount get the total number of listings
func (k Keeper) GetNftListingCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.PrefixNftListingsCount
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetNftListingCount set the total number of listings
func (k Keeper) SetNftListingCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.PrefixNftListingsCount
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetNftListing set a specific listing in the store
func (k Keeper) SetNftListing(ctx sdk.Context, listing types.NftListing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftListingId)
	b := k.cdc.MustMarshal(&listing)
	store.Set(types.KeyNftListingIdPrefix(listing.Id), b)
}

// GetNftListing returns a listing from its id
func (k Keeper) GetNftListing(ctx sdk.Context, id string) (val types.NftListing, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftListingId)
	b := store.Get(types.KeyNftListingIdPrefix(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetNftListing returns a listing from its nft id
func (k Keeper) GetNftListingIdByNftId(ctx sdk.Context, nftId string) (val string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftListingNFTID)
	b := store.Get(types.KeyNftListingNFTIDPrefix(nftId))
	if b == nil {
		return val, false
	}
	var listingId gogotypes.StringValue
	k.cdc.MustUnmarshal(b, &listingId)
	return listingId.Value, true
}

// RemoveListing removes a listing from the store
func (k Keeper) RemoveNftListing(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftListingId)
	store.Delete(types.KeyNftListingIdPrefix(id))
}

// GetAllNftListings returns all listings
func (k Keeper) GetAllNftListings(ctx sdk.Context) (list []types.NftListing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftListingId)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftListing
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetNftListingsByOwner returns all listings of specific owner
func (k Keeper) GetNftListingsByOwner(ctx sdk.Context, owner sdk.AccAddress) (listings []types.NftListing) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, append(types.PrefixNftListingOwner, owner.Bytes()...))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var id gogotypes.StringValue
		k.cdc.MustUnmarshal(iterator.Value(), &id)
		listing, found := k.GetNftListing(ctx, id.Value)
		if !found {
			continue
		}
		listings = append(listings, listing)
	}

	return
}

func (k Keeper) HasNftListing(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyNftListingIdPrefix(id))
}

func (k Keeper) SetWithOwner(ctx sdk.Context, owner sdk.AccAddress, id string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.StringValue{Value: id})

	store.Set(types.KeyNftListingOwnerPrefix(owner, id), bz)
}
func (k Keeper) UnsetWithOwner(ctx sdk.Context, owner sdk.AccAddress, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNftListingOwnerPrefix(owner, id))
}

func (k Keeper) SetWithNFTID(ctx sdk.Context, nftId, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftListingNFTID)
	bz := k.cdc.MustMarshal(&gogotypes.StringValue{Value: id})
	store.Set(types.KeyNftListingNFTIDPrefix(nftId), bz)
}
func (k Keeper) UnsetWithNFTID(ctx sdk.Context, nftId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftListingNFTID)
	store.Delete(types.KeyNftListingNFTIDPrefix(nftId))
}

func (k Keeper) SetWithPriceDenom(ctx sdk.Context, priceDenom, id string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.StringValue{Value: id})

	store.Set(types.KeyNftListingPriceDenomPrefix(priceDenom, id), bz)
}
func (k Keeper) UnsetWithPriceDenom(ctx sdk.Context, priceDenom, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNftListingPriceDenomPrefix(priceDenom, id))
}
