package keeper

import (
	"fmt"
	"strings"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ollo-station/ollo/x/market/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryParams:
			return queryParams(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryNftListing:
			return queryNftNftListing(ctx, req, k, legacyQuerierCdc)
		case types.QueryAllNftListings:
			return queryAllNftNftListings(ctx, req, k, legacyQuerierCdc)
		case types.QueryNftListingsByOwner:
			return queryNftNftListingsByOwner(ctx, req, k, legacyQuerierCdc)
		case types.QueryNftAuction:
			return queryNftAuction(ctx, req, k, legacyQuerierCdc)
		case types.QueryAllNftAuctions:
			return queryAllNftAuctions(ctx, req, k, legacyQuerierCdc)
		case types.QueryNftAuctionsByOwner:
			return queryNftAuctionsByOwner(ctx, req, k, legacyQuerierCdc)
		case types.QueryNftAuctionBid:
			return queryNftAuctionBid(ctx, req, k, legacyQuerierCdc)
		case types.QueryAllNftAuctionBids:
			return queryAllNftAuctionBids(ctx, req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryParams(ctx sdk.Context, _ []string, _ abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNftNftListing(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNftListingParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	id := strings.ToLower(strings.TrimSpace(params.Id))

	listing, found := k.GetNftListing(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNftListingDoesNotExists, fmt.Sprintf("listing %s does not exist", id))
	}
	return codec.MarshalJSONIndent(legacyQuerierCdc, listing)
}

func queryAllNftNftListings(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAllNftListingsParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	listings := k.GetAllNftListings(ctx)

	return codec.MarshalJSONIndent(legacyQuerierCdc, listings)
}

func queryNftNftListingsByOwner(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNftListingsByOwnerParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	listings := k.GetNftListingsByOwner(ctx, params.Owner)
	return codec.MarshalJSONIndent(legacyQuerierCdc, listings)
}

func queryNftAuction(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNftAuctionParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	auction, found := k.GetNftAuction(ctx, params.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNftAuctionDoesNotExists, fmt.Sprintf("auction %d does not exist", params.Id))
	}
	return codec.MarshalJSONIndent(legacyQuerierCdc, auction)
}

func queryAllNftAuctions(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAllNftAuctionsParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	auctions := k.GetAllNftAuctions(ctx)

	return codec.MarshalJSONIndent(legacyQuerierCdc, auctions)
}

func queryNftAuctionsByOwner(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNftAuctionsByOwnerParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	auctions := k.GetNftAuctionsByOwner(ctx, params.Owner)
	return codec.MarshalJSONIndent(legacyQuerierCdc, auctions)
}

func queryNftAuctionBid(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNftAuctionBidParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	bid, found := k.GetNftAuctionBid(ctx, params.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNftBidDoesNotExists, fmt.Sprintf("auction %d does not have any bid", params.Id))
	}
	return codec.MarshalJSONIndent(legacyQuerierCdc, bid)
}

func queryAllNftAuctionBids(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAllNftAuctionBidsParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	bids := k.GetAllNftAuctionBids(ctx)

	return codec.MarshalJSONIndent(legacyQuerierCdc, bids)
}
