package cmd

import (
	"bufio"
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/client/keys"

	cosmoshd "github.com/cosmos/cosmos-sdk/crypto/hd"
	etherminthd "github.com/evmos/ethermint/crypto/hd"

	// ethermintclient "github.com/evmos/ethermint/client"
	clientkeys "github.com/evmos/ethermint/client/keys"
)

func KeyCommands(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keyring",
		Short: "Keyring management commands",
		Long:  "Keyring management commands",
	}
	addCmd := keys.AddKeyCommand()
	algoFlag := addCmd.Flag(flags.FlagKeyAlgorithm)
	algoFlag.DefValue = string(cosmoshd.Secp256k1Type)
	err := algoFlag.Value.Set(string(cosmoshd.Secp256k1Type))
	if err != nil {
		panic(err)
	}
	addCmd.RunE = runAddCmd

	cmd.AddCommand(
		keys.MnemonicKeyCommand(),
		keys.ExportKeyCommand(),
		importKeyCommand(),
		keys.AddKeyCommand(),
		keys.ShowKeysCmd(),
		keys.ListKeysCmd(),
		flags.LineBreak,
		keys.DeleteKeyCommand(),
		keys.ParseKeyStringCommand(),
		keys.MigrateCommand(),
		// ethermintclient.UnsafeExportEthKeyCommand(),
		// ethermintclient.UnsafeImportKeyCommand(),

		addCmd,
	)
	cmd.PersistentFlags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.PersistentFlags().String(flags.FlagKeyringDir, "", "The client Keyring directory; if omitted, the default 'home' directory will be used")
	cmd.PersistentFlags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	cmd.PersistentFlags().String(cli.OutputFlag, "text", "Output format (text|json)")
	return cmd
}
func importKeyCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "import <name> <keyfile>",
		Short: "Import private keys into the local keybase",
		Long:  "Import a ASCII armored private key into the local keybase.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			buf := bufio.NewReader(cmd.InOrStdin())
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			bz, err := ioutil.ReadFile(args[1])
			if err != nil {
				return err
			}

			passphrase, err := input.GetPassword("Enter passphrase to decrypt your key:", buf)
			if err != nil {
				return err
			}

			armor, err := getArmor(bz, passphrase)
			if err != nil {
				return err
			}
			return clientCtx.Keyring.ImportPrivKey(args[0], armor, passphrase)
		},
	}
}

func getArmor(privBytes []byte, passphrase string) (string, error) {
	if !json.Valid(privBytes) {
		return string(privBytes), nil
	}
	// return keystore.RecoveryAndExportPrivKeyArmor(privBytes, passphrase)
	return "", nil
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	clientCtx := client.GetClientContextFromCmd(cmd).WithKeyringOptions(etherminthd.EthSecp256k1Option())
	clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
	if err != nil {
		return err
	}
	buf := bufio.NewReader(clientCtx.Input)
	return clientkeys.RunAddCmd(clientCtx, cmd, args, buf)
}
