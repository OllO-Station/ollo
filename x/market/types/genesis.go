package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

func NewGenesisState(listings []NftListing, listingCount uint64, params Params,
	auctions []NftAuction, bids []NftAuctionBid, nextAuctionNumber uint64) *GenesisState {
	return &GenesisState{
		Listings:          listings,
		ListingCount:      listingCount,
		Params:            params,
		Auctions:          auctions,
		Bids:              bids,
		NextAuctionNumber: nextAuctionNumber,
	}
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Listings:          []NftListing{},
		ListingCount:      0,
		Params:            Params{},
		Auctions:          []NftAuction{},
		Bids:              []NftAuctionBid{},
		NextAuctionNumber: 0,
	}
}

func (m *GenesisState) Validate() error {
	return m.ValidateGenesis()
}
func (m *GenesisState) ValidateGenesis() error {
	for _, l := range m.Listings {
		if l.GetOwner().Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing nft owner")
		}
		if err := ValidateNftListing(l); err != nil {
			return err
		}
	}
	if m.ListingCount < 0 {
		return sdkerrors.Wrap(ErrNonPositiveNumber, "must be a positive number")
	}
	if err := m.Params.ValidateBasic(); err != nil {
		return err
	}
	// for _, auction := range m.Auctions {
	// 	if err := validateNftAuctionId(auction); err != nil {
	// 		return err
	// 	}
	// }
	// for _, bid := range m.Bids {
	// 	if err := ValidateNftAuctionBid(bid); err != nil {
	// 		return err
	// 	}
	// }
	if m.NextAuctionNumber <= 0 {
		return sdkerrors.Wrap(ErrNonPositiveNumber, "must be a number and greater than 0.")
	}
	return nil
}
