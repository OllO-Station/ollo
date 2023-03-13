package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/ollo-station/ollo/x/market/types"
)

// SetBid set a specific bid for an auction listing in the store
func (k Keeper) SetNftAuctionBid(ctx sdk.Context, bid types.NftAuctionBid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixBidByNftAuctionId)
	bz := k.cdc.MustMarshal(&bid)
	store.Set(types.KeyBidPrefix(bid.AuctionId), bz)
}

// GetBid returns a bid of an auction listing by its id
func (k Keeper) GetNftAuctionBid(ctx sdk.Context, id uint64) (val types.NftAuctionBid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixBidByNftAuctionId)
	b := store.Get(types.KeyBidPrefix(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBid removes a bid of an auction listing from the store
func (k Keeper) RemoveNftAuctionBid(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixBidByNftAuctionId)
	store.Delete(types.KeyBidPrefix(id))
}

// GetAllBids returns all bids
func (k Keeper) GetAllNftAuctionBids(ctx sdk.Context) (list []types.NftAuctionBid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixBidByNftAuctionId)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftAuctionBid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBidsByBidder returns all bids of specific bidder
func (k Keeper) GetNftAuctionBidsByBidder(ctx sdk.Context, bidder sdk.AccAddress) (bids []types.NftAuctionBid) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, append(types.PrefixBidByNftBidder, bidder.Bytes()...))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var id gogotypes.UInt64Value
		k.cdc.MustUnmarshal(iterator.Value(), &id)
		bid, found := k.GetNftAuctionBid(ctx, id.Value)
		if !found {
			continue
		}
		bids = append(bids, bid)
	}

	return
}

func (k Keeper) HasNftAuctionBid(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixBidByNftAuctionId)
	return store.Has(types.KeyBidPrefix(id))
}
