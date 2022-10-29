package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"ollo/x/ons/types"
)

var _ = strconv.Itoa(0)

func CmdBuyName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-name [name] [bidder-addr] [bid]",
		Short: "Broadcast message buy-name",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argBidderAddr := args[1]
			argBid := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyName(
				clientCtx.GetFromAddress().String(),
				argName,
				argBidderAddr,
				argBid,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
