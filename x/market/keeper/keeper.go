package keeper

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	"github.com/ollo-station/ollo/x/market/types"

	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"

	// "github.com/cosmos/ibc-go/v6/modules/core/exported"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc            codec.BinaryCodec
		storeKey       storetypes.StoreKey
		memKey         storetypes.StoreKey
		paramstore     paramtypes.Subspace
		channelKeeper  types.ChannelKeeper
		portKeeper     types.PortKeeper
		scopedKeeper   capabilitykeeper.ScopedKeeper
		accountKeeper  types.AccountKeeper
		bankKeeper     types.BankKeeper
		feegrantKeeper types.FeeGrantKeeper
		groupKeeper    types.GroupKeeper
		nftKeeper      types.NFTKeeper
		onsKeeper      types.OnsKeeper
		authzKeeper    types.AuthzKeeper
		distrKeeper    types.DistributionKeeper
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
	accountKeeper types.AccountKeeper,
	authzKeeper types.AuthzKeeper,
	bankKeeper types.BankKeeper,
	feegrantKeeper types.FeeGrantKeeper,
	groupKeeper types.GroupKeeper,
	nftKeeper types.NFTKeeper,
	onsKeeper types.OnsKeeper,
	distrKeeper types.DistributionKeeper,
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
		// 	scopedKeeper,
		// ),
		cdc:            cdc,
		channelKeeper:  channelKeeper,
		portKeeper:     portKeeper,
		scopedKeeper:   scopedKeeper,
		storeKey:       storeKey,
		memKey:         memKey,
		paramstore:     ps,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		feegrantKeeper: feegrantKeeper,
		nftKeeper:      nftKeeper,
		onsKeeper:      onsKeeper,
		authzKeeper:    authzKeeper,
		distrKeeper:    distrKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// ChanCloseInit defines a wrapper function for the channel Keeper's function.
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.scopedKeeper.GetCapability(ctx, capName)
	if !ok {
		return sdkerrors.Wrapf(
			channeltypes.ErrChannelCapabilityNotFound,
			"could not retrieve channel capability at: %s",
			capName,
		)
	}
	return k.channelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// IsBound checks if the IBC app module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the IBC app module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the IBC app module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(
	ctx sdk.Context,
	cap *capabilitytypes.Capability,
	name string,
) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the IBC app module to claim a capability that core IBC
// passes to it
func (k Keeper) ClaimCapability(
	ctx sdk.Context,
	cap *capabilitytypes.Capability,
	name string,
) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

// AddNftListing adds a listing in the store and set owner to listing and updates the count
func (k Keeper) AddNftListing(ctx sdk.Context, listing types.NftListing) error {
	// check listing already exists
	if k.HasNftListing(ctx, listing.GetId()) {
		return sdkerrors.Wrapf(
			types.ErrNftListingAlreadyExists,
			"listing already exists: %s",
			listing.GetId(),
		)
	}

	err := k.nftKeeper.TransferNFT(ctx,
		listing.GetDenomId(), listing.GetNftId(), listing.GetOwner(),
		k.accountKeeper.GetModuleAddress(types.ModuleName))

	if err != nil {
		return err
	}
	// set listing
	k.SetNftListing(ctx, listing)
	if len(listing.GetOwner()) != 0 {
		// set listing id with owner prefix
		k.SetWithOwner(ctx, listing.GetOwner(), listing.GetId())
	}
	// Update listing count
	count := k.GetNftListingCount(ctx)
	k.SetNftListingCount(ctx, count+1)
	k.SetWithNFTID(ctx, listing.NftId, listing.Id)

	if len(listing.Price.Denom) > 0 {
		k.SetWithPriceDenom(ctx, listing.Price.Denom, listing.Id)
	}
	return nil
}

func (k Keeper) DeleteNftListing(ctx sdk.Context, listing types.NftListing) {
	k.RemoveNftListing(ctx, listing.GetId())
	k.UnsetWithOwner(ctx, listing.GetOwner(), listing.GetId())
	k.UnsetWithNFTID(ctx, listing.GetNftId())
	k.UnsetWithPriceDenom(ctx, listing.Price.Denom, listing.GetId())
}

func (k Keeper) Buy(ctx sdk.Context, listing types.NftListing, buyer sdk.AccAddress) error {
	owner, err := sdk.AccAddressFromBech32(listing.Creator)
	if err != nil {
		return err
	}
	denom, err := k.nftKeeper.GetDenom(ctx, listing.DenomId)
	if err != nil {
		return err
	}
	listingPriceCoin := listing.Price
	listingSaleAmountCoin := listingPriceCoin
	nft, err := k.nftKeeper.GetNFT(ctx, listing.DenomId, listing.NftId)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		buyer,
		types.ModuleName,
		sdk.NewCoins(listingPriceCoin),
	)
	if err != nil {
		return err
	}
	err = k.nftKeeper.TransferNFT(ctx, listing.GetDenomId(), listing.GetNftId(),
		k.accountKeeper.GetModuleAddress(types.ModuleName), buyer)
	if err != nil {
		_ = k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			buyer,
			sdk.NewCoins(listing.Price),
		)
		return err
	}
	saleCommission := k.GetSaleCommission(ctx)
	marketplaceCoin := k.GetProportions(listing.Price, saleCommission)
	if marketplaceCoin.Amount.GTE(sdk.OneInt()) {
		err = k.DistributeCommission(ctx, marketplaceCoin)
		if err != nil {
			return err
		}
		listingSaleAmountCoin = listingPriceCoin.Sub(marketplaceCoin)
	}
	if nft.GetRoyaltyShare().GT(sdk.ZeroDec()) {
		nftRoyaltyShareCoin := k.GetProportions(listingSaleAmountCoin, nft.GetRoyaltyShare())
		creator, err := sdk.AccAddressFromBech32(denom.Creator)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			creator,
			sdk.NewCoins(nftRoyaltyShareCoin),
		)
		if err != nil {
			return err
		}
		k.CreateRoyaltyShareTransferEvent(
			ctx,
			k.accountKeeper.GetModuleAddress(types.ModuleName),
			creator,
			nftRoyaltyShareCoin,
		)
		listingSaleAmountCoin = listingSaleAmountCoin.Sub(nftRoyaltyShareCoin)
	}
	remaining := listingSaleAmountCoin

	// if len(listing.SplitShares) > 0 {
	// 	for _, share := range listing.SplitShares {
	// 		sharePortionCoin := k.GetProportions(listingSaleAmountCoin, share.Weight)
	// 		sharePortionCoins := sdk.NewCoins(sharePortionCoin)
	// 		if share.Address == "" {
	// 			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, sharePortionCoins)
	// 			if err != nil {
	// 				return err
	// 			}
	// 		} else {
	// 			saleSplitAddr, err := sdk.AccAddressFromBech32(share.Address)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			err = k.bankKeeper.SendCoinsFromModuleToAccount(
	// 				ctx, types.ModuleName, saleSplitAddr, sharePortionCoins)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			k.CreateSplitShareTransferEvent(ctx, k.accountKeeper.GetModuleAddress(types.ModuleName), saleSplitAddr, sharePortionCoin)
	// 		}
	// 		remaining = remaining.Sub(sharePortionCoin)
	// 	}
	// 	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, sdk.NewCoins(remaining))
	// 	if err != nil {
	// 		return err
	// 	}
	// } else {
	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		owner,
		sdk.NewCoins(remaining),
	)
	if err != nil {
		return err
	}
	// }

	k.DeleteNftListing(ctx, listing)
	return nil
}

func (k Keeper) GetProportions(totalCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(
		totalCoin.Denom,
		sdk.NewDecFromInt(totalCoin.Amount).Mul(ratio).TruncateInt(),
	)
}

func (k Keeper) DistributeCommission(ctx sdk.Context, marketplaceCoin sdk.Coin) error {
	distrParams := k.GetMarketplaceDistributionParams(ctx)
	stakingCommissionCoin := k.GetProportions(marketplaceCoin, distrParams.Staking)
	if distrParams.Staking.GT(sdk.ZeroDec()) && stakingCommissionCoin.Amount.GT(sdk.ZeroInt()) {
		err := k.bankKeeper.SendCoinsFromModuleToModule(
			ctx,
			types.ModuleName,
			authtypes.FeeCollectorName,
			sdk.NewCoins(stakingCommissionCoin),
		)
		if err != nil {
			return err
		}
		k.CreateCommissionTransferEvent(ctx,
			k.accountKeeper.GetModuleAddress(types.ModuleName),
			k.accountKeeper.GetModuleAddress(authtypes.FeeCollectorName),
			stakingCommissionCoin,
		)
		marketplaceCoin = marketplaceCoin.Sub(stakingCommissionCoin)
	}
	communityPoolCommissionCoin := marketplaceCoin

	err := k.distrKeeper.FundCommunityPool(
		ctx,
		sdk.NewCoins(communityPoolCommissionCoin),
		k.accountKeeper.GetModuleAddress(types.ModuleName),
	)
	if err != nil {
		return err
	}
	k.CreateCommissionTransferEvent(ctx,
		k.accountKeeper.GetModuleAddress(types.ModuleName),
		k.accountKeeper.GetModuleAddress("distribution"),
		communityPoolCommissionCoin,
	)

	return nil
}

// CreateAuctionNftListing creates a auction in the store and set owner to auction and updates the next auction number
func (k Keeper) CreateNftAuction(ctx sdk.Context, auction types.NftAuction) error {

	// check auction already exists or not
	if k.HasNftAuction(ctx, auction.GetId()) {
		return sdkerrors.Wrapf(
			types.ErrNftListingAlreadyExists,
			"auction listing already exists: %s",
			auction.GetId(),
		)
	}

	err := k.nftKeeper.TransferNFT(ctx,
		auction.GetDenomId(), auction.GetNftId(), auction.GetOwner(),
		k.accountKeeper.GetModuleAddress(types.ModuleName))

	if err != nil {
		return err
	}
	// set auction listing
	k.SetNftAuction(ctx, auction)

	if len(auction.GetOwner()) != 0 {
		// set auction listing id with owner prefix
		k.SetNftAuctionWithOwner(ctx, auction.GetOwner(), auction.GetId())
	}
	// Update auction listing next number
	auctionId := k.GetNextNftAuctionNumber(ctx)
	k.SetNextNftAuctionNumber(ctx, auctionId+1)
	k.SetNftAuctionWithNFTID(ctx, auction.NftId, auction.Id)

	if len(auction.StartPrice.Denom) > 0 {
		k.SetNftAuctionWithPriceDenom(ctx, auction.StartPrice.Denom, auction.Id)
	}
	return nil
}

func (k Keeper) CancelNftAuction(ctx sdk.Context, auction types.NftAuction) error {
	// Check bid Exists or Not
	if k.HasNftAuctionBid(ctx, auction.Id) {
		return sdkerrors.Wrapf(
			types.ErrNftBidExists,
			"cannot cancel auction %d, bid exists ",
			auction.Id,
		)
	}

	// Transfer Back NFT ownership to auction owner
	err := k.nftKeeper.TransferNFT(ctx, auction.GetDenomId(), auction.GetNftId(),
		k.accountKeeper.GetModuleAddress(types.ModuleName), auction.GetOwner())
	if err != nil {
		return err
	}
	k.RemoveNftAuction(ctx, auction.GetId())
	k.UnsetNftAuctionWithOwner(ctx, auction.GetOwner(), auction.GetId())
	k.UnsetNftAuctionWithNFTID(ctx, auction.GetNftId())
	k.UnsetNftAuctionWithPriceDenom(ctx, auction.StartPrice.Denom, auction.GetId())

	return nil
}

func (k Keeper) PlaceNftBid(
	ctx sdk.Context,
	auction types.NftAuction,
	newBid types.NftAuctionBid,
) error {
	// Check bids of auction
	newBidPrice := auction.StartPrice
	prevBid, bidExists := k.GetNftAuctionBid(ctx, auction.Id)
	if bidExists {
		newBidPrice = k.GetNewBidPrice(
			auction.StartPrice.Denom,
			prevBid.Amount,
			auction.IncrementPercentage,
		)
	}
	if newBid.Amount.IsLT(newBidPrice) {
		return sdkerrors.Wrapf(
			types.ErrNftBidAmountNotEnough,
			"cannot place bid for given auction %d, required amount to bid is %s",
			auction.Id,
			newBidPrice.String(),
		)
	}

	// Transfer amount from bidder to module account
	err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		newBid.GetBidder(),
		types.ModuleName,
		sdk.NewCoins(newBid.Amount),
	)
	if err != nil {
		return err
	}
	// Release previous Bid
	if bidExists {
		_ = k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			prevBid.GetBidder(),
			sdk.NewCoins(prevBid.Amount),
		)
		k.RemoveNftAuctionBid(ctx, prevBid.AuctionId)
	}
	// Set new bid
	k.SetNftAuctionBid(ctx, newBid)

	return nil
}

func (k Keeper) GetNewBidPrice(denom string, amount sdk.Coin, increment sdk.Dec) sdk.Coin {
	return sdk.NewCoin(
		denom,
		amount.Amount.Add(sdk.NewDecFromInt(amount.Amount).Mul(increment).TruncateInt()),
	)
}
