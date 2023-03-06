package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/ollo-station/ollo/app"
	"github.com/ollo-station/ollo/cmd/ollod/cmd"
)

// func config() {
// config := sdk.GetConfig()
// cfg.SetBech32Prefixes(config)
// cfg.SetBip44CoinType(config)
// cfg.RegisterDenoms()
// config.Seal()

// }
func main() {
	// config()
	rootCmd, _ := cmd.NewRootCmd()
	// fmt.Println("encodingConfig ", encodingConfig)
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
