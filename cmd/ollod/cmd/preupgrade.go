package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

func HandlePreUpgrade() error {
	return errors.New("not implemented")
}
func preUpgradeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pre-upgrade",
		Short: "Pre-upgrade command",
		Long:  "Pre-upgrade command to implement custom pre-upgrade handling",
		Run: func(cmd *cobra.Command, args []string) {

			err := HandlePreUpgrade()

			if err != nil {
				os.Exit(30)
			}

			os.Exit(0)

		},
	}

	return cmd
}
