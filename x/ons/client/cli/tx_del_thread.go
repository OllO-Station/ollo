package cli

import (
	"strconv"

	"ollo/x/ons/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDelThread() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del-thread [name] [thread] [addr]",
		Short: "Broadcast message del-thread",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argThread := args[1]
			argAddr := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteThread(
				clientCtx.GetFromAddress().String(),
				argName,
				argThread,
				argAddr,
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
