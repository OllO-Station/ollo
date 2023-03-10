package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	// "path/filepath"
	"regexp"
	"strings"

	// ethermintdebug "github.com/evmos/ethermint/client/debug"
	// ethermintclient "github.com/evmos/ethermint/client"
	// ethermintencoding "github.com/evmos/ethermint/encoding"
	// ethermintserver "github.com/evmos/ethermint/server"
	// ethermintservercfg "github.com/evmos/ethermint/server/config"
	// ethermintsrvflags "github.com/evmos/ethermint/server/flags"

	// "github.com/cosmos/cosmos-sdk/baseapp"
	// "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/version"
	cfg "github.com/ollo-station/ollo/cmd/config"
	"github.com/prometheus/client_golang/prometheus"

	// "github.com/cosmos/cosmos-sdk/snapshots"
	// snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	// "github.com/cosmos/cosmos-sdk/store"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/fatih/color"

	// "github.com/mint/k/ignite/services/network"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	// this line is used by starport scaffolding # root/moduleImport

	appparams "github.com/ollo-station/ollo/app/params"

	"github.com/ollo-station/ollo/app"

	// "github.com/ollo-station/ollo/testutil/network"
	"github.com/ollo-station/ollo/x/wasm"
	wasmkeeper "github.com/ollo-station/ollo/x/wasm/keeper"

	// wasmtypes "github.com/ollo-station/ollo/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	etherminthd "github.com/evmos/ethermint/crypto/hd"
)

// NewRootCmd creates a new root command for a Cosmos SDK application
func NewRootCmd() (*cobra.Command, appparams.EncodingConfig) {
	encodingConfig := app.MakeEncodingConfig()
	// initClientCtx := InitClientCtx()
	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithBroadcastMode(flags.BroadcastBlock).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome).
		WithKeyringOptions(etherminthd.EthSecp256k1Option()).
		WithViper(cfg.EnvPrefix)

	// fgMagenta := color.New(color.FgHiMagenta, color.Bold).SprintFunc()
	// fgBlue := color.New(color.FgHiBlue, color.Italic, color.Bold).SprintFunc()
	fgDesc := color.New(color.Italic, color.Faint).SprintFunc()
	fgBold := color.New(color.Bold).SprintFunc()

	rootCmd := &cobra.Command{
		Use: version.AppName,
		Short: fgBold(
			"ollo-local ",
		) + fgDesc(
			"The OLLO Station network node v0.0.2 | ",
		),
		Long: fgBold(
			"ollo-local ",
		) + fgDesc(
			"The OLLO Station network node v0.0.2 | ",
		),
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			useLedger, _ := cmd.Flags().GetBool(flags.FlagUseLedger)
			if useLedger {
				return errors.New("--ledger: ledger not currently supported")

			}
			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			customAppTemplate, customAppConfig := cfg.InitAppConfig()
			customTMConfig := cfg.InitTendermintConfig()

			return server.InterceptConfigsPreRunHandler(
				cmd, customAppTemplate, customAppConfig, customTMConfig,
			)
		},
	}

	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgHiGreen, color.Bold).SprintFunc())
	cobra.AddTemplateFunc("StyleParam", color.New(color.FgHiYellow).SprintFunc())
	cobra.AddTemplateFunc("StyleTx", color.New(color.FgHiCyan, color.Bold).SprintFunc())
	cobra.AddTemplateFunc("StyleSubcmd", color.New(color.FgHiBlue, color.Bold).SprintFunc())
	cobra.AddTemplateFunc("StyleFlags", color.New(color.FgBlue, color.Bold).SprintFunc())
	cobra.AddTemplateFunc("StyleError", color.New(color.FgHiRed, color.Bold).SprintFunc())
	usageTmpl := rootCmd.UsageTemplate()
	usageTmpl = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
		`Flags:`, `{{StyleHeading "Flags:"}}`,
		`[command]`, `{{StyleParam "[command]"}}`,
		`[flags]`, `{{StyleParam "[command]"}}`,
		`query`, `{{StyleSubcmd "query"}}`,
		`tx`, `{{StyleTx "tx"}}`,
		`Error`, `{{StyleError "Error"}}`,
		// The following one steps on "Global Flags:"
	).Replace(usageTmpl)

	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTmpl = re.ReplaceAllLiteralString(usageTmpl, `{{StyleHeading "Flags:"}}`)
	rootCmd.SetUsageTemplate(usageTmpl)
	initRootCmd(rootCmd, encodingConfig)
	overwriteFlagDefaults(rootCmd, map[string]string{
		flags.FlagChainID:        strings.ReplaceAll(app.Name, "-", ""),
		flags.FlagKeyringBackend: "test",
	})

	return rootCmd, encodingConfig
}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
	wasm.AddModuleInitFlags(startCmd)
}
func initRootCmd(
	rootCmd *cobra.Command,
	encodingConfig appparams.EncodingConfig,
) {
	// Set config
	cfg.InitSDKConfig()

	rootCmd.AddCommand(

		// ethermintclient.ValidateChainID(
		// 	genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		// ),
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(),
		genutilcli.GenTxCmd(
			app.ModuleBasics,
			encodingConfig.TxConfig,
			banktypes.GenesisBalancesIterator{},
			app.DefaultNodeHome,
		),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		AddGenesisAccountCmd(app.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		completionCmd,
		debug.Cmd(),
		config.Cmd(),
		// ethermintclient.KeyCommands(app.DefaultNodeHome),
		// ethermintserver.NewIndexTxCmd(),
		server.RosettaCommand(encodingConfig.InterfaceRegistry, encodingConfig.Marshaler),
		// tmcmd.ResetAllCmd,
		// tmcmd.LightCmd,
		// tmcmd.ReplayCmd,
		// nftcli.GetTxCmd(),
		// ExportBalancesCmd(),
		testnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		// this line is used by starport scaffolding # root/commands
	)

	a := appCreator{encodingConfig}

	// add server commands
	server.AddCommands(
		rootCmd,
		app.DefaultNodeHome,
		a.newApp,
		a.appExport,
		AddModuleInitFlags,
	)

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		rpc.BlockCommand(),
		queryCommand(),
		txCommand(),
		keys.Commands(app.DefaultNodeHome),
		// startWithTunnelingCommand(a, app.DefaultNodeHome),
	)
}

func AddModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
	wasm.AddModuleInitFlags(startCmd)
}

// queryCommand returns the sub-command to send queries to the app
func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		Long:                       "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// txCommand returns the sub-command to send transactions to the app
func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
		Long:  "Transactions subcommands",
		// Aliases:                    []string{"t"},
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// startWithTunnelingCommand returns a new start command with http tunneling
// enabled.
func startWithTunnelingCommand(appCreator appCreator, defaultNodeHome string) *cobra.Command {
	startCmd := server.StartCmd(appCreator.newApp, defaultNodeHome)
	startCmd.Use = "start-with-http-tunneling"
	startCmd.Short = "Run the full node with http tunneling"
	// Backup existing PreRunE, since we'll override it.
	startPreRunE := startCmd.PreRunE
	startCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		// var (
		// 	ctx       = cmd.Context()
		// 	clientCtx = client.GetClientContextFromCmd(cmd)
		// 	serverCtx = server.GetServerContextFromCmd(cmd)
		// )
		// network.StartProxyForTunneledPeers(ctx, clientCtx, serverCtx)
		if startPreRunE == nil {
			return nil
		}
		return startPreRunE(cmd, args)
	}
	return startCmd
}

func overwriteFlagDefaults(c *cobra.Command, defaults map[string]string) {
	set := func(s *pflag.FlagSet, key, val string) {
		if f := s.Lookup(key); f != nil {
			f.DefValue = val
			f.Value.Set(val)
		}
	}
	for key, val := range defaults {
		set(c.Flags(), key, val)
		set(c.PersistentFlags(), key, val)
	}
	for _, c := range c.Commands() {
		overwriteFlagDefaults(c, defaults)
	}
}

type appCreator struct {
	encodingConfig appparams.EncodingConfig
}

// newApp creates a new Cosmos SDK app
func (a appCreator) newApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	var cache sdk.MultiStorePersistentCache

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := dbm.NewDB("metadata", dbm.GoLevelDBBackend, snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	var wasmOpts []wasm.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	snapshotOptions := snapshottypes.NewSnapshotOptions(
		cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval)),
		cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent)),
	)

	return app.New(
		logger,
		db,
		traceStore,
		true,
		skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		a.encodingConfig,
		appOpts,
		wasmOpts,
		app.GetEnabledProposals(),
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),

		baseapp.SetSnapshot(snapshotStore, snapshotOptions),
		baseapp.SetIAVLCacheSize(
			cast.ToInt(appOpts.Get(server.FlagIAVLCacheSize)),
		), // 1
		baseapp.SetIAVLDisableFastNode(
			cast.ToBool(appOpts.Get(server.FlagDisableIAVLFastNode)),
		), // 1
	)
}

// appExport creates a new simapp (optionally at a given height)
func (a appCreator) appExport(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
) (servertypes.ExportedApp, error) {
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home not set")
	}

	loadLatest := height == -1
	var emptyWasmOpts []wasm.Option
	app := app.New(
		logger,
		db,
		traceStore,
		loadLatest,
		map[int64]bool{},
		homePath,
		uint(1),
		a.encodingConfig,
		appOpts,
		emptyWasmOpts,
		app.GetEnabledProposals(),
	)

	if height != -1 {
		if err := app.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return app.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
