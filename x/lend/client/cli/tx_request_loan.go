package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ollo-station/ollo/x/lend/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRequestLoan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-loan [amount] [collateral] [fee] [deadline]",
		Short: "Broadcast message request-loan",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argCollateral := args[1]
			argFee := args[2]
			argDeadline := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestLoan(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argCollateral,
				argFee,
				argDeadline,
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
