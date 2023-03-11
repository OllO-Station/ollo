package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	types "github.com/ollo-station/ollo/x/epoch/types"
	cobra "github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "epoch",
		Short: "Querying commands for the epoch module",
	}
	cmd.AddCommand(
		GetCmdListEpoch(),
		GetCmdGetCurrentEpoch(),
	)
	return cmd
}

func GetCmdListEpoch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Long:  "list all epoch",
		Short: "list all epoch",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			// clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			// if err != nil {
			// return err
			// }
			queryClient := types.NewQueryClient(clientCtx)
			e, err := queryClient.Epochs(context.Background(), &types.QueryEpochsRequest{})
			if err != nil {
				return err
			}
			fmt.Println(e)
			return nil
		},
	}
	return cmd
}

func GetCmdGetCurrentEpoch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-current [key]",
		Short: "Query the current epoch info",
		Long:  "Query the current epoch info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			// clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			// if err != nil {
			//     return err
			// }
			queryClient := types.NewQueryClient(clientCtx)
			e, err := queryClient.CurrentEpoch(context.Background(), &types.QueryCurrentEpochRequest{})
			if err != nil {
				return err
			}
			fmt.Println(e)
			return nil
		},
	}
	return cmd
}

// func GetCmdGetEpoch() *cobra.Command {
//     cmd := &cobra.Command{
//         Use:   "get-epoch [num]",
//         Short: "Query a epoch by its epoch number",
//         Long: "Query a epoch by its epoch number",
//         Args:  cobra.ExactArgs(1),
//         RunE: func(cmd *cobra.Command, args []string) error {
//             clientCtx := client.GetClientContextFromCmd(cmd)
//             clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
//             if err != nil {
//                 return err
//             }
//             queryClient := types.NewQueryClient(clientCtx)
//             queryClient.Epochs(context.Background(), &types.QueryEpochsRequest{})
//         },
//     }
//     return cmd
// }
