package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	QueryParams             = "params"
	QueryNftListing         = "listing"
	QueryAllNftListings     = "listings"
	QueryNftListingsByOwner = "listings-by-owner"
	QueryNftAuction         = "auction"
	QueryAllNftAuctions     = "auctions"
	QueryNftAuctionBid      = "bid"
	QueryAllNftAuctionBids  = "bids"
	QueryNftAuctionsByOwner = "auctions-by-owner"
)

// QueryNftListingParams is the query parameters for '/marketplace/listings/{id}'
type QueryNftListingParams struct {
	Id string
}

// NewQueryNftListingParams
func NewQueryNftListingParams(id string) QueryNftListingParams {
	return QueryNftListingParams{
		Id: id,
	}
}

// QueryAllNftListingsParams is the query parameters for 'marketplace/listings'
type QueryAllNftListingsParams struct {
}

// NewQueryAllNftListingsParams
func NewQueryAllNftListingsParams() QueryAllNftListingsParams {
	return QueryAllNftListingsParams{}
}

// QueryNftListingsByOwnerParams is the query parameters for 'marketplace/listings/{owner}'
type QueryNftListingsByOwnerParams struct {
	Owner sdk.AccAddress
}

// NewQueryNftListingsByOwnerParams
func NewQueryNftListingsByOwnerParams(owner sdk.AccAddress) QueryNftListingsByOwnerParams {
	return QueryNftListingsByOwnerParams{
		Owner: owner,
	}
}

// QueryNftAuctionParams is the query parameters for '/marketplace/auctions/{id}'
type QueryNftAuctionParams struct {
	Id uint64
}

// NewQueryNftAuctionParams
func NewQueryNftAuctionParams(id uint64) QueryNftAuctionParams {
	return QueryNftAuctionParams{
		Id: id,
	}
}

// QueryAllNftListingsParams is the query parameters for 'marketplace/auctions'
type QueryAllNftAuctionsParams struct {
}

// NewQueryAllNftListingsParams
func NewQueryAllNftAuctionsParams() QueryAllNftAuctionsParams {
	return QueryAllNftAuctionsParams{}
}

// QueryNftListingsByOwnerParams is the query parameters for 'marketplace/auctions/{owner}'
type QueryNftAuctionsByOwnerParams struct {
	Owner sdk.AccAddress
}

// NewQueryNftAuctionsByOwnerParams
func NewQueryNftAuctionsByOwnerParams(owner sdk.AccAddress) QueryNftAuctionsByOwnerParams {
	return QueryNftAuctionsByOwnerParams{
		Owner: owner,
	}
}

// QueryNftAuctionBidParams is the query parameters for '/marketplace/bids/{id}'
type QueryNftAuctionBidParams struct {
	Id uint64
}

// NewQueryNftAuctionParams
func NewQueryNftAuctionBidParams(id uint64) QueryNftAuctionBidParams {
	return QueryNftAuctionBidParams{
		Id: id,
	}
}

// QueryAllNftAuctionBidsParams is the query parameters for 'marketplace/bids'
type QueryAllNftAuctionBidsParams struct {
}

// NewQueryAllNftListingsParams
func NewQueryAllNftAuctionBidsParams() QueryAllNftAuctionBidsParams {
	return QueryAllNftAuctionBidsParams{}
}
