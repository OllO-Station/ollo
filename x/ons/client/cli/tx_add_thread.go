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

func CmdAddThread() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-thread [name] [thread] [addr] [offer]",
		Short: "Broadcast message add-thread",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argThread := args[1]
			argAddr := args[2]
			argOffer := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddThread(
				clientCtx.GetFromAddress().String(),
				argName,
				argThread,
				argAddr,
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
