package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/ollo-station/ollo/x/market/types"
)

// GetNextNftAuctionNumber get the next auction number
func (k Keeper) GetNextNftAuctionNumber(ctx sdk.Context) uint64 {
	var nextAuctionNumber uint64
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.PrefixNextNftAuctionNumber)
	if bz == nil {
		panic(fmt.Errorf("%s module not initialized -- Should have been done in InitGenesis", types.ModuleName))
	} else {
		val := gogotypes.UInt64Value{}

		err := k.cdc.Unmarshal(bz, &val)
		if err != nil {
			panic(err)
		}

		nextAuctionNumber = val.GetValue()
	}
	return nextAuctionNumber
}

// SetNextNftAuctionNumber set the next auction number
func (k Keeper) SetNextNftAuctionNumber(ctx sdk.Context, number uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: number})
	store.Set(types.PrefixNextNftAuctionNumber, bz)
}

// SetNftAuction set a specific auction listing in the store
func (k Keeper) SetNftAuction(ctx sdk.Context, auctionListing types.NftAuction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftAuctionId)
	bz := k.cdc.MustMarshal(&auctionListing)
	store.Set(types.KeyNftAuctionIdPrefix(auctionListing.Id), bz)
}

// GetNftAuction returns a auction listing by its id
func (k Keeper) GetNftAuction(ctx sdk.Context, id uint64) (val types.NftAuction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftAuctionId)
	b := store.Get(types.KeyNftAuctionIdPrefix(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetListing returns a listing from its nft id
func (k Keeper) GetNftAuctionIdByNftId(ctx sdk.Context, nftId string) (val uint64, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftAuctionNFTID)
	bz := store.Get(types.KeyNftAuctionNFTIDPrefix(nftId))
	if bz == nil {
		return val, false
	}
	var auctionId gogotypes.UInt64Value
	k.cdc.MustUnmarshal(bz, &auctionId)
	return auctionId.Value, true
}

// RemoveNftAuction removes a auction listing from the store
func (k Keeper) RemoveNftAuction(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftAuctionId)
	store.Delete(types.KeyNftAuctionIdPrefix(id))
}

// GetAllNftAuctions returns all auction listings
func (k Keeper) GetAllNftAuctions(ctx sdk.Context) (list []types.NftAuction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftAuctionId)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NftAuction
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetNftAuctionsByOwner returns all auction listings of specific owner
func (k Keeper) GetNftAuctionsByOwner(ctx sdk.Context, owner sdk.AccAddress) (auctionListings []types.NftAuction) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, append(types.PrefixNftAuctionOwner, owner.Bytes()...))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var id gogotypes.UInt64Value
		k.cdc.MustUnmarshal(iterator.Value(), &id)
		listing, found := k.GetNftAuction(ctx, id.Value)
		if !found {
			continue
		}
		auctionListings = append(auctionListings, listing)
	}

	return
}

func (k Keeper) HasNftAuction(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyNftAuctionIdPrefix(id))
}

func (k Keeper) SetNftAuctionWithOwner(ctx sdk.Context, owner sdk.AccAddress, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})

	store.Set(types.KeyNftAuctionOwnerPrefix(owner, id), bz)
}
func (k Keeper) UnsetNftAuctionWithOwner(ctx sdk.Context, owner sdk.AccAddress, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNftAuctionOwnerPrefix(owner, id))
}

func (k Keeper) SetNftAuctionWithNFTID(ctx sdk.Context, nftId string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftAuctionNFTID)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.KeyNftAuctionNFTIDPrefix(nftId), bz)
}

func (k Keeper) UnsetNftAuctionWithNFTID(ctx sdk.Context, nftId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftAuctionNFTID)
	store.Delete(types.KeyNftAuctionNFTIDPrefix(nftId))
}

func (k Keeper) SetNftAuctionWithPriceDenom(ctx sdk.Context, priceDenom string, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})

	store.Set(types.KeyNftAuctionPriceDenomPrefix(priceDenom, id), bz)
}

func (k Keeper) UnsetNftAuctionWithPriceDenom(ctx sdk.Context, priceDenom string, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNftAuctionPriceDenomPrefix(priceDenom, id))
}

func (k Keeper) SetInactiveNftAuction(ctx sdk.Context, auctionId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: auctionId})

	store.Set(types.KeyInActiveNftAuctionPrefix(auctionId), bz)
}

func (k Keeper) UnsetInactiveNftAuction(ctx sdk.Context, auctionId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyInActiveNftAuctionPrefix(auctionId))
}

func (k Keeper) SetActiveNftAuction(ctx sdk.Context, auctionId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: auctionId})

	store.Set(types.KeyActiveNftAuctionPrefix(auctionId), bz)
}

func (k Keeper) UnsetActiveNftAuction(ctx sdk.Context, auctionId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyActiveNftAuctionPrefix(auctionId))
}

func (k Keeper) IterateInactiveNftAuctions(ctx sdk.Context, fn func(index int, item types.NftAuction) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixInactiveNftAuction)
	iter := sdk.KVStorePrefixIterator(store, []byte{})
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var id gogotypes.UInt64Value
		k.cdc.MustUnmarshal(iter.Value(), &id)
		var (
			auction, _ = k.GetNftAuction(ctx, id.Value)
		)

		if stop := fn(i, auction); stop {
			break
		}
		i++
	}
}

func (k Keeper) IterateActiveNftAuctions(ctx sdk.Context, fn func(index int, item types.NftAuction) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixActiveNftAuction)
	iter := sdk.KVStorePrefixIterator(store, []byte{})
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var id gogotypes.UInt64Value
		k.cdc.MustUnmarshal(iter.Value(), &id)
		var (
			auction, _ = k.GetNftAuction(ctx, id.Value)
		)

		if stop := fn(i, auction); stop {
			break
		}
		i++
	}
}

// UpdateAuctionStatusesAndProcessBids update all auction listings status
func (k Keeper) UpdateAuctionStatusesAndProcessBids(ctx sdk.Context) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixNftAuctionId)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var auction types.NftAuction
		k.cdc.MustUnmarshal(iterator.Value(), &auction)
		if auction.StartTime.Before(ctx.BlockTime()) {
			bid, found := k.GetNftAuctionBid(ctx, auction.GetId())
			if !found && &auction.EndTime != nil && auction.EndTime.Before(ctx.BlockTime()) {
				err := k.nftKeeper.TransferNFT(ctx, auction.GetDenomId(), auction.GetNftId(),
					k.accountKeeper.GetModuleAddress(types.ModuleName), auction.GetOwner())
				if err != nil {
					return err
				}
				k.RemoveNftAuction(ctx, auction.GetId())
				k.RemoveNftAuctionEvent(ctx, auction)
			} else if !found && &auction.EndTime == nil &&
				ctx.BlockTime().Sub(auction.StartTime).Seconds() > k.GetBidCloseDuration(ctx).Seconds() {
				err := k.nftKeeper.TransferNFT(ctx, auction.GetDenomId(), auction.GetNftId(),
					k.accountKeeper.GetModuleAddress(types.ModuleName), auction.GetOwner())
				if err != nil {
					return err
				}
				k.RemoveNftAuction(ctx, auction.GetId())
				k.RemoveNftAuctionEvent(ctx, auction)

			} else if found && ctx.BlockTime().Sub(bid.Time).Seconds() > k.GetBidCloseDuration(ctx).Seconds() {
				err := k.processBid(ctx, auction, bid)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (k Keeper) processBid(ctx sdk.Context, auction types.NftAuction, bid types.NftAuctionBid) error {
	owner, err := sdk.AccAddressFromBech32(auction.Owner)
	if err != nil {
		return err
	}
	denom, err := k.nftKeeper.GetDenom(ctx, auction.DenomId)
	if err != nil {
		return err
	}
	nft, err := k.nftKeeper.GetNFT(ctx, auction.DenomId, auction.NftId)
	if err != nil {
		return err
	}
	BidAmountCoin := bid.Amount
	auctionSaleAmountCoin := BidAmountCoin
	err = k.nftKeeper.TransferNFT(ctx, auction.GetDenomId(), auction.GetNftId(),
		k.accountKeeper.GetModuleAddress(types.ModuleName), bid.GetBidder())
	if err != nil {
		return err
	}
	saleCommission := k.GetSaleCommission(ctx)
	marketplaceCoin := k.GetProportions(bid.Amount, saleCommission)
	if marketplaceCoin.Amount.GTE(sdk.OneInt()) {
		err = k.DistributeCommission(ctx, marketplaceCoin)
		if err != nil {
			return err
		}
		auctionSaleAmountCoin = BidAmountCoin.Sub(marketplaceCoin)
	}
	if nft.GetRoyaltyShare().GT(sdk.ZeroDec()) {
		nftRoyaltyShareCoin := k.GetProportions(auctionSaleAmountCoin, nft.GetRoyaltyShare())
		creator, err := sdk.AccAddressFromBech32(denom.Creator)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(nftRoyaltyShareCoin))
		if err != nil {
			return err
		}
		k.CreateRoyaltyShareTransferEvent(ctx, k.accountKeeper.GetModuleAddress(types.ModuleName), creator, nftRoyaltyShareCoin)
		auctionSaleAmountCoin = auctionSaleAmountCoin.Sub(nftRoyaltyShareCoin)
	}
	remaining := auctionSaleAmountCoin

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, sdk.NewCoins(remaining))
	if err != nil {
		return err
	}
	k.ProcessNftAuctionBidEvent(ctx, auction, bid)
	k.RemoveNftAuction(ctx, auction.GetId())
	k.RemoveNftAuctionBid(ctx, auction.GetId())
	return nil
}
