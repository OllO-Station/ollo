package cli

import (
	"fmt"

	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ollo-station/ollo/x/reserve/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	return cmd
}
func CmdCreateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create [subdenom]",
		Args: cobra.ExactArgs(1),
		// Short: "Remove a terminated private plan",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress()

			msg := types.NewMsgCreateDenom(string(creator), args[0])

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func CmdMintDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mint [amount] [denom]",
		Args: cobra.ExactArgs(2),
		// Short: "Remove a terminated private plan",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress()

			coin, err := sdk.ParseCoinNormalized(string(args[0]))
			if err != nil {
				return err
			}
			msg := types.NewMsgMintDenom(string(creator), coin)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func CmdBurnDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "burn [amount][denom]",
		Args: cobra.ExactArgs(1),
		// Short: "Remove a terminated private plan",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress()

			coin, err := sdk.ParseCoinNormalized(string(args[0]))
			if err != nil {
				return err
			}
			msg := types.NewMsgBurnDenom(string(creator), coin)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func CmdChangeDenomAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "change-denom-admin",
		Args: cobra.ExactArgs(1),
		// Short: "Remove a terminated private plan",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// creator := clientCtx.GetFromAddress()

			// msg := types.NewMsgRemovePlan(creator, planId)
			msg := types.MsgChangeDenomAdmin{}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func CmdForceTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "force-transfer",
		Args: cobra.ExactArgs(1),
		// Short: "Remove a terminated private plan",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// creator := clientCtx.GetFromAddress()

			// msg := types.NewMsgRemovePlan(creator, planId)
			msg := types.MsgForceTransfer{}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func CmdSetDenomMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "set-denom-metadata",
		Args: cobra.ExactArgs(1),
		// Short: "Remove a terminated private plan",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// creator := clientCtx.GetFromAddress()

			// planId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("parse plan id: %w", err)
			}

			// msg := types.NewMsgSetDenomMetadata(creator, types.MsgSetDenomMetadata{})
			msg := types.MsgSetDenomMetadata{}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
