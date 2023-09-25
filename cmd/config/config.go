package config

import (
	"os"

	"github.com/ollo-station/ollo/app"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	// wasmtypes "github.com/ollo-station/ollo/x/wasm/types"
	tmcfg "github.com/tendermint/tendermint/config"
)

func InitSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	// config.SetAddressVerifier(wasmtypes.VerifyAddressLen())
	config.Seal()
}
func InitClientCtx() client.Context {
	encodingConfig := app.MakeEncodingConfig()
	return client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(app.DefaultNodeHome).
		WithViper("")
}

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func InitAppConfig() (string, interface{}) {
	// The following code snippet is just for reference.

	// WASMConfig defines configuration for the wasm module.
	type WASMConfig struct {
		// This is the maximum sdk gas (wasm and storage) that we allow for any x/wasm "smart" queries
		QueryGasLimit uint64 `mapstructure:"query_gas_limit"`

		// Address defines the gRPC-web server to listen on
		LruSize uint64 `mapstructure:"lru_size"`
	}

	type CustomAppConfig struct {
		serverconfig.Config

		WASM WASMConfig `mapstructure:"wasm"`
	}

	srvCfg := serverconfig.DefaultConfig()
	srvCfg.MinGasPrices = "0uollo"

	customAppConfig := CustomAppConfig{
		Config: *srvCfg,
		WASM: WASMConfig{
			LruSize:       1,
			QueryGasLimit: 300000,
		},
	}

	customAppTemplate := serverconfig.DefaultConfigTemplate + `
[wasm]
# This is the maximum sdk gas (wasm and storage) that we allow for any x/wasm "smart" queries
query_gas_limit = 300000
# This is the number of wasm vm instances we keep cached in memory for speed-up
# Warning: this is currently unstable and may lead to crashes, best to keep for 0 unless testing locally
lru_size = 0`

	return customAppTemplate, customAppConfig
}

func InitTendermintConfig() *tmcfg.Config {
	cfg := tmcfg.DefaultConfig()
	return cfg
}
