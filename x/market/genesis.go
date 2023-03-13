package market

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ollo-station/ollo/x/market/keeper"
	"github.com/ollo-station/ollo/x/market/types"
)

// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.ValidateGenesis(); err != nil {
		panic(err.Error())
	}
	for _, l := range genState.Listings {
		k.SetNftListing(ctx, l)
		k.SetWithOwner(ctx, l.GetOwner(), l.GetId())
		k.SetWithNFTID(ctx, l.GetNftId(), l.GetId())
		k.SetWithPriceDenom(ctx, l.Price.GetDenom(), l.GetId())
	}
	k.SetNftListingCount(ctx, genState.ListingCount)
	k.SetParams(ctx, genState.Params)

	for _, al := range genState.Auctions {
		k.SetNftAuction(ctx, al)
		k.SetNftAuctionWithOwner(ctx, al.GetOwner(), al.GetId())
		k.SetNftAuctionWithNFTID(ctx, al.GetNftId(), al.GetId())
		k.SetNftAuctionWithPriceDenom(ctx, al.StartPrice.GetDenom(), al.GetId())
	}

	for _, b := range genState.Bids {
		k.SetNftAuctionBid(ctx, b)
	}
	k.SetNextNftAuctionNumber(ctx, genState.NextAuctionNumber)

	// check if the module account exists
	moduleAcc := k.GetMarketplaceAccount(ctx)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetAllNftListings(ctx),
		k.GetNftListingCount(ctx),
		k.GetParams(ctx),
		k.GetAllNftAuctions(ctx),
		k.GetAllNftAuctionBids(ctx),
		k.GetNextNftAuctionNumber(ctx),
	)
}

func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.NftListing{}, 0, types.DefaultParams(), []types.NftAuction{}, []types.NftAuctionBid{}, 1)
}
