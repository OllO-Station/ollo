package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MsgRoute = "marketplace"

	TypeMsgListNft          = "list_nft"
	TypeMsgEditListing      = "edit_listing"
	TypeMsgDelistNft        = "de_list_nft"
	TypeMsgBuyNft           = "buy_nft"
	TypeMsgCreateAuction    = "create_auction"
	TypeMsgCancelNftAuction = "cancel_auction"
	TypeMsgPlaceNftBid      = "place_bid"

	// DoNotModify used to indicate that some field should not be updated
	DoNotModify = "[do-not-modify]"
	IdPrefix    = "list"
)

var (
	_ sdk.Msg = &MsgListNft{}
	_ sdk.Msg = &MsgEditNftListing{}
	_ sdk.Msg = &MsgDelistNft{}
	_ sdk.Msg = &MsgBuyNft{}
	_ sdk.Msg = &MsgCreateNftAuction{}
	_ sdk.Msg = &MsgCancelNftAuction{}
	_ sdk.Msg = &MsgPlaceNftBid{}
)

func NewMsgListNft(denomId, nftId string, price sdk.Coin, owner sdk.AccAddress) *MsgListNft {
	return &MsgListNft{
		Id:      GenUniqueID(IdPrefix),
		NftId:   nftId,
		DenomId: denomId,
		Price:   price,
		Seller:  owner.String(),
	}
}

func (msg MsgListNft) Route() string { return MsgRoute }

func (msg MsgListNft) Type() string { return TypeMsgListNft }

func (msg MsgListNft) ValidateBasic() error {
	return ValidateNftListing(
		NewNftListing(
			msg.Id,
			msg.NftId,
			msg.DenomId,
			msg.Price,
			sdk.AccAddress(msg.Seller),
		),
	)
}

// GetSignBytes Implements Msg.
func (msg MsgListNft) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgListNft) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgEditListing(id string, price sdk.Coin, owner sdk.AccAddress) *MsgEditNftListing {
	return &MsgEditNftListing{
		Id:    id,
		Price: price,
		Owner: owner.String(),
	}
}

func (msg MsgEditNftListing) Route() string { return MsgRoute }

func (msg MsgEditNftListing) Type() string { return TypeMsgEditListing }

func (msg MsgEditNftListing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return ValidatePrice(msg.Price)
}

// GetSignBytes Implements Msg.
func (msg MsgEditNftListing) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgEditNftListing) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgDelistNft
func NewMsgDelistNft(id string, owner sdk.AccAddress) *MsgDelistNft {
	return &MsgDelistNft{
		Id:    id,
		Owner: owner.String(),
	}
}

// Route Implements Msg.
func (msg MsgDelistNft) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgDelistNft) Type() string { return TypeMsgDelistNft }

// ValidateBasic Implements Msg.
func (msg MsgDelistNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgDelistNft) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgDelistNft) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgBuyNft
func NewMsgBuyNft(id string, price sdk.Coin, buyer sdk.AccAddress) *MsgBuyNft {
	return &MsgBuyNft{
		Id:    id,
		Price: price,
		Buyer: buyer.String(),
	}
}

// Route Implements Msg.
func (msg MsgBuyNft) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgBuyNft) Type() string { return TypeMsgBuyNft }

// ValidateBasic Implements Msg.
func (msg MsgBuyNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidatePrice(msg.Price)
}

// GetSignBytes Implements Msg.
func (msg MsgBuyNft) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgBuyNft) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// Auction messages

func NewMsgCreateAuction(denomId, nftId string, startTime time.Time, duration *time.Duration, startPrice sdk.Coin, owner sdk.AccAddress,
	incrementPercentage sdk.Dec) *MsgCreateNftAuction {
	return &MsgCreateNftAuction{
		NftId:               nftId,
		DenomId:             denomId,
		Duration:            duration,
		StartTime:           startTime,
		StartPrice:          startPrice,
		Owner:               owner.String(),
		IncrementPercentage: incrementPercentage,
	}
}

func (msg MsgCreateNftAuction) Route() string { return MsgRoute }

func (msg MsgCreateNftAuction) Type() string { return TypeMsgCreateAuction }

func (msg MsgCreateNftAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	if err = ValidatePrice(msg.StartPrice); err != nil {
		return err
	}
	if msg.Duration != nil {
		if err = ValidateDuration(msg.Duration); err != nil {
			return err
		}
	}
	if !msg.IncrementPercentage.IsPositive() || !msg.IncrementPercentage.LTE(sdk.NewDec(1)) {
		return sdkerrors.Wrapf(ErrInvalidPercentage, "invalid percentage value (%s)", msg.IncrementPercentage.String())
	}
	return nil
}
func (msg MsgCreateNftAuction) Validate(now time.Time) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	if msg.StartTime.Before(now) {
		return sdkerrors.Wrapf(ErrInvalidStartTime, "start time must be after current time %s", now.String())
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCreateNftAuction) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgCreateNftAuction) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgCancelNftAuction(auctionId uint64, owner sdk.AccAddress) *MsgCancelNftAuction {
	return &MsgCancelNftAuction{
		AuctionId: auctionId,
		Owner:     owner.String(),
	}
}

func (msg MsgCancelNftAuction) Route() string { return MsgRoute }

func (msg MsgCancelNftAuction) Type() string { return TypeMsgCancelNftAuction }

func (msg MsgCancelNftAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCancelNftAuction) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgCancelNftAuction) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgPlaceNftBid(auctionId uint64, amount sdk.Coin, bidder sdk.AccAddress) *MsgPlaceNftBid {
	return &MsgPlaceNftBid{
		AuctionId: auctionId,
		Amount:    amount,
		Bidder:    bidder.String(),
	}
}

func (msg MsgPlaceNftBid) Route() string { return MsgRoute }

func (msg MsgPlaceNftBid) Type() string { return TypeMsgPlaceNftBid }

func (msg MsgPlaceNftBid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bidder address (%s)", err)
	}
	if err := ValidatePrice(msg.Amount); err != nil {
		return err
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgPlaceNftBid) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgPlaceNftBid) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
