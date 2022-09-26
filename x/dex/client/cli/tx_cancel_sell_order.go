package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"ollo/x/dex/types"
)

var _ = strconv.Itoa(0)

func CmdCancelSellOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-sell-order [port] [channel] [amount-denom] [price-denom] [order-id]",
		Short: "Cancel a sell order",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPort := args[0]
			argChannel := args[1]
			argAmountDenom := args[2]
			argPriceDenom := args[3]
			argOrderID, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelSellOrder(
				clientCtx.GetFromAddress().String(),
				argPort,
				argChannel,
				argAmountDenom,
				argPriceDenom,
				argOrderID,
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
