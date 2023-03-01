package cli

import (
	"fmt"
	// "strings"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/x/reserve/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group reserve queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdGetDenomsFromCreator())
	cmd.AddCommand(CmdGetQueryDenom())
	cmd.AddCommand(CmdGetWhitelist())
	cmd.AddCommand(CmdGetQueryDenoms())
	cmd.AddCommand(CmdGetQueryParams())

	// this line is used by starport scaffolding # 1

	return cmd
}

// GetCmdQueryParams implements a command to return the current minting
// parameters.
func CmdGetQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current minting parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryParamsRequest{}
			res, err := queryClient.Params(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(fmt.Sprintf("%s\n", res.Params))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetQueryDenoms() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-denoms",
		Short: "",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetDenomsRequest{}
			res, err := queryClient.GetDenoms(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(fmt.Sprintf("%s\n", res.GetDenoms))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
func CmdGetQueryDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-denom",
		Short: "",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetDenomRequest{}
			res, err := queryClient.GetDenom(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(fmt.Sprintf("%s\n", res))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
func CmdGetDenomsFromCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "denoms-from-creator",
		Short: "",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDenomsFromCreatorRequest{}
			res, err := queryClient.DenomsFromCreator(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(fmt.Sprintf("%s\n", res))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
func CmdGetWhitelist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whitelist",
		Short: "",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetDenomWhitelistRequest{}
			res, err := queryClient.GetDenomWhitelist(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(fmt.Sprintf("%s\n", res))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
