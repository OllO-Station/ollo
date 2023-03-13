package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"strconv"
	"strings"

	"github.com/ollo-station/ollo/x/market/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group marketplace queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdQueryNftListing(),
		GetCmdQueryAllNftListings(),
		GetCmdQueryNftListingsByOwner(),
		GetCmdQueryNftAuction(),
		GetCmdQueryAllNftAuctions(),
		GetCmdQueryNftAuctionsByOwner(),
		GetCmdQueryNftAuctionBid(),
		GetCmdQueryAllBids(),
	)

	return cmd
}

// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query Marketplace params",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryNftListing implements the query listing command.
func GetCmdQueryNftListing() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "listing [id]",
		Long:    "Query a listing by id.",
		Example: fmt.Sprintf("$ %s query marketplace listing <id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}

			listingId := strings.ToLower(strings.TrimSpace(args[0]))

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NftListing(context.Background(), &types.QueryNftListingRequest{
				Id: listingId,
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Listing)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryAllNftListings implements the query all listings command.
func GetCmdQueryAllNftListings() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "listings",
		Long:    "Query listings.",
		Example: fmt.Sprintf("$ %s query marketplace listings", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}
			owner, err := cmd.Flags().GetString(FlagOwner)
			if err != nil {
				return err
			}
			priceDenom, err := cmd.Flags().GetString(FlagPriceDenom)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			resp, err := queryClient.NftListings(
				context.Background(),
				&types.QueryNftListingsRequest{
					Owner:      owner,
					Denom:      priceDenom,
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(FlagOwner, "", "filter by owner address")
	cmd.Flags().String(FlagPriceDenom, "", "filter by listing price-denom")
	flags.AddPaginationFlagsToCmd(cmd, "all listings")

	return cmd
}

// GetCmdQueryNftListingsByOwner implements the query listings by owner command.
func GetCmdQueryNftListingsByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "listings-by-owner [owner]",
		Long:    "Query listings by the owner.",
		Example: fmt.Sprintf("$ %s query marketplace listings <owner>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			var owner sdk.AccAddress
			if len(args) > 0 {
				owner, err = sdk.AccAddressFromBech32(args[0])
				if err != nil {
					return err
				}
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			resp, err := queryClient.NftListingsByOwner(
				context.Background(),
				&types.QueryNftListingsByOwnerRequest{
					Owner:      owner.String(),
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "owner listings")

	return cmd
}

// GetCmdQueryNftAuctionNftListing implements the query auction command.
func GetCmdQueryNftAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auction [id]",
		Long:    "Query a auction by id.",
		Example: fmt.Sprintf("$ %s query marketplace auction <id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}

			auctionId, err := strconv.ParseUint(strings.ToLower(strings.TrimSpace(args[0])), 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NftAuction(context.Background(), &types.QueryNftAuctionRequest{
				Id: auctionId,
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Auction)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryAllNftAuctions implements the query all auctions command.
func GetCmdQueryAllNftAuctions() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auctions",
		Long:    "Query auctions.",
		Example: fmt.Sprintf("$ %s query marketplace auctions", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}
			owner, err := cmd.Flags().GetString(FlagOwner)
			if err != nil {
				return err
			}
			priceDenom, err := cmd.Flags().GetString(FlagPriceDenom)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			resp, err := queryClient.NftAuctions(
				context.Background(),
				&types.QueryNftAuctionsRequest{
					Owner:      owner,
					Denom:      priceDenom,
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(FlagOwner, "", "filter by owner address")
	cmd.Flags().String(FlagPriceDenom, "", "filter by auction price-denom")
	flags.AddPaginationFlagsToCmd(cmd, "all auctions")

	return cmd
}

// GetCmdQueryNftAuctionsByOwner implements the query auctions by owner command.
func GetCmdQueryNftAuctionsByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auctions-by-owner [owner]",
		Long:    "Query auctions by the owner.",
		Example: fmt.Sprintf("$ %s query marketplace auctions <owner>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			var owner sdk.AccAddress
			if len(args) > 0 {
				owner, err = sdk.AccAddressFromBech32(args[0])
				if err != nil {
					return err
				}
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			resp, err := queryClient.NftAuctionsByOwner(
				context.Background(),
				&types.QueryNftAuctionsByOwnerRequest{
					Owner:      owner.String(),
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "owner auctions")

	return cmd
}

// GetCmdQueryNftAuctionBid implements the query bid command.
func GetCmdQueryNftAuctionBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bid [id]",
		Long:    "Query a bid by auction id.",
		Example: fmt.Sprintf("$ %s query marketplace bid <id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}

			auctionId, err := strconv.ParseUint(strings.ToLower(strings.TrimSpace(args[0])), 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NftAuctionBid(context.Background(), &types.QueryNftAuctionBidRequest{
				Id: auctionId,
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Bid)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryAllBids implements the query all bids command.
func GetCmdQueryAllBids() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bids",
		Long:    "Query bids.",
		Example: fmt.Sprintf("$ %s query marketplace bids", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}
			bidder, err := cmd.Flags().GetString(FlagBidder)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			resp, err := queryClient.NftAuctionBids(
				context.Background(),
				&types.QueryNftAuctionBidsRequest{
					Bidder:     bidder,
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(FlagBidder, "", "filter by bidder address")
	flags.AddPaginationFlagsToCmd(cmd, "all bids")

	return cmd
}
