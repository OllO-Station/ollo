package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/market module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")

	ErrNftListingNotExists      = sdkerrors.Register(ModuleName, 2, "NftListing does not exist")
	ErrInvalidOwner             = sdkerrors.Register(ModuleName, 3, "invalid NftListing owner")
	ErrInvalidPrice             = sdkerrors.Register(ModuleName, 4, "invalid amount")
	ErrInvalidNftListing        = sdkerrors.Register(ModuleName, 5, "invalid NftListing")
	ErrNftListingAlreadyExists  = sdkerrors.Register(ModuleName, 6, "NftListing already exists")
	ErrNotEnoughAmount          = sdkerrors.Register(ModuleName, 7, "amount is not enough to buy")
	ErrInvalidPriceDenom        = sdkerrors.Register(ModuleName, 8, "invalid price denom")
	ErrInvalidNftListingId      = sdkerrors.Register(ModuleName, 9, "invalid NftListing id")
	ErrInvalidNftId             = sdkerrors.Register(ModuleName, 10, "invalid nft id")
	ErrNftNotExists             = sdkerrors.Register(ModuleName, 11, "nft not exists with given details")
	ErrUnauthorized             = sdkerrors.Register(ModuleName, 12, "unauthorized")
	ErrNftNonTransferable       = sdkerrors.Register(ModuleName, 13, "non-transferable nft")
	ErrNftListingDoesNotExists  = sdkerrors.Register(ModuleName, 14, "nft listing doesn't exists")
	ErrInvalidSplits            = sdkerrors.Register(ModuleName, 15, "invalid split shares")
	ErrNonPositiveNumber        = sdkerrors.Register(ModuleName, 16, "non positive number")
	ErrInvalidNftAuctionId      = sdkerrors.Register(ModuleName, 17, "invalid nft auction id")
	ErrInvalidWhitelistAccounts = sdkerrors.Register(ModuleName, 18, "invalid whitelist accounts")
	ErrNftAuctionDoesNotExists  = sdkerrors.Register(ModuleName, 19, "nft auction nft listing doesn't exists")
	ErrNftBidExists             = sdkerrors.Register(ModuleName, 20, "nft bid exists")
	ErrEndedNftAuction          = sdkerrors.Register(ModuleName, 21, "nft auction ended")
	ErrInActiveNftAuction       = sdkerrors.Register(ModuleName, 22, "inactive nft auction")
	ErrNftBidAmountNotEnough    = sdkerrors.Register(ModuleName, 23, "amount is not enough to nft bid")
	ErrNftBidDoesNotExists      = sdkerrors.Register(ModuleName, 24, "nft bid does not exists")
	ErrInvalidStartTime         = sdkerrors.Register(ModuleName, 25, "invalid start time")
	ErrInvalidPercentage        = sdkerrors.Register(ModuleName, 26, "invalid percentage decimal value")
	ErrInvalidTime              = sdkerrors.Register(ModuleName, 27, "invalid timestamp value")
	ErrInvalidDuration          = sdkerrors.Register(ModuleName, 28, "invalid duration")
)
