package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ollo-station/ollo/x/ons/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSellName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-name [name] [seller-addr] [offer]",
		Short: "Broadcast message sell-name",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argSellerAddr := args[1]
			argOffer := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSellName(
				clientCtx.GetFromAddress().String(),
				argName,
				argSellerAddr,
				argOffer,
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
