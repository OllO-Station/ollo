package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
	"time"
)

var (
	allowedDenoms = []string{}
)

// ValidateNftListing checks listing is valid or not
func ValidateNftListing(listing NftListing) error {
	if len(listing.Creator) > 0 {
		if _, err := sdk.AccAddressFromBech32(listing.Creator); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
		}
	}
	if err := ValidateId(listing.Id); err != nil {
		return err
	}
	if err := ValidatePrice(listing.Price); err != nil {
		return err
	}
	return nil
}

// ValidatePrice
func ValidatePrice(price sdk.Coin) error {
	if price.IsZero() || price.IsNegative() {
		return sdkerrors.Wrapf(ErrInvalidPrice, "invalid price %s, only accepts positive amount", price.String())
	}
	/*
		if !StringInSlice(price.Denom, allowedDenoms) {
			return sdkerrors.Wrapf(ErrInvalidPriceDenom, "invalid denom %s", price.Denom)
		}
	*/
	return nil
}

func ValidateDuration(t interface{}) error {
	duration, ok := t.(*time.Duration)
	if !ok {
		return sdkerrors.Wrapf(ErrInvalidDuration, "invalid value for duration: %T", t)
	}
	if duration.Nanoseconds() <= 0 {
		return sdkerrors.Wrapf(ErrInvalidDuration, "invalid duration %s, only accepts positive value", duration.String())
	}
	return nil
}

func ValidateId(id string) error {
	id = strings.TrimSpace(id)
	if len(id) < MinListingIdLength || len(id) > MaxListingIdLength {

		return sdkerrors.Wrapf(
			ErrInvalidNftListingId,
			"invalid id %s, only accepts value [%d, %d]", id, MinListingIdLength, MaxListingIdLength,
		)
	}
	if !IsBeginWithAlpha(id) || !IsAlphaNumeric(id) {
		return sdkerrors.Wrapf(ErrInvalidNftListingId, "invalid id %s, only accepts alphanumeric characters,and begin with an english letter", id)
	}
	return nil
}

func ValidateWhiteListAccounts(whitelistAccounts []string) error {
	if len(whitelistAccounts) > MaxWhitelistAccounts {
		return sdkerrors.Wrapf(ErrInvalidWhitelistAccounts,
			"number of whitelist accounts are more than the limit, len must be less than or equal to %d ", MaxWhitelistAccounts)
	}
	for _, address := range whitelistAccounts {
		_, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateIncrementPercentage(increment sdk.Dec) error {
	if !increment.IsPositive() || !increment.LTE(sdk.NewDec(1)) {
		return sdkerrors.Wrapf(ErrInvalidPercentage, "invalid percentage value (%s)", increment.String())
	}
	return nil
}

func validateNftAuctionId(id uint64) error {
	if id <= 0 {
		return sdkerrors.Wrapf(ErrInvalidNftAuctionId, "invalid auction id (%d)", id)
	}
	return nil
}

// ValidateNftAuctionNftListing checks auction listing is valid or not
func ValidateNftAuctionNftListing(auction NftAuction) error {
	if len(auction.Owner) > 0 {
		if _, err := sdk.AccAddressFromBech32(auction.Owner); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
		}
	}
	if err := validateNftAuctionId(auction.Id); err != nil {
		return err
	}
	if err := ValidatePrice(auction.StartPrice); err != nil {
		return err
	}
	return nil
}

// ValidateNftNftAuctionBid checks bid is valid or not
func ValidateNftNftAuctionBid(bid NftAuctionBid) error {
	if len(bid.Bidder) > 0 {
		if _, err := sdk.AccAddressFromBech32(bid.Bidder); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bidder address (%s)", bid.Bidder)
		}
	}
	if err := ValidatePrice(bid.Amount); err != nil {
		return err
	}
	if bid.Time.IsZero() {
		return sdkerrors.Wrapf(ErrInvalidTime, "invalid time (%s)", bid.Time.String())
	}
	return nil
}
