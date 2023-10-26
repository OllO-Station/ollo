package app

import (
	"fmt"

	// "strings"
	"time"

	// icagenesistypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/genesis/types"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	// ibctesting "github.com/cosmos/ibc-go/v6/testing"
	// "github.com/evmos/ethermint/ethereum/eip712"
	srvflags "github.com/evmos/ethermint/server/flags"
	// ethermint "github.com/evmos/ethermint/types"
	"github.com/cosmos/cosmos-sdk/store/streaming"
	"github.com/evmos/ethermint/x/evm"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/evmos/ethermint/x/evm/vm/geth"
	"github.com/evmos/ethermint/x/feemarket"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
	"github.com/ollo-station/ollo/docs"
	epoch "github.com/ollo-station/ollo/x/epoch"

	// epochkeeper "github.com/ollo-station/ollo/x/epoch/keeper"
	// epochtypes "github.com/ollo-station/ollo/x/epoch/types"

	"github.com/ollo-station/ollo/x/farming"
	"github.com/ollo-station/ollo/x/wasm"
	wasmclient "github.com/ollo-station/ollo/x/wasm/client"

	// "github.com/cosmos/cosmos-sdk/client/docs/statik"
	// "github.com/cosmos/cosmos-sdk/snapshots"
	// "github.com/cosmos/cosmos-sdk/snapshots/types"

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"

	"github.com/cosmos/cosmos-sdk/baseapp"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	emissionsmodule "github.com/ollo-station/ollo/x/emissions"
	emissionskeeper "github.com/ollo-station/ollo/x/emissions/keeper"
	emissionstypes "github.com/ollo-station/ollo/x/emissions/types"

	automationmodule "github.com/ollo-station/ollo/x/automation"
	automationkeeper "github.com/ollo-station/ollo/x/automation/keeper"
	automationtypes "github.com/ollo-station/ollo/x/automation/types"

	hooks "github.com/ollo-station/ollo/x/hooks"
	hookskeeper "github.com/ollo-station/ollo/x/hooks/keeper"
	hookstypes "github.com/ollo-station/ollo/x/hooks/types"

	engine "github.com/ollo-station/ollo/x/engine"
	enginekeeper "github.com/ollo-station/ollo/x/engine/keeper"
	enginetypes "github.com/ollo-station/ollo/x/engine/types"

	vault "github.com/ollo-station/ollo/x/vault"
	vaultkeeper "github.com/ollo-station/ollo/x/vault/keeper"
	vaulttupes "github.com/ollo-station/ollo/x/vault/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	farmingkeeper "github.com/ollo-station/ollo/x/farming/keeper"
	farmingtypes "github.com/ollo-station/ollo/x/farming/types"
	grants "github.com/ollo-station/ollo/x/grants"
	grantskeeper "github.com/ollo-station/ollo/x/grants/keeper"
	grantstypes "github.com/ollo-station/ollo/x/grants/types"

	// "github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"

	// "github.com/cosmos/cosmos-sdk/x/staking/"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	// "github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	// consensus "github.com/cosmos/cosmos-sdk/x/consensus"
	// consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	inter_tx "github.com/cosmos/interchain-accounts/x/inter-tx"
	intertxkeepers "github.com/cosmos/interchain-accounts/x/inter-tx/keeper"
	intertxtypes "github.com/cosmos/interchain-accounts/x/inter-tx/types"

	wasmkeeper "github.com/ollo-station/ollo/x/wasm/keeper"

	// solomachine "github.com/cosmos/ibc-go/v6/modules/light-clients/06-solomachine"
	// tmint "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint"

	ibcmock "github.com/cosmos/ibc-go/v6/testing/mock"
	"github.com/cosmos/ibc-go/v6/testing/simapp"

	// "github.com/cosmos/ibc-go/v6/modules/core/04-channel"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
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

	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	// "github.com/cosmos/cosmos-sdk/x/nft"
	// nftnativekeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	// nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	// nftmodule "github.com/cosmos/cosmos-sdk/x/nft/module"

	nftmodule "github.com/ollo-station/ollo/x/nft"
	nftkeeper "github.com/ollo-station/ollo/x/nft/keeper"
	nfttypes "github.com/ollo-station/ollo/x/nft/types"

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
	icagenesistypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/genesis/types"
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

	// ibcexported "github.com/cosmos/ibc-go/v6/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"

	// v6 "github.com/cosmos/ibc-go/v6/testing/simapp/upgrades/v6"

	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	claimmodule "github.com/ollo-station/ollo/x/claim"
	claimmodulekeeper "github.com/ollo-station/ollo/x/claim/keeper"
	claimmoduletypes "github.com/ollo-station/ollo/x/claim/types"
	lendmodule "github.com/ollo-station/ollo/x/lend"
	lendmodulekeeper "github.com/ollo-station/ollo/x/lend/keeper"
	lendmoduletypes "github.com/ollo-station/ollo/x/lend/types"
	liquiditymodule "github.com/ollo-station/ollo/x/liquidity"
	liquiditymodulekeeper "github.com/ollo-station/ollo/x/liquidity/keeper"
	liquiditymoduletypes "github.com/ollo-station/ollo/x/liquidity/types"
	marketmodule "github.com/ollo-station/ollo/x/market"
	marketmodulekeeper "github.com/ollo-station/ollo/x/market/keeper"
	marketmoduletypes "github.com/ollo-station/ollo/x/market/types"
	onsmodule "github.com/ollo-station/ollo/x/ons"
	onsmodulekeeper "github.com/ollo-station/ollo/x/ons/keeper"
	onsmoduletypes "github.com/ollo-station/ollo/x/ons/types"
	reservemodule "github.com/ollo-station/ollo/x/reserve"
	reservemodulekeeper "github.com/ollo-station/ollo/x/reserve/keeper"
	reservemoduletypes "github.com/ollo-station/ollo/x/reserve/types"

	// vaultmodulekeeper "github.com/ollo-station/ollo/x/vault/keeper"
	// vaultmodule "github.com/ollo-station/ollo/x/vault"
	// vaultmoduletypes "github.com/ollo-station/ollo/x/vault/types"

	// feesmodulekeeper "github.com/ollo-station/ollo/x/fees/keeper"
	// feesmodule "github.com/ollo-station/ollo/x/fees"
	// feesmoduletypes "github.com/ollo-station/ollo/x/fees/types"

	// lockmodulekeeper "github.com/ollo-station/ollo/x/lock/keeper"
	// lockmodule "github.com/ollo-station/ollo/x/lock"
	// lockmoduletypes "github.com/ollo-station/ollo/x/lock/types"

	// ratelimitmodulekeeper "github.com/ollo-station/ollo/x/ratelimit/keeper"
	// ratelimitmodule "github.com/ollo-station/ollo/x/ratelimit"
	// ratelimitmoduletypes "github.com/ollo-station/ollo/x/ratelimit/types"
	// mintmodule "github.com/ollo-station/ollo/x/mint"
	// mintmodulekeeper "github.com/ollo-station/ollo/x/mint/keeper"
	// mintmoduletypes "github.com/ollo-station/ollo/x/mint/types"

	tokenmodule "github.com/ollo-station/ollo/x/token"
	tokenmodulekeeper "github.com/ollo-station/ollo/x/token/keeper"
	tokenmoduletypes "github.com/ollo-station/ollo/x/token/types"

	// ibctm "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	// ibctmtypes "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"

	// oraclemodule "github.com/ollo-station/ollo/x/oracle"
	// oraclemodulekeeper "github.com/ollo-station/ollo/x/oracle/keeper"
	// oraclemoduletypes "github.com/ollo-station/ollo/x/oracle/types"

	// this line is used by starport scaffolding # stargate/app/moduleImport

	"github.com/ollo-station/ollo/x/exchange"
	exchangekeeper "github.com/ollo-station/ollo/x/exchange/keeper"
	exchangetypes "github.com/ollo-station/ollo/x/exchange/types"

	appparams "github.com/ollo-station/ollo/app/params"
)

const (
	AccountAddressPrefix           = "ollo"
	appName                        = "OLLO Station"
	Bech32Prefix                   = "ollo"
	Name                           = "ollo"
	EnableSpecificProposals        = ""
	ProposalsEnabled               = "false"
	AppBinary                      = "ollod"
	MockFeePort             string = ibcmock.ModuleName + ibcfeetypes.ModuleName
)

func GetEnabledProposals() []wasm.ProposalType {
	if EnableSpecificProposals == "" {
		if ProposalsEnabled == "true" {
			return wasm.EnableAllProposals
		}
		return wasm.DisableAllProposals
	}
	chunks := strings.Split(EnableSpecificProposals, ",")
	proposals, err := wasm.ConvertToProposals(chunks)
	if err != nil {
		panic(err)
	}
	return proposals
}

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
	govProposalHandlers = append(govProposalHandlers,
		wasmclient.ProposalHandlers...,
	)

	return govProposalHandlers
}

var (

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	DefaultNodeHome      = os.ExpandEnv("$HOME/.ollo")
	Bech32PrefixAccAddr  = Bech32Prefix
	Bech32PrefixAccPub   = Bech32Prefix + sdk.PrefixPublic
	Bech32PrefixValAddr  = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	Bech32PrefixValPub   = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	Bech32PrefixConsAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	Bech32PrefixConsPub  = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
	ModuleBasics         = module.NewBasicManager(
		auth.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
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
		evm.AppModuleBasic{},
		feemarket.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		ica.AppModuleBasic{},
		vesting.AppModuleBasic{},
		epoch.AppModuleBasic{},
		liquiditymodule.AppModuleBasic{},
		onsmodule.AppModuleBasic{},
		automationmodule.AppModuleBasic{},
		vault.AppModuleBasic{},
		hooks.AppModuleBasic{},
		engine.AppModuleBasic{},
		emissionsmodule.AppModuleBasic{},
		// lockmodule.AppModuleBasic{},
		marketmodule.AppModuleBasic{},
		claimmodule.AppModuleBasic{},
		reservemodule.AppModuleBasic{},
		lendmodule.AppModuleBasic{},
		// incentive.AppModuleBasic{},
		grants.AppModuleBasic{},
		farming.AppModuleBasic{},
		tokenmodule.AppModuleBasic{},
		wasm.AppModuleBasic{},
		ibcfee.AppModuleBasic{},
		// consensus.AppModuleBasic{},
		inter_tx.AppModuleBasic{},
		emissionsmodule.AppModuleBasic{},
		ibcmock.AppModuleBasic{},
		exchange.AppModuleBasic{},
		// oraclemodule.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName: nil,
		distrtypes.ModuleName:      nil,
		icatypes.ModuleName:        nil,
		ibcfeetypes.ModuleName:     nil,
		exchangetypes.ModuleName:   nil,

		vaulttupes.ModuleName:      {authtypes.Minter, authtypes.Burner},
		emissionstypes.ModuleName:  {authtypes.Minter, authtypes.Burner},
		automationtypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		enginetypes.ModuleName:     {authtypes.Minter, authtypes.Burner},
		hookstypes.ModuleName:      {authtypes.Minter, authtypes.Burner},

		farmingtypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		grantstypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
		reservemoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		// nfttypes.ModuleName:           { authtypes.Minter,authtypes.Burner},
		nfttypes.ModuleName: nil, // { authtypes.Minter,authtypes.Burner},
		// epochtypes.ModuleName:           {authtypes.Minter, authtypes.Burner},
		wasm.ModuleName:                 {authtypes.Burner},
		stakingtypes.BondedPoolName:     {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:  {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:             {authtypes.Burner},
		ibctransfertypes.ModuleName:     {authtypes.Minter, authtypes.Burner},
		liquiditymoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		onsmoduletypes.ModuleName:       {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		marketmoduletypes.ModuleName:    {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		claimmoduletypes.ModuleName:     {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		lendmoduletypes.ModuleName:      {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		evmtypes.ModuleName: {
			authtypes.Minter,
			authtypes.Burner,
		}, // used for secure addition and subtraction of balance using module account

		// emissionsmoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		minttypes.ModuleName:        {authtypes.Minter},
		tokenmoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		// oraclemoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		// this line is used by starport scaffolding # stargate/app/maccPerms
		ibcmock.ModuleName: nil,
	}
	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName:       true,
		claimmoduletypes.ModuleName: true,
		// TODO: Add vaulttupes.Foundation, vaulttypes.Team for vested vaults
	}
)

var (
	_ servertypes.Application = (*App)(nil)
	_ simapp.App              = (*App)(nil)
	// _ ibctesting.TestingApp   = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
	// manually update the power reduction by replacing micro (u) -> atto (a) Canto
	// sdk.DefaultPowerReduction = ethermint.PowerReduction
	// modify fee market parameter defaults through global
	feemarkettypes.DefaultMinGasPrice = sdk.NewDec(20_000_000_000)
	feemarkettypes.DefaultMinGasMultiplier = sdk.NewDecWithPrec(5, 1)
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
	// tpsCounter *tpsCounter
	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	AuthzKeeper      authzkeeper.Keeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	// MintKeeper          mintmodulekeeper.Keeper
	DistrKeeper    distrkeeper.Keeper
	EpochingKeeper epochingkeeper.Keeper
	// EpochKeeper         *epochkeeper.Keeper
	ExchangeKeeper      exchangekeeper.Keeper
	GovKeeper           govkeeper.Keeper
	CrisisKeeper        crisiskeeper.Keeper
	UpgradeKeeper       upgradekeeper.Keeper
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ParamsKeeper        paramskeeper.Keeper
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	InterTxKeeper       intertxkeepers.Keeper
	EvidenceKeeper      evidencekeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	FeeGrantKeeper      feegrantkeeper.Keeper

	HooksKeeper      hookskeeper.Keeper
	AutomationKeeper automationkeeper.Keeper
	EngineKeeper     enginekeeper.Keeper
	VaultKeeper      vaultkeeper.Keeper
	EmissionsKeeper  emissionskeeper.Keeper

	GroupKeeper   groupkeeper.Keeper
	NFTKeeper     nftkeeper.Keeper
	GrantsKeeper  grantskeeper.Keeper
	FarmingKeeper farmingkeeper.Keeper
	TokenKeeper   tokenmodulekeeper.Keeper
	WasmKeeper    wasm.Keeper

	// Ethermint keepers
	EvmKeeper       *evmkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper

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
	ScopedInterTxKeeper       capabilitykeeper.ScopedKeeper
	ScopedEngineKeeper        capabilitykeeper.ScopedKeeper
	ScopedAutomationKeeper    capabilitykeeper.ScopedKeeper
	ScopedHooksKeeper         capabilitykeeper.ScopedKeeper

	// ICAAuthModule ibcmock.IBCModule
	FeeMockModule ibcmock.IBCModule

	LiquidityKeeper    liquiditymodulekeeper.Keeper
	ScopedOnsKeeper    capabilitykeeper.ScopedKeeper
	OnsKeeper          onsmodulekeeper.Keeper
	ScopedMarketKeeper capabilitykeeper.ScopedKeeper
	MarketKeeper       marketmodulekeeper.Keeper
	ScopedClaimKeeper  capabilitykeeper.ScopedKeeper
	ClaimKeeper        claimmodulekeeper.Keeper
	ReserveKeeper      reservemodulekeeper.Keeper
	LendKeeper         lendmodulekeeper.Keeper
	// VaultKeeper   vaultmodulekeeper.Keeper
	// RatelimitKeeper ratelimitmodulekeeper.Keeper
	// LockKeeper	lockmodulekeeper.Keeper
	// IncentiveKeeper incentivekeeper.Keeper
	// FeesKeeper   feesmodulekeeper.Keeper

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
	wasmOpts []wasmkeeper.Option,
	enabledProposals []wasm.ProposalType,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Marshaler

	cdc := encodingConfig.Amino

	interfaceRegistry := encodingConfig.InterfaceRegistry

	// eip712.SetEncodingConfig(encodingConfig)
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
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		upgradetypes.StoreKey,
		// consensustypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		ibctransfertypes.StoreKey,
		// nftnativekeeper.StoreKey,
		icahosttypes.StoreKey,
		ibchost.StoreKey,
		capabilitytypes.StoreKey,
		group.StoreKey,
		ibcfeetypes.StoreKey,
		icacontrollertypes.StoreKey,
		liquiditymoduletypes.StoreKey,
		onsmoduletypes.StoreKey,
		marketmoduletypes.StoreKey,
		nfttypes.StoreKey,
		claimmoduletypes.MemStoreKey,
		claimmoduletypes.StoreKey,
		reservemoduletypes.StoreKey,
		grantstypes.StoreKey,
		grantstypes.MemStoreKey,
		farmingtypes.StoreKey,
		lendmoduletypes.StoreKey,
		intertxtypes.StoreKey,
		emissionstypes.StoreKey,
		enginetypes.StoreKey,
		automationtypes.StoreKey,
		hookstypes.StoreKey,
		vaulttupes.StoreKey,
		wasm.StoreKey,
		exchangetypes.StoreKey,
		// tokenmoduletypes.StoreKey,
		// ethermint keys
		evmtypes.StoreKey, feemarkettypes.StoreKey,
		// epochtypes.StoreKey,
		string(
			epochingkeeper.ActionStoreKey(
				epochingkeeper.DefaultEpochNumber,
				epochingkeeper.DefaultEpochActionID,
			),
		),
		// emissionsmoduletypes.StoreKey,
		minttypes.StoreKey,
		// oraclemoduletypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(
		paramstypes.TStoreKey,
		evmtypes.TransientKey,
		feemarkettypes.TransientKey,
	)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
	if _, _, err := streaming.LoadStreamingServices(bApp, appOpts, appCodec, keys); err != nil {
		fmt.Printf("failed to load state streaming: %s", err)
		os.Exit(1)
	}
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
	scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)
	// this line is used by starport scaffolding # stargate/app/scopedKeeper
	scopedInterTxKeeper := app.CapabilityKeeper.ScopeToModule(intertxtypes.ModuleName)
	scopedIBCMockKeeper := app.CapabilityKeeper.ScopeToModule(ibcmock.ModuleName)
	scopedFeeMockKeeper := app.CapabilityKeeper.ScopeToModule(MockFeePort)
	scopedClaimKeeper := app.CapabilityKeeper.ScopeToModule(claimmoduletypes.ModuleName)
	scopedOnsKeeper := app.CapabilityKeeper.ScopeToModule(onsmoduletypes.ModuleName)
	scopedMarketKeeper := app.CapabilityKeeper.ScopeToModule(marketmoduletypes.ModuleName)
	scopedEngineKeeper := app.CapabilityKeeper.ScopeToModule(enginetypes.ModuleName)
	scopedHooksKeeper := app.CapabilityKeeper.ScopeToModule(hookstypes.ModuleName)
	scopedAutomationKeeper := app.CapabilityKeeper.ScopeToModule(automationtypes.ModuleName)
	scopedICAMockKeeper := app.CapabilityKeeper.ScopeToModule(
		ibcmock.ModuleName + icacontrollertypes.SubModuleName,
	)
	// app.CapabilityKeeper.Seal()

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		// ethermint.ProtoAccount,
		authtypes.ProtoBaseAccount,
		maccPerms, AccountAddressPrefix,
	)
	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authz.ModuleName],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	// Create Ethermint keepers
	tracer := cast.ToString(appOpts.Get(srvflags.EVMTracer))
	feeMarketSs := app.GetSubspace(feemarkettypes.ModuleName)
	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec, authtypes.NewModuleAddress(govtypes.ModuleName),
		keys[feemarkettypes.StoreKey], tkeys[feemarkettypes.TransientKey], feeMarketSs,
	)
	// ethAddr := sdk.MustAccAddressFromBech32("ollo1phrdcmje043ydeqy750czh9fdk5h2ue7a5ref")
	evmSs := app.GetSubspace(evmtypes.ModuleName)
	app.EvmKeeper = evmkeeper.NewKeeper(
		appCodec,
		keys[evmtypes.StoreKey],
		tkeys[evmtypes.TransientKey],
		authtypes.NewModuleAddress(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.FeeMarketKeeper,
		nil,
		geth.NewEVM,
		tracer,
		evmSs,
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedModuleAccountAddrs(),
	)
	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)
	mintKeeper := mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		app.StakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)

	app.MintKeeper = mintKeeper
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.GetSubspace(distrtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&app.StakingKeeper,
		authtypes.FeeCollectorName,
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		keys[slashingtypes.StoreKey],
		&app.StakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)
	app.ExchangeKeeper = exchangekeeper.NewKeeper(
		appCodec,
		keys[exchangetypes.StoreKey],
		app.GetSubspace(exchangetypes.ModuleName),
		app.BankKeeper,
	)
	// app.EpochKeeper = epochkeeper.NewKeeper(
	// 	keys[epochtypes.StoreKey],
	// )
	// epochModule := epoch.NewAppModule(*app.EpochKeeper)

	// mintKeeper := mintmodulekeeper.NewKeeper(
	// 	appCodec,
	// 	keys[mintmoduletypes.StoreKey],
	// 	app.GetSubspace(mintmoduletypes.ModuleName),
	// 	app.StakingKeeper,
	// 	app.AccountKeeper,
	// 	app.BankKeeper,
	// 	app.DistrKeeper,
	// 	authtypes.FeeCollectorName,
	// )
	// app.MintKeeper = mintKeeper
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

	// app.MintKeeper = mintkeeper.NewKeeper(
	// 	appCodec,
	// 	keys[minttypes.StoreKey],
	// 	app.GetSubspace(minttypes.ModuleName),
	// 	&stakingKeeper,
	// 	app.AccountKeeper,
	// 	app.BankKeeper,
	// 	authtypes.FeeCollectorName,
	// )

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

	transferModule := transfer.NewAppModule(app.TransferKeeper)
	transferIBCModule := transfer.NewIBCModule(app.TransferKeeper)
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

	icaModule := ica.NewAppModule(&icaControllerKeeper, &app.ICAHostKeeper)
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

	vaultKeeper := vaultkeeper.NewKeeper(
		appCodec,
		keys[vaulttupes.StoreKey],
		memKeys[vaulttupes.MemStoreKey],
		app.GetSubspace(vaulttupes.ModuleName),
		app.AccountKeeper,
		app.EpochingKeeper,
		app.GroupKeeper,
		app.DistrKeeper,
		app.BankKeeper,
	)
	app.VaultKeeper = *vaultKeeper
	vaultModule := vault.NewAppModule(appCodec, app.VaultKeeper, app.AccountKeeper, app.BankKeeper)

	engineKeeper := enginekeeper.NewKeeper(
		appCodec,
		keys[enginetypes.StoreKey],
		memKeys[enginetypes.MemStoreKey],
		app.GetSubspace(enginetypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedEngineKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.LiquidityKeeper,
		app.StakingKeeper,
		app.AuthzKeeper,
		app.FeeGrantKeeper,
		app.LendKeeper,
		app.NFTKeeper,
		app.MarketKeeper,
		app.DistrKeeper,
		app.TokenKeeper,
		app.MintKeeper,
	)
	app.EngineKeeper = *engineKeeper
	engineModule := engine.NewAppModule(appCodec, app.EngineKeeper, app.AccountKeeper, app.BankKeeper)

	emissionsKeeper := emissionskeeper.NewKeeper(
		appCodec,
		keys[emissionstypes.StoreKey],
		memKeys[emissionstypes.MemStoreKey],
		app.GetSubspace(emissionstypes.ModuleName),
		app.BankKeeper,
		app.DistrKeeper,
		app.AccountKeeper,
		app.StakingKeeper,
		app.EpochingKeeper,
		app.MintKeeper,
		app.GovKeeper, app.LiquidityKeeper, app.LendKeeper,
	)
	app.EmissionsKeeper = *emissionsKeeper
	emissionsModule := emissionsmodule.NewAppModule(appCodec, app.EmissionsKeeper, app.AccountKeeper, app.BankKeeper)

	automationKeeper := automationkeeper.NewKeeper(
		appCodec,
		keys[automationtypes.StoreKey],
		memKeys[automationtypes.MemStoreKey],
		app.GetSubspace(automationtypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedAutomationKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.LiquidityKeeper,
		app.StakingKeeper,
		app.AuthzKeeper,
		app.FeeGrantKeeper,
		app.LendKeeper,
		app.NFTKeeper,
		app.MarketKeeper,
		app.DistrKeeper,
		app.TokenKeeper,
		app.MintKeeper,
	)
	app.AutomationKeeper = *automationKeeper
	automationModule := automationmodule.NewAppModule(appCodec, app.AutomationKeeper, app.AccountKeeper, app.BankKeeper)

	hooksKeeper := hookskeeper.NewKeeper(
		appCodec,
		keys[hookstypes.StoreKey],
		memKeys[hookstypes.MemStoreKey],
		app.GetSubspace(hookstypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedHooksKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.LiquidityKeeper,
		app.StakingKeeper,
		app.AuthzKeeper,
		app.FeeGrantKeeper,
		app.LendKeeper,
		app.NFTKeeper,
		app.MarketKeeper,
		app.DistrKeeper,
		app.TokenKeeper,
		app.MintKeeper,
	)
	app.HooksKeeper = *hooksKeeper
	hooksModule := hooks.NewAppModule(appCodec, app.HooksKeeper, app.AccountKeeper, app.BankKeeper)
	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	evidenceRouter := evidencetypes.NewRouter()
	evidenceKeeper.SetRouter(evidenceRouter)
	app.EvidenceKeeper = *evidenceKeeper
	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	availableCapabilities := "iterator,staking,stargate,cosmwasm_1_1,cosmwasm_1_2"
	app.WasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		availableCapabilities,
		wasmOpts...,
	)

	nftKeeper := nftkeeper.NewKeeper(
		appCodec,
		keys[nfttypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		// nftKeeper,
	)
	app.NFTKeeper = nftKeeper

	app.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		appCodec, keys[ibcfeetypes.StoreKey],
		// app.GetSubspace(ibcfeetypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with IBC middleware
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper, app.AccountKeeper, app.BankKeeper,
	)

	// icaFeeStack := ibcfee.NewIBCMiddleware(icaControllerStack, app.IBCFeeKeeper)
	var transferStack ibcporttypes.IBCModule
	transferStack = ibcfee.NewIBCMiddleware(transferIBCModule, app.IBCFeeKeeper)

	var icaControllerStack ibcporttypes.IBCModule
	// icaControllerStack = ibcmock.NewIBCModule(
	// 	&mockModule,
	// 	ibcmock.NewIBCApp("", scopedICAMockKeeper),
	// )
	// app.ICAAuthModule = icaControllerStack.(ibcmock.IBCModule)
	// icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, app.ICAControllerKeeper)
	// icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerStack, app.IBCFeeKeeper)
	icaControllerStack = inter_tx.NewIBCModule(app.InterTxKeeper)
	icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, app.ICAControllerKeeper)
	icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerStack, app.IBCFeeKeeper)

	var icaHostStack ibcporttypes.IBCModule
	icaHostStack = icahost.NewIBCModule(app.ICAHostKeeper)
	icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, app.IBCFeeKeeper)

	var wasmStack ibcporttypes.IBCModule
	wasmStack = wasm.NewIBCHandler(app.WasmKeeper, app.IBCKeeper.ChannelKeeper, app.IBCFeeKeeper)
	wasmStack = ibcfee.NewIBCMiddleware(wasmStack, app.IBCFeeKeeper)

	app.InterTxKeeper = intertxkeepers.NewKeeper(
		appCodec,
		keys[intertxtypes.StoreKey],
		app.ICAControllerKeeper,
		scopedInterTxKeeper,
	)
	interTxModule := inter_tx.NewAppModule(appCodec, app.InterTxKeeper)
	// interTxIBCModule := inter_tx.NewIBCModule(app.InterTxKeeper)

	liquidityKeeper := liquiditymodulekeeper.NewKeeper(
		appCodec,
		keys[liquiditymoduletypes.StoreKey],
		app.GetSubspace(liquiditymoduletypes.ModuleName),
		app.AccountKeeper, app.BankKeeper,
	)
	app.LiquidityKeeper = liquidityKeeper
	liquidityModule := liquiditymodule.NewAppModule(
		appCodec, app.LiquidityKeeper, app.AccountKeeper, app.BankKeeper,
	)

	app.ScopedOnsKeeper = scopedOnsKeeper
	app.OnsKeeper = *onsmodulekeeper.NewKeeper(
		appCodec,
		keys[onsmoduletypes.StoreKey],
		keys[onsmoduletypes.MemStoreKey],
		app.GetSubspace(onsmoduletypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedOnsKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.GroupKeeper,
		app.GovKeeper,
		app.EpochingKeeper,
		app.NFTKeeper,
		app.TokenKeeper,
		app.StakingKeeper,
		app.AuthzKeeper,
		app.FeeGrantKeeper,
	)
	onsModule := onsmodule.NewAppModule(appCodec, app.OnsKeeper, app.AccountKeeper, app.BankKeeper)

	onsIBCModule := onsmodule.NewIBCModule(app.OnsKeeper)
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
		app.AuthzKeeper,
		app.BankKeeper,
		app.FeeGrantKeeper,
		app.GroupKeeper,
		app.NFTKeeper,
		app.OnsKeeper,
		app.DistrKeeper,
	)
	marketModule := marketmodule.NewAppModule(
		appCodec,
		app.MarketKeeper,
		app.AccountKeeper,
		app.BankKeeper,
	)

	marketIBCModule := marketmodule.NewIBCModule(app.MarketKeeper)
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

	app.LendKeeper = *lendmodulekeeper.NewKeeper(
		appCodec,
		keys[lendmoduletypes.StoreKey],
		keys[lendmoduletypes.MemStoreKey],
		app.GetSubspace(lendmoduletypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.StakingKeeper,
		app.LiquidityKeeper,
	)
	lendModule := lendmodule.NewAppModule(
		appCodec,
		app.LendKeeper,
		app.AccountKeeper,
		app.BankKeeper,
	)

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
	// app.CapabilityKeeper.Seal()

	// Create static IBC router, add transfer route, then set and seal it

	govRouter := govv1beta1.NewRouter()
	if len(enabledProposals) != 0 {
		govRouter.AddRoute(
			wasm.RouterKey,
			wasm.NewWasmProposalHandler(app.WasmKeeper, enabledProposals),
		)
	}
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
		&app.StakingKeeper,
		govRouter,
		app.MsgServiceRouter(),
		govConfig,
	)

	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)

	app.EvmKeeper = app.EvmKeeper.SetHooks(
		evmkeeper.NewMultiEvmHooks(
		// app.Erc20Keeper.Hooks(),
		// app.FeesKeeper.Hooks(),
		// app.CSRKeeper.Hooks(),
		),
	)
	app.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			app.DistrKeeper.Hooks(),
			app.SlashingKeeper.Hooks(),
			// app.ClaimKeeper.NewGoalVoteHooks
		),
	)
	app.CapabilityKeeper.Seal()
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.
		AddRoute(onsmoduletypes.ModuleName, onsIBCModule).
		// AddRoute(ibcmock.ModuleName, feeMockModule).
		AddRoute(intertxtypes.ModuleName, icaControllerStack).
		AddRoute(icacontrollertypes.SubModuleName, icaControllerStack).
		AddRoute(icahosttypes.SubModuleName, icaHostStack).
		AddRoute(ibcmock.ModuleName+icacontrollertypes.SubModuleName, icaControllerStack).
		AddRoute(marketmoduletypes.ModuleName, marketIBCModule).
		AddRoute(ibcmock.ModuleName, mockIBCModule).
		AddRoute(ibctransfertypes.ModuleName, transferStack).
		AddRoute(MockFeePort, feeWithMockModule).
		AddRoute(wasm.ModuleName, wasmStack)
		// the ICA Controller middleware needs to be explicitly added to the IBC Router because the
		// ICA controller module owns the port capability for ICA. The ICA authentication module
		// owns the channel capability.
		// AddRoute(oraclemoduletypes.ModuleName, oracleIBCModule)
	// ibcRouter.AddRoute(icacontrollertypes.SubModuleName, icaControllerIBCModule)
	// ibcRouter.AddRoute(ibcfeetypes.ModuleName, icaControllerStack)
	// ibcRouter.AddRoute(claimmoduletypes.ModuleName, claimIBCModule)
	// ibcRouter.AddRoute(oraclemoduletypes.ModuleName, oracleIBCModule)

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
			app.AccountKeeper, app.StakingKeeper,
			app.BaseApp.DeliverTx,
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
		mint.NewAppModule(
			appCodec,
			app.MintKeeper,
			app.AccountKeeper,
			minttypes.DefaultInflationCalculationFn,
		),
		liquidityModule,
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
		// epochModule,
		exchange.NewAppModule(appCodec, app.ExchangeKeeper, app.AccountKeeper, app.BankKeeper),
		nftmodule.NewAppModule(
			appCodec,
			nftKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			// app.interfaceRegistry,
		),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		wasm.NewAppModule(
			appCodec,
			&app.WasmKeeper,
			app.StakingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
		),
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
		params.NewAppModule(app.ParamsKeeper),
		// tokenmodule.NewAppModule(appCodec, app.TokenKeeper, app.AccountKeeper, app.BankKeeper),
		transferModule,
		ibcfee.NewAppModule(app.IBCFeeKeeper),
		emissionsModule,
		vaultModule,
		engineModule,
		automationModule,
		hooksModule,
		icaModule,
		mockModule,
		onsModule,
		marketModule,
		claimModule,
		reserveModule,
		interTxModule,
		lendModule,
		// epochModule,
		// Ethermint app modules
		// evm.NewAppModule(
		// 	app.EvmKeeper,
		// 	app.AccountKeeper,
		// 	app.ParamsKeeper.Subspace(evmtypes.ModuleName),
		// ),
		// feemarket.NewAppModule(
		// 	app.FeeMarketKeeper,
		// 	app.ParamsKeeper.Subspace(feemarkettypes.ModuleName),
		// ),
		// feemarket.NewAppModule(app.FeeMarketKeeper, feeMarketSs),
		// evm.NewAppModule(app.EvmKeeper, app.AccountKeeper, evmSs),
		emissionsModule,
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
		// feemarkettypes.ModuleName,
		// evmtypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		ibctransfertypes.ModuleName,
		// ibchost.ModuleName,
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
		marketmoduletypes.ModuleName,
		claimmoduletypes.ModuleName,
		reservemoduletypes.ModuleName,
		lendmoduletypes.ModuleName,
		grantstypes.ModuleName,
		farmingtypes.ModuleName,
		nfttypes.ModuleName,
		intertxtypes.ModuleName,
		wasm.ModuleName,
		// tokenmoduletypes.ModuleName,
		emissionstypes.ModuleName,
		vaulttupes.ModuleName,
		enginetypes.ModuleName,
		automationtypes.ModuleName,
		hookstypes.ModuleName,
		// mintmoduletypes.ModuleName,
		// oraclemoduletypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibcmock.ModuleName,
		exchangetypes.ModuleName,
		// epochtypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/beginBlockers
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		stakingtypes.ModuleName,
		// evmtypes.ModuleName,
		// feemarkettypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
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
		reservemoduletypes.ModuleName,
		lendmoduletypes.ModuleName,
		grantstypes.ModuleName,
		farmingtypes.ModuleName,
		tokenmoduletypes.ModuleName,
		nfttypes.ModuleName,
		intertxtypes.ModuleName,
		wasm.ModuleName,
		// epochtypes.ModuleName,
		emissionstypes.ModuleName,
		vaulttupes.ModuleName,
		enginetypes.ModuleName,
		automationtypes.ModuleName,
		hookstypes.ModuleName,
		// mintmoduletypes.ModuleName,
		// oraclemoduletypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibcmock.ModuleName,
		exchangetypes.ModuleName,
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
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		ibctransfertypes.ModuleName,
		// ibchost.ModuleName,
		ibchost.ModuleName,
		exchangetypes.ModuleName,
		// evmtypes.ModuleName,
		// feemarkettypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		group.ModuleName,
		liquiditymoduletypes.ModuleName,
		onsmoduletypes.ModuleName,
		marketmoduletypes.ModuleName,
		claimmoduletypes.ModuleName,
		reservemoduletypes.ModuleName,
		lendmoduletypes.ModuleName,
		grantstypes.ModuleName,
		farmingtypes.ModuleName,
		// tokenmoduletypes.ModuleName,
		emissionstypes.ModuleName,
		vaulttupes.ModuleName,
		enginetypes.ModuleName,
		automationtypes.ModuleName,
		hookstypes.ModuleName,
		// mintmoduletypes.ModuleName,
		// oraclemoduletypes.ModuleName,
		intertxtypes.ModuleName,
		wasm.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibcmock.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		nfttypes.ModuleName,
		// epochtypes.ModuleName,

		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	// Uncomment if you want to set a custom migration order here.
	// app.mm.SetOrderMigrations(...)

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
		authzmodule.NewAppModule(
			appCodec,
			app.AuthzKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
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
		mint.NewAppModule(
			appCodec,
			app.MintKeeper,
			app.AccountKeeper,
			minttypes.DefaultInflationCalculationFn,
		),
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
		groupmodule.NewAppModule(
			appCodec,
			app.GroupKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
		exchange.NewAppModule(appCodec, app.ExchangeKeeper, app.AccountKeeper, app.BankKeeper),
		// epochModule,

		// evm.NewAppModule(
		// 	app.EvmKeeper,
		// 	app.AccountKeeper,
		// 	app.ParamsKeeper.Subspace(evmtypes.ModuleName),
		// ),
		// feemarket.NewAppModule(
		// 	app.FeeMarketKeeper,
		// 	app.ParamsKeeper.Subspace(feemarkettypes.ModuleName),
		// ),
		// this line is used by starport scaffolding # stargate/app/appModule
	)
	app.sm.RegisterStoreDecoders()
	// overrideModules := map[string]module.AppModuleSimulation{
	// authtypes.ModuleName: auth.NewAppModule(
	// 	app.appCodec,
	// 	app.AccountKeeper,
	// 	authsims.RandomGenesisAccounts,
	// 	// app.GetSubspace(authtypes.ModuleName)
	// ),
	// }
	// reflectionSvc, err := runtimeservices.NewReflectionService()

	// app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)
	// app.sm.RegisterStoreDecoders()

	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	// maxGasWanted := cast.ToUint64(appOpts.Get(srvflags.EVMMaxTxGasWanted))
	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeeGrantKeeper,
				SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			// EvmKeeper:         app.EvmKeeper,
			// StakingKeeper:     app.StakingKeeper,
			IBCKeeper: app.IBCKeeper,
			// FeeMarketKeeper:   app.FeeMarketKeeper,
			// Cdc:               appCodec,
			// MaxTxGasWanted:    maxGasWanted,
			WasmConfig: &wasmConfig,
			// IBCKeeper:         app.IBCKeeper,
			TXCounterStoreKey: keys[wasm.StoreKey],
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %s", err))
	}

	app.SetAnteHandler(anteHandler)
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	// app.SetPostHandler(postHandler)
	// app.tpsCounter = newTPSCounter(logger)
	// go func() {
	// Unfortunately golangci-lint is so pedantic
	// so we have to ignore this error explicitly.
	// _ = app.tpsCounter.start(context.Background())
	// }()

	if manager := app.SnapshotManager(); manager != nil {
		err := manager.RegisterExtensions(
		// &wasmkeeper.NewWasmSnapshotter(
		// 	app.CommitMultiStore(),
		// 	&app.WasmKeeper,
		// ),
		)
		if err != nil {
			panic(fmt.Errorf("failed to register snapshot extensions: %s", err))
		}

	}
	// app.setupUpgradeHandlers()

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			logger.Error("error on loading last version", "err", err)
			tmos.Exit(fmt.Sprintf("failed to load latest version: %s", err))
		}
		ctx := app.NewUncachedContext(true, tmproto.Header{})
		if err := app.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit(fmt.Sprintf("failed initialize pinned codes %s", err))
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedClaimKeeper = scopedClaimKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.ScopedICAControllerKeeper = scopedICAControllerKeeper
	app.ScopedICAHostKeeper = scopedICAHostKeeper
	app.ScopedIBCFeeKeeper = scopedIBCFeeKeeper
	app.ScopedEngineKeeper = scopedEngineKeeper
	app.ScopedHooksKeeper = scopedHooksKeeper
	app.ScopedIBCMockKeeper = scopedIBCMockKeeper
	app.ScopedICAMockKeeper = scopedICAMockKeeper
	app.ScopedFeeMockKeeper = scopedFeeMockKeeper
	app.ScopedWasmKeeper = scopedWasmKeeper
	app.ScopedInterTxKeeper = scopedInterTxKeeper
	app.ScopedMarketKeeper = scopedMarketKeeper
	app.ScopedClaimKeeper = scopedClaimKeeper
	app.ScopedOnsKeeper = scopedOnsKeeper
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

	return app
}

func (app *App) ModuleManager() module.Manager {
	return *app.mm
}

func (app *App) ModuleConfigurator() module.Configurator {
	return app.configurator
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
	icaRawGenesisState := genesisState[icatypes.ModuleName]
	var icaGenesisState icagenesistypes.GenesisState
	if err := app.cdc.UnmarshalJSON(icaRawGenesisState, &icaGenesisState); err != nil {
		panic(err)
	}
	icaGenesisState.HostGenesisState.Params.AllowMessages = []string{"*"} // allow all messages
	genesisJSON, err := app.appCodec.MarshalJSON(&icaGenesisState)
	if err != nil {
		panic(err)
	}

	genesisState[icatypes.ModuleName] = genesisJSON
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
	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}
	return blockedAddrs
	// modAccAddrs := app.ModuleAccountAddrs()

	// delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())
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

func RegisterSwaggerAPI(rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
	// panic(err)
	// }

	// register app's OpenAPI routes.
	RegisterSwaggerAPI(apiSvr.Router)
	apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.FS(docs.Docs)))
	apiSvr.Router.HandleFunc("/", docs.Handler(Name, "/static/openapi.yml"))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(
		app.GRPCQueryRouter(),
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
	paramsKeeper.Subspace(tokenmoduletypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(grantstypes.ModuleName)
	paramsKeeper.Subspace(farmingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	// paramsKeeper.Subspace(mintmoduletypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)
	// paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(icatypes.ModuleName)
	paramsKeeper.Subspace(liquiditymoduletypes.ModuleName)
	paramsKeeper.Subspace(onsmoduletypes.ModuleName)
	paramsKeeper.Subspace(marketmoduletypes.ModuleName)
	paramsKeeper.Subspace(claimmoduletypes.ModuleName)
	// paramsKeeper.Subspace(epochtypes.ModuleName)
	paramsKeeper.Subspace(exchangetypes.ModuleName)
	paramsKeeper.Subspace(reservemoduletypes.ModuleName)
	paramsKeeper.Subspace(lendmoduletypes.ModuleName)
	paramsKeeper.Subspace(emissionstypes.ModuleName)
	paramsKeeper.Subspace(enginetypes.ModuleName)
	paramsKeeper.Subspace(automationtypes.ModuleName)
	paramsKeeper.Subspace(hookstypes.ModuleName)
	paramsKeeper.Subspace(vaulttupes.ModuleName)
	paramsKeeper.Subspace(ibcfeetypes.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	// ethermint subspaces
	paramsKeeper.Subspace(evmtypes.ModuleName).
		WithKeyTable(evmtypes.ParamKeyTable())
		//nolint: staticcheck
	paramsKeeper.Subspace(feemarkettypes.ModuleName).WithKeyTable(feemarkettypes.ParamKeyTable())
	// paramsKeeper.Subspace(oraclemoduletypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}

//	func (app *App) setupUpgradeHandlers() {
//		app.UpgradeKeeper.SetUpgradeHandler(
//			v6.UpgradeName,
//			v6.CreateUpgradeHandler(
//				app.mm,
//				app.configurator,
//				app.appCodec,
//				app.keys[capabilitytypes.ModuleName],
//				app.CapabilityKeeper,
//				intertxtypes.ModuleName,
//			),
//		)
//	}
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

type EmptyAppOptions struct{}

func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

func BlockedAddresses() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range GetMaccPerms() {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	// allow the following addresses to receive funds
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

func (app *App) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}
