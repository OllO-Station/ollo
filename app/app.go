package app

import (
	"fmt"
	"time"

	// icagenesistypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/genesis/types"
	"io"
	"net/http"
	"ollo/docs"
	"ollo/x/farming"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"

	farmingkeeper "ollo/x/farming/keeper"
	farmingtypes "ollo/x/farming/types"
	grants "ollo/x/grants"
	grantskeeper "ollo/x/grants/keeper"
	grantstypes "ollo/x/grants/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	// "github.com/CosmWasm/wasmd/x/wasm"

	// solomachine "github.com/cosmos/ibc-go/v6/modules/light-clients/06-solomachine"
	// tmint "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint"

	ibcmock "github.com/cosmos/ibc-go/v6/testing/mock"

	// "github.com/cosmos/ibc-go/v6/modules/core/04-channel"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	epochingkeeper "github.com/cosmos/cosmos-sdk/x/epoching/keeper"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"

	// "github.com/cosmos/cosmos-sdk/x/mint"
	// mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	// minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	nftmodule "github.com/cosmos/cosmos-sdk/x/nft/module"

	// nftmodule "github.com/cosmos/cosmos-sdk/x/nft/client/cli"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
	ibcfee "github.com/cosmos/ibc-go/v6/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v6/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v6/modules/apps/29-fee/types"
	"github.com/cosmos/ibc-go/v6/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v6/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v6/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v6/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v6/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"

	// v6 "github.com/cosmos/ibc-go/v6/testing/simapp/upgrades/v6"

	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	claimmodule "ollo/x/claim"
	claimmodulekeeper "ollo/x/claim/keeper"
	claimmoduletypes "ollo/x/claim/types"
	liquiditymodule "ollo/x/liquidity"
	liquiditymodulekeeper "ollo/x/liquidity/keeper"
	liquiditymoduletypes "ollo/x/liquidity/types"
	loanmodule "ollo/x/loan"
	loanmodulekeeper "ollo/x/loan/keeper"
	loanmoduletypes "ollo/x/loan/types"
	marketmodule "ollo/x/market"
	marketmodulekeeper "ollo/x/market/keeper"
	marketmoduletypes "ollo/x/market/types"
	onsmodule "ollo/x/ons"
	onsmodulekeeper "ollo/x/ons/keeper"
	onsmoduletypes "ollo/x/ons/types"
	reservemodule "ollo/x/reserve"
	reservemodulekeeper "ollo/x/reserve/keeper"
	reservemoduletypes "ollo/x/reserve/types"

	mintmodule "ollo/x/mint"
	mintmodulekeeper "ollo/x/mint/keeper"
	mintmoduletypes "ollo/x/mint/types"

	tokenmodule "ollo/x/token"
	tokenmodulekeeper "ollo/x/token/keeper"
	tokenmoduletypes "ollo/x/token/types"

	// oraclemodule "ollo/x/oracle"
	// oraclemodulekeeper "ollo/x/oracle/keeper"
	// oraclemoduletypes "ollo/x/oracle/types"

	// this line is used by starport scaffolding # stargate/app/moduleImport

	appparams "ollo/app/params"
)

const (
	AccountAddressPrefix           = "ollo"
	Name                           = "ollo"
	EnableSpecificProposals        = ""
	ProposalsEnabled               = "true"
	AppBinary                      = "oxd"
	MockFeePort             string = ibcmock.ModuleName + ibcfeetypes.ModuleName
)

// func GetEnabledProposals() []wasm.ProposalType {
// 	if EnableSpecificProposals == "" {
// 		if ProposalsEnabled == "true" {
// 			return wasm.EnableAllProposals
// 		}
// 		return wasm.DisableAllProposals
// 	}
// 	chunks := strings.Split(EnableSpecificProposals, ",")
// 	proposals, err := wasm.ConvertToProposals(chunks)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return proposals
// }

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.LegacyProposalHandler,
		upgradeclient.LegacyCancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mintmodule.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		nftmodule.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		ica.AppModuleBasic{},
		vesting.AppModuleBasic{},
		liquiditymodule.AppModuleBasic{},
		onsmodule.AppModuleBasic{},
		marketmodule.AppModuleBasic{},
		claimmodule.AppModuleBasic{},
		reservemodule.AppModuleBasic{},
		loanmodule.AppModuleBasic{},
		ibcfee.AppModuleBasic{},
		grants.AppModuleBasic{},
		farming.AppModuleBasic{},
		tokenmodule.AppModuleBasic{},
		// wasm.AppModuleBasic{},
		// emissionsmodule.AppModuleBasic{},
		ibcmock.AppModuleBasic{},
		// oraclemodule.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName: nil,
		distrtypes.ModuleName:      nil,
		icatypes.ModuleName:        nil,
		ibcfeetypes.ModuleName:     nil,

		farmingtypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		grantstypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
		reservemoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		nft.ModuleName:                nil,
		// wasm.ModuleName:                 {authtypes.Burner},
		stakingtypes.BondedPoolName:     {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:  {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:             {authtypes.Burner},
		ibctransfertypes.ModuleName:     {authtypes.Minter, authtypes.Burner},
		liquiditymoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		onsmoduletypes.ModuleName:       {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		marketmoduletypes.ModuleName:    {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		claimmoduletypes.ModuleName:     {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		loanmoduletypes.ModuleName:      {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		// emissionsmoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		mintmoduletypes.ModuleName:  {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		tokenmoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},

		// oraclemoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		// this line is used by starport scaffolding # stargate/app/maccPerms
		ibcmock.ModuleName: nil,
	}
)

var (
	_ servertypes.Application = (*App)(nil)
	_ simapp.App              = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	AuthzKeeper      authzkeeper.Keeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	// MintKeeper          mintkeeper.Keeper
	MintKeeper          *mintmodulekeeper.Keeper
	DistrKeeper         distrkeeper.Keeper
	EpochingKeeper      epochingkeeper.Keeper
	GovKeeper           govkeeper.Keeper
	CrisisKeeper        crisiskeeper.Keeper
	UpgradeKeeper       upgradekeeper.Keeper
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ParamsKeeper        paramskeeper.Keeper
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper      evidencekeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	FeeGrantKeeper      feegrantkeeper.Keeper
	GroupKeeper         groupkeeper.Keeper
	NFTKeeper           nftkeeper.Keeper
	GrantsKeeper        grantskeeper.Keeper
	FarmingKeeper       farmingkeeper.Keeper
	TokenKeeper         tokenmodulekeeper.Keeper
	// WasmKeeper       wasm.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedIBCFeeKeeper        capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper          capabilitykeeper.ScopedKeeper
	ScopedFeeMockKeeper       capabilitykeeper.ScopedKeeper
	ScopedIBCMockKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAMockKeeper       capabilitykeeper.ScopedKeeper

	ICAAuthModule ibcmock.IBCModule
	FeeMockModule ibcmock.IBCModule

	LiquidityKeeper    liquiditymodulekeeper.Keeper
	ScopedOnsKeeper    capabilitykeeper.ScopedKeeper
	OnsKeeper          onsmodulekeeper.Keeper
	ScopedMarketKeeper capabilitykeeper.ScopedKeeper
	MarketKeeper       marketmodulekeeper.Keeper
	ScopedClaimKeeper  capabilitykeeper.ScopedKeeper
	ClaimKeeper        claimmodulekeeper.Keeper

	ReserveKeeper    reservemodulekeeper.Keeper
	ScopedLoanKeeper capabilitykeeper.ScopedKeeper
	LoanKeeper       loanmodulekeeper.Keeper

	// EmissionsKeeper emissionsmodulekeeper.Keeper

	// mintKeeper mintmodulekeeper.Keeper
	// ScopedOracleKeeper capabilitykeeper.ScopedKeeper
	// OracleKeeper       oraclemodulekeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// mm is the module manager
	mm *module.Manager

	// sm is the simulation manager
	sm           *module.SimulationManager
	configurator module.Configurator
}

// New returns a reference to an initialized blockchain app
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig appparams.EncodingConfig,
	appOpts servertypes.AppOptions,
	// enabledProposals []wasm.ProposalType,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(
		Name,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...,
	)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		authz.ModuleName,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		// minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		ibchost.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		ibctransfertypes.StoreKey,
		icahosttypes.StoreKey,
		capabilitytypes.StoreKey,
		group.StoreKey,
		ibcfeetypes.StoreKey,
		icacontrollertypes.StoreKey,
		liquiditymoduletypes.StoreKey,
		onsmoduletypes.StoreKey,
		marketmoduletypes.StoreKey,
		nftkeeper.StoreKey,
		claimmoduletypes.StoreKey,
		claimmoduletypes.MemStoreKey,
		reservemoduletypes.StoreKey,
		grantstypes.StoreKey,
		grantstypes.MemStoreKey,
		farmingtypes.StoreKey,
		loanmoduletypes.StoreKey,
		// tokenmoduletypes.StoreKey,
		string(
			epochingkeeper.ActionStoreKey(
				epochingkeeper.DefaultEpochNumber,
				epochingkeeper.DefaultEpochActionID,
			),
		),
		// emissionsmoduletypes.StoreKey,
		mintmoduletypes.StoreKey,
		// oraclemoduletypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(
		app.ParamsKeeper.Subspace(baseapp.Paramspace).
			WithKeyTable(paramstypes.ConsensusParamsKeyTable()),
	)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(
		icacontrollertypes.SubModuleName,
	)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	scopedIBCFeeKeeper := app.CapabilityKeeper.ScopeToModule(ibcfeetypes.ModuleName)
	// scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)
	// this line is used by starport scaffolding # stargate/app/scopedKeeper
	scopedIBCMockKeeper := app.CapabilityKeeper.ScopeToModule(ibcmock.ModuleName)
	scopedFeeMockKeeper := app.CapabilityKeeper.ScopeToModule(MockFeePort)
	scopedICAMockKeeper := app.CapabilityKeeper.ScopeToModule(
		ibcmock.ModuleName + icacontrollertypes.SubModuleName,
	)

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms, AccountAddressPrefix,
	)

	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authz.ModuleName],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	mintKeeper := mintmodulekeeper.NewKeeper(
		appCodec,
		keys[mintmoduletypes.StoreKey],
		app.GetSubspace(mintmoduletypes.ModuleName),
		app.StakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.DistrKeeper,
		authtypes.FeeCollectorName,
	)
	app.MintKeeper = &mintKeeper
	mintModule := mintmodule.NewAppModule(appCodec, *app.MintKeeper, app.AccountKeeper)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedModuleAccountAddrs(),
	)

	epochingKeeper := epochingkeeper.NewKeeper(
		appCodec,
		keys[string(epochingkeeper.ActionStoreKey(epochingkeeper.DefaultEpochNumber, epochingkeeper.DefaultEpochActionID))],
		time.Duration(1)*time.Second*60*60*24*14,
	)
	app.EpochingKeeper = epochingKeeper

	app.FarmingKeeper = farmingkeeper.NewKeeper(
		appCodec,
		keys[farmingtypes.StoreKey],
		app.GetSubspace(farmingtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.BlockedModuleAccountAddrs(),
	)
	grantsKeeper := grantskeeper.NewKeeper(appCodec,
		keys[grantstypes.StoreKey],
		keys[grantstypes.MemStoreKey],
		app.GetSubspace(grantstypes.ModuleName), app.AccountKeeper, app.BankKeeper, app.DistrKeeper)
	app.GrantsKeeper = grantsKeeper

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)

	// app.MintKeeper = mintkeeper.NewKeeper(
	// 	appCodec,
	// 	keys[minttypes.StoreKey],
	// 	app.GetSubspace(minttypes.ModuleName),
	// 	&stakingKeeper,
	// 	app.AccountKeeper,
	// 	app.BankKeeper,
	// 	authtypes.FeeCollectorName,
	// )

	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.GetSubspace(distrtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		authtypes.FeeCollectorName,
	)

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		keys[slashingtypes.StoreKey],
		&stakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)

	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName),
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)

	groupConfig := group.DefaultConfig()
	/*
		Example of setting group params:
		groupConfig.MaxMetadataLen = 1000
	*/
	app.GroupKeeper = groupkeeper.NewKeeper(
		keys[group.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
		groupConfig,
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)

	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	// ... other modules keepers

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibchost.StoreKey],
		app.GetSubspace(ibchost.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
	)

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)

	mockModule := ibcmock.NewAppModule(&app.IBCKeeper.PortKeeper)

	// The mock module is used for testing IBC
	mockIBCModule := ibcmock.NewIBCModule(
		&mockModule,
		ibcmock.NewIBCApp(ibcmock.ModuleName, scopedIBCMockKeeper),
	)

	icaControllerKeeper := icacontrollerkeeper.NewKeeper(
		appCodec,
		keys[icacontrollertypes.StoreKey],
		app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCFeeKeeper, // may be replaced with middleware such as ics29 fee
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper,
		app.MsgServiceRouter(),
	)
	app.ICAControllerKeeper = icaControllerKeeper
	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec, keys[icahosttypes.StoreKey], app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCFeeKeeper,
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, scopedICAHostKeeper, app.MsgServiceRouter(),
	)

	tokenKeeper := tokenmodulekeeper.NewKeeper(
		appCodec,
		keys[tokenmoduletypes.StoreKey],
		app.GetSubspace(tokenmoduletypes.ModuleName),
		app.BankKeeper,
		app.BlockedModuleAccountAddrs(),
		authtypes.FeeCollectorName,
	)
	app.TokenKeeper = tokenKeeper
	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper
	// wasmDir := filepath.Join(homePath, "wasm")
	// wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	// if err != nil {
	// panic(fmt.Sprintf("error while reading wasm config: %s", err))
	// }

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	// availableCapabilities := "iterator,staking,stargate,cosmwasm_1_1"
	// app.WasmKeeper = wasm.NewKeeper(
	// 	appCodec,
	// 	keys[wasm.StoreKey],
	// 	app.GetSubspace(wasm.ModuleName),
	// 	app.AccountKeeper,
	// 	app.BankKeeper,
	// 	app.StakingKeeper,
	// 	app.DistrKeeper,
	// 	app.IBCKeeper.ChannelKeeper,
	// 	&app.IBCKeeper.PortKeeper,
	// 	// scopedWasmKeeper,
	// 	app.TransferKeeper,
	// 	app.MsgServiceRouter(),
	// 	app.GRPCQueryRouter(),
	// wasmDir,
	// wasmConfig,
	// availableCapabilities,
	// wasmOpts...,
	// )
	app.NFTKeeper = nftkeeper.NewKeeper(
		keys[nftkeeper.StoreKey],
		appCodec,
		app.AccountKeeper,
		app.BankKeeper,
	)

	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))
	govConfig := govtypes.DefaultConfig()
	govKeeper := govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		govRouter,
		app.MsgServiceRouter(),
		govConfig,
	)
	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)
	app.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		appCodec, keys[ibcfeetypes.StoreKey],
		// app.GetSubspace(ibcfeetypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with IBC middleware
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper, app.AccountKeeper, app.BankKeeper,
	)

	// icaFeeStack := ibcfee.NewIBCMiddleware(icaControllerStack, app.IBCFeeKeeper)
	var transferStack ibcporttypes.IBCModule
	transferStack = transfer.NewIBCModule(app.TransferKeeper)
	transferStack = ibcfee.NewIBCMiddleware(transferStack, app.IBCFeeKeeper)

	var icaControllerStack ibcporttypes.IBCModule
	icaControllerStack = ibcmock.NewIBCModule(
		&mockModule,
		ibcmock.NewIBCApp("", scopedICAMockKeeper),
	)
	app.ICAAuthModule = icaControllerStack.(ibcmock.IBCModule)
	icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, app.ICAControllerKeeper)
	icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerStack, app.IBCFeeKeeper)

	var icaHostStack ibcporttypes.IBCModule
	icaHostStack = icahost.NewIBCModule(app.ICAHostKeeper)
	icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, app.IBCFeeKeeper)

	app.LiquidityKeeper = *liquiditymodulekeeper.NewKeeper(
		appCodec,
		keys[liquiditymoduletypes.StoreKey],
		keys[liquiditymoduletypes.MemStoreKey],
		app.GetSubspace(liquiditymoduletypes.ModuleName),

		app.AccountKeeper,
		app.BankKeeper,
	)
	liquidityModule := liquiditymodule.NewAppModule(
		appCodec,
		app.LiquidityKeeper,
		app.AccountKeeper,
		app.BankKeeper,
	)

	scopedOnsKeeper := app.CapabilityKeeper.ScopeToModule(onsmoduletypes.ModuleName)
	app.ScopedOnsKeeper = scopedOnsKeeper
	app.OnsKeeper = *onsmodulekeeper.NewKeeper(
		appCodec,
		keys[onsmoduletypes.StoreKey],
		keys[onsmoduletypes.MemStoreKey],
		app.GetSubspace(onsmoduletypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedOnsKeeper,
		app.BankKeeper,
		app.AccountKeeper,
		app.GroupKeeper,
	)
	onsModule := onsmodule.NewAppModule(appCodec, app.OnsKeeper, app.AccountKeeper, app.BankKeeper)

	onsIBCModule := onsmodule.NewIBCModule(app.OnsKeeper)
	scopedMarketKeeper := app.CapabilityKeeper.ScopeToModule(marketmoduletypes.ModuleName)
	app.ScopedMarketKeeper = scopedMarketKeeper
	app.MarketKeeper = *marketmodulekeeper.NewKeeper(
		appCodec,
		keys[marketmoduletypes.StoreKey],
		keys[marketmoduletypes.MemStoreKey],
		app.GetSubspace(marketmoduletypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedMarketKeeper,
		app.AccountKeeper,
		app.BankKeeper,
	)
	marketModule := marketmodule.NewAppModule(
		appCodec,
		app.MarketKeeper,
		app.AccountKeeper,
		app.BankKeeper,
	)

	marketIBCModule := marketmodule.NewIBCModule(app.MarketKeeper)
	scopedClaimKeeper := app.CapabilityKeeper.ScopeToModule(claimmoduletypes.ModuleName)
	app.ScopedClaimKeeper = scopedClaimKeeper
	claimKeeper := *claimmodulekeeper.NewKeeper(
		appCodec,
		keys[claimmoduletypes.StoreKey],
		keys[claimmoduletypes.MemStoreKey],
		app.GetSubspace(claimmoduletypes.ModuleName),
		app.AccountKeeper,
		app.DistrKeeper,
		app.BankKeeper,
		app.StakingKeeper,
	)
	app.ClaimKeeper = claimKeeper
	claimModule := claimmodule.NewAppModule(
		appCodec,
		app.ClaimKeeper,
		app.AccountKeeper,
		app.BankKeeper,
	)

	// claimIBCModule := claimmodule.NewIBCModule(app.ClaimKeeper)

	reserveKeeper := reservemodulekeeper.NewKeeper(
		// appCodec,
		keys[reservemoduletypes.StoreKey],
		// keys[reservemoduletypes.MemStoreKey],
		app.GetSubspace(reservemoduletypes.ModuleName),

		app.AccountKeeper,

		app.BankKeeper.WithMintCoinsRestriction(
			reservemoduletypes.NewTokenFactoryDenomMintCoinsRestriction(),
		),
		app.DistrKeeper,
	)
	app.ReserveKeeper = reserveKeeper

	reserveModule := reservemodule.NewAppModule(
		app.ReserveKeeper,
		app.AccountKeeper,
		app.BankKeeper,
	)

	scopedLoanKeeper := app.CapabilityKeeper.ScopeToModule(loanmoduletypes.ModuleName)
	app.ScopedLoanKeeper = scopedLoanKeeper
	app.LoanKeeper = *loanmodulekeeper.NewKeeper(
		appCodec,
		keys[loanmoduletypes.StoreKey],
		keys[loanmoduletypes.MemStoreKey],
		app.GetSubspace(loanmoduletypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedLoanKeeper,
		app.BankKeeper,
		app.AccountKeeper,
		app.StakingKeeper,
		app.LiquidityKeeper,
	)
	loanModule := loanmodule.NewAppModule(
		appCodec,
		app.LoanKeeper,
		app.AccountKeeper,
		app.BankKeeper,
	)

	loanIBCModule := loanmodule.NewIBCModule(app.LoanKeeper)

	// app.EmissionsKeeper = *emissionsmodulekeeper.NewKeeper(
	// 	appCodec,
	// 	keys[emissionsmoduletypes.StoreKey],
	// 	keys[emissionsmoduletypes.MemStoreKey],
	// 	app.GetSubspace(emissionsmoduletypes.ModuleName),

	// 	app.StakingKeeper,
	// 	app.BankKeeper,
	// 	app.AccountKeeper,
	// 	app.DistrKeeper,
	// )
	// emissionsModule := emissionsmodule.NewAppModule(appCodec, app.EmissionsKeeper, app.AccountKeeper, app.BankKeeper)

	// scopedOracleKeeper := app.CapabilityKeeper.ScopeToModule(oraclemoduletypes.ModuleName)
	// app.ScopedOracleKeeper = scopedOracleKeeper
	// app.OracleKeeper = *oraclemodulekeeper.NewKeeper(
	// 	appCodec,
	// 	keys[oraclemoduletypes.StoreKey],
	// 	keys[oraclemoduletypes.MemStoreKey],
	// 	app.GetSubspace(oraclemoduletypes.ModuleName),
	// 	app.IBCKeeper.ChannelKeeper,
	// 	&app.IBCKeeper.PortKeeper,
	// 	scopedOracleKeeper,
	// 	app.AccountKeeper,
	// 	app.BankKeeper,
	// )

	// oracleModule := oraclemodule.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper)

	// oracleIBCModule := oraclemodule.NewIBCModule(app.OracleKeeper)

	feeMockModule := ibcmock.NewIBCModule(
		&mockModule,
		ibcmock.NewIBCApp(MockFeePort, scopedFeeMockKeeper),
	)
	app.FeeMockModule = feeMockModule
	feeWithMockModule := ibcfee.NewIBCMiddleware(feeMockModule, app.IBCFeeKeeper)

	// Sealing prevents other modules from creating scoped sub-keepers
	app.CapabilityKeeper.Seal()

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	// if len(enabledProposals) != 0 {
	// govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.WasmKeeper, enabledProposals))
	// }
	// ibcRouter.AddRoute(icahosttypes.SubModuleName, icaHostIBCModule).
	ibcRouter.AddRoute(onsmoduletypes.ModuleName, onsIBCModule)
	ibcRouter.
		// the ICA Controller middleware needs to be explicitly added to the IBC Router because the
		// ICA controller module owns the port capability for ICA. The ICA authentication module
		// owns the channel capability.
		AddRoute(icacontrollertypes.SubModuleName, icaControllerStack).
		AddRoute(icahosttypes.SubModuleName, icaHostStack).
		AddRoute(ibcmock.ModuleName+icacontrollertypes.SubModuleName, icaControllerStack)
		// AddRoute(oraclemoduletypes.ModuleName, oracleIBCModule)
	// ibcRouter.AddRoute(icacontrollertypes.SubModuleName, icaControllerIBCModule)
	ibcRouter.AddRoute(marketmoduletypes.ModuleName, marketIBCModule)
	// ibcRouter.AddRoute(ibcfeetypes.ModuleName, icaControllerStack)
	// ibcRouter.AddRoute(claimmoduletypes.ModuleName, claimIBCModule)
	ibcRouter.AddRoute(loanmoduletypes.ModuleName, loanIBCModule)
	// ibcRouter.AddRoute(oraclemoduletypes.ModuleName, oracleIBCModule)
	ibcRouter.AddRoute(ibcmock.ModuleName, mockIBCModule)
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack)
	ibcRouter.AddRoute(MockFeePort, feeWithMockModule)

	// this line is used by starport scaffolding # ibc/app/router
	app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		authzmodule.NewAppModule(
			appCodec,
			app.AuthzKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(
			appCodec,
			app.AccountKeeper,
			app.BankKeeper,
			app.FeeGrantKeeper,
			app.interfaceRegistry,
		),
		groupmodule.NewAppModule(
			appCodec,
			app.GroupKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		// mintmodule.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, mintmoduletypes.DefaultmintCalculationFn),
		mintModule,
		slashing.NewAppModule(
			appCodec,
			app.SlashingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
		),
		distr.NewAppModule(
			appCodec,
			app.DistrKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
		),
		nftmodule.NewAppModule(
			appCodec,
			app.NFTKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		grants.NewAppModule(
			appCodec,
			app.GrantsKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.DistrKeeper,
		),
		farming.NewAppModule(appCodec, app.FarmingKeeper, app.AccountKeeper, app.BankKeeper),
		// wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		params.NewAppModule(app.ParamsKeeper),
		// tokenmodule.NewAppModule(appCodec, app.TokenKeeper, app.AccountKeeper, app.BankKeeper),
		transfer.NewAppModule(app.TransferKeeper),
		ibcfee.NewAppModule(app.IBCFeeKeeper),
		ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper),
		mockModule,
		liquidityModule,
		onsModule,
		marketModule,
		claimModule,
		reserveModule,

		loanModule,
		// emissionsModule,
		// oracleModule,
		// this line is used by starport scaffolding # stargate/app/appModule
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		// upgrades should be run first
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		mintmoduletypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		liquiditymoduletypes.ModuleName,
		onsmoduletypes.ModuleName,
		nft.ModuleName,
		marketmoduletypes.ModuleName,
		claimmoduletypes.ModuleName,
		reservemoduletypes.ModuleName,
		loanmoduletypes.ModuleName,
		grantstypes.ModuleName,
		farmingtypes.ModuleName,
		// tokenmoduletypes.ModuleName,
		// emissionsmoduletypes.ModuleName,
		// mintmoduletypes.ModuleName,
		// oraclemoduletypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibcmock.ModuleName,
		// this line is used by starport scaffolding # stargate/app/beginBlockers
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		mintmoduletypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		liquiditymoduletypes.ModuleName,
		onsmoduletypes.ModuleName,
		marketmoduletypes.ModuleName,
		claimmoduletypes.ModuleName,
		nft.ModuleName,
		reservemoduletypes.ModuleName,
		loanmoduletypes.ModuleName,
		grantstypes.ModuleName,
		farmingtypes.ModuleName,
		tokenmoduletypes.ModuleName,
		// emissionsmoduletypes.ModuleName,
		// mintmoduletypes.ModuleName,
		// oraclemoduletypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibcmock.ModuleName,
		// this line is used by starport scaffolding # stargate/app/endBlockers
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		mintmoduletypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		group.ModuleName,
		liquiditymoduletypes.ModuleName,
		onsmoduletypes.ModuleName,
		marketmoduletypes.ModuleName,
		claimmoduletypes.ModuleName,
		reservemoduletypes.ModuleName,
		loanmoduletypes.ModuleName,
		nft.ModuleName,
		grantstypes.ModuleName,
		farmingtypes.ModuleName,
		// tokenmoduletypes.ModuleName,
		// emissionsmoduletypes.ModuleName,
		// mintmoduletypes.ModuleName,
		// oraclemoduletypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibcmock.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,

		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	// Uncomment if you want to set a custom migration order here.
	// app.mm.SetOrderMigrations(custom order)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)

	app.configurator = module.NewConfigurator(
		app.appCodec,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
	)
	app.mm.RegisterServices(app.configurator)
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(
			appCodec,
			app.AccountKeeper,
			app.BankKeeper,
			app.FeeGrantKeeper,
			app.interfaceRegistry,
		),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		// mintmodule.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil),
		// mintModule,
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(
			appCodec,
			app.DistrKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
		),
		slashing.NewAppModule(
			appCodec,
			app.SlashingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
		),
		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		authzmodule.NewAppModule(
			appCodec,
			app.AuthzKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		ibc.NewAppModule(app.IBCKeeper),
		transfer.NewAppModule(app.TransferKeeper),
		ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper),
		nftmodule.NewAppModule(
			appCodec,
			app.NFTKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		groupmodule.NewAppModule(
			appCodec,
			app.GroupKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		tokenmodule.NewAppModule(appCodec, app.TokenKeeper, app.AccountKeeper, app.BankKeeper),
		// grants.NewAppModule(appCodec, app.GrantsKeeper, app.AccountKeeper, app.BankKeeper, app.DistrKeeper),
		// farming.NewAppModule(appCodec, app.FarmingKeeper, app.AccountKeeper, app.BankKeeper, ),
		// claimmodule.NewAppModule(appCodec, app.ClaimKeeper, app.AccountKeeper, app.BankKeeper),
		liquidityModule,
		onsModule,
		marketModule,
		// claimModule,

		loanModule,
		// emissionsModule,
		// mintModule,
		// oracleModule,
	)
	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
				FeegrantKeeper:  app.FeeGrantKeeper,
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCKeeper: app.IBCKeeper,
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %s", err))
	}

	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.ScopedICAControllerKeeper = scopedICAControllerKeeper
	app.ScopedICAHostKeeper = scopedICAHostKeeper
	app.ScopedIBCFeeKeeper = scopedIBCFeeKeeper
	app.ScopedIBCMockKeeper = scopedIBCMockKeeper
	app.ScopedICAMockKeeper = scopedICAMockKeeper
	app.ScopedFeeMockKeeper = scopedFeeMockKeeper
	// app.ScopedWasmKeeper = scopedWasmKeeper
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	// icaRawGenesisState := genesisState[icatypes.ModuleName]

	// var icaGenesisState icagenesistypes.GenesisState
	// if err := app.cdc.UnmarshalJSON(icaRawGenesisState, &icaGenesisState); err != nil {
	// 	panic(err)
	// }

	// icaGenesisState.HostGenesisState.Params.AllowMessages = []string{"*"}
	// genesisJson, err := app.cdc.MarshalJSON(icaGenesisState)
	// if err != nil {
	// 	panic(err)
	// }

	// genesisState[icatypes.ModuleName] = genesisJson
	// app.UpgradeKeeper.SetUpgradeHandler(
	// v6.UpgradeName,
	// v6.CreateUpgradeHandler(
	// app.mm,
	// app.configurator,
	// app.appCodec,
	// app.keys[capabilitytypes.ModuleName],
	// app.CapabilityKeeper,
	// authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	// ),
	// )

	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedModuleAccountAddrs returns all the app's blocked module account
// addresses.
func (app *App) BlockedModuleAccountAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns an app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns an InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register app's OpenAPI routes.
	apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.FS(docs.Docs)))
	apiSvr.Router.HandleFunc("/", docs.Handler(Name, "/static/openapi.yml"))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(
		app.BaseApp.GRPCQueryRouter(),
		clientCtx,
		app.BaseApp.Simulate,
		app.interfaceRegistry,
	)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(
	appCodec codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino,
	key, tkey storetypes.StoreKey,
) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(grantstypes.ModuleName)
	paramsKeeper.Subspace(farmingtypes.ModuleName)
	// paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(mintmoduletypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	// paramsKeeper.Subspace(wasm.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(icatypes.ModuleName)
	paramsKeeper.Subspace(liquiditymoduletypes.ModuleName)
	paramsKeeper.Subspace(onsmoduletypes.ModuleName)
	paramsKeeper.Subspace(marketmoduletypes.ModuleName)
	paramsKeeper.Subspace(claimmoduletypes.ModuleName)
	paramsKeeper.Subspace(reservemoduletypes.ModuleName)
	paramsKeeper.Subspace(loanmoduletypes.ModuleName)
	// paramsKeeper.Subspace(emissionsmoduletypes.ModuleName)
	paramsKeeper.Subspace(ibcfeetypes.ModuleName)
	paramsKeeper.Subspace(tokenmoduletypes.ModuleName)
	// paramsKeeper.Subspace(oraclemoduletypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}

func (app *App) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetStakingKeeper implements the TestingApp interface.
func (app *App) GetStakingKeeper() stakingkeeper.Keeper {
	return app.StakingKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *App) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *App) GetTxConfig() client.TxConfig {
	return MakeEncodingConfig().TxConfig
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}
