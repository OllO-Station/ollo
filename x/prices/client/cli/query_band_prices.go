package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"ollo/x/prices/types"
)

// CmdBandPricesResult queries request result by reqID
func CmdBandPricesResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "band-prices-result [request-id]",
		Short: "Query the BandPrices result data by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			r, err := queryClient.BandPricesResult(context.Background(), &types.QueryBandPricesRequest{RequestId: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(r)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdLastBandPricesID queries latest request
func CmdLastBandPricesID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "last-band-prices-id",
		Short: "Query the last request id returned by BandPrices ack packet",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			r, err := queryClient.LastBandPricesId(context.Background(), &types.QueryLastBandPricesIdRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(r)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
