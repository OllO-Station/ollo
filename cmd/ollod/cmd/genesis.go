package cmd

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"time"

	"github.com/ollo-station/ollo/app"

	sdkmath "cosmossdk.io/math"
	"github.com/tendermint/tendermint/crypto/ed25519"

	// "github.com/mint/k/ignite/pkg/cliui/colors"
	"github.com/spf13/cobra"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	color "github.com/fatih/color"
)

func GenesisCmd(defaultNodeHome string, mbm module.BasicManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "genesis",
		Short:   color.New(color.Faint, color.Italic).SprintFunc()("Genesis subcommands "),
		Aliases: []string{"gen", "g"},
		Long: `Genesis utilites & subcommands
Example:
	ollod genesis [command]
  `,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			depCdc := clientCtx.Codec
			cdc := depCdc
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}
			if err = mbm.ValidateGenesis(cdc, clientCtx.TxConfig, appState); err != nil {
				return fmt.Errorf("error validating genesis file: %s", err.Error())
			}
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	flags.AddQueryFlagsToCmd(cmd)
	cmd.AddCommand(
		PregenesisCmd(app.DefaultNodeHome, app.ModuleBasics),
		ExportBalancesCmd(),
		StakedCSVCmd(),
		genutilcli.MigrateGenesisCmd(),
		AddGenesisAccountCmd(app.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		preUpgradeCommand(),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.GenTxCmd(
			app.ModuleBasics,
			app.MakeEncodingConfig().TxConfig,
			banktypes.GenesisBalancesIterator{},
			app.DefaultNodeHome,
		),
	)
	// )
	return cmd
}

func PregenesisCmd(defaultNodeHome string, mbm module.BasicManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pregenesis",
		Short: "Prepare a genesis file ",
		Long: `Prepare a genesis file 
Example:
	ollod g pregenesis mainnet ollo-1
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			depCdc := clientCtx.Codec
			cdc := depCdc
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}
			var genesisParams GenesisParams
			network := args[0]
			if network == "testnet" {
				genesisParams = TestnetGenesisParams()
			} else if network == "mainnet" {
				genesisParams = MainnetGenesisParams()
			} else {
				return fmt.Errorf("please choose 'mainnet' or 'testnet'")
			}
			chainID := args[1]
			appState, genDoc, err = PrepareGenesis(
				clientCtx,
				appState,
				genDoc,
				genesisParams,
				chainID,
			)
			if err != nil {
				return err
			}
			if err = mbm.ValidateGenesis(cdc, clientCtx.TxConfig, appState); err != nil {
				return fmt.Errorf("error validating genesis file: %s", err.Error())
			}
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON
			err = genutil.ExportGenesisFile(genDoc, genFile)
			return err
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func PrepareGenesis(
	clientCtx client.Context,
	appState map[string]json.RawMessage,
	genDoc *tmtypes.GenesisDoc,
	genesisParams GenesisParams,
	chainID string,
) (map[string]json.RawMessage, *tmtypes.GenesisDoc, error) {
	depCdc := clientCtx.Codec
	cdc := depCdc
	genDoc.ChainID = chainID
	genDoc.GenesisTime = genesisParams.GenesisTime
	genDoc.ConsensusParams = genesisParams.ConsensusParams
	stakingGenState := stakingtypes.GetGenesisStateFromAppState(depCdc, appState)
	stakingGenState.Params = genesisParams.StakingParams
	stakingGenStateBz, err := cdc.MarshalJSON(stakingGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal staking genesis state: %w", err)
	}
	appState[stakingtypes.ModuleName] = stakingGenStateBz

	mintGenState := minttypes.DefaultGenesisState()
	mintGenState.Params = genesisParams.MintParams
	mintGenStateBz, err := cdc.MarshalJSON(mintGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal mint genesis state: %w", err)
	}
	appState[minttypes.ModuleName] = mintGenStateBz

	distributionGenState := distributiontypes.DefaultGenesisState()
	distributionGenState.Params = genesisParams.DistributionParams
	distributionGenStateBz, err := cdc.MarshalJSON(distributionGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal distribution genesis state: %w", err)
	}
	appState[distributiontypes.ModuleName] = distributionGenStateBz

	govGenState := govtypes.DefaultGenesisState()
	govGenState.DepositParams = genesisParams.GovParams.DepositParams
	govGenState.TallyParams = genesisParams.GovParams.TallyParams
	govGenState.VotingParams = genesisParams.GovParams.VotingParams
	govGenStateBz, err := cdc.MarshalJSON(govGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal gov genesis state: %w", err)
	}
	appState["gov"] = govGenStateBz

	crisisGenState := crisistypes.DefaultGenesisState()
	crisisGenState.ConstantFee = genesisParams.CrisisConstantFee
	crisisGenStateBz, err := cdc.MarshalJSON(crisisGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal crisis genesis state: %w", err)
	}
	appState[crisistypes.ModuleName] = crisisGenStateBz

	slashingGenState := slashingtypes.DefaultGenesisState()
	slashingGenState.Params = genesisParams.SlashingParams
	slashingGenStateBz, err := cdc.MarshalJSON(slashingGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal slashing genesis state: %w", err)
	}
	appState[slashingtypes.ModuleName] = slashingGenStateBz
	return appState, genDoc, nil
}

type GenesisParams struct {
	AirdropSupply sdk.Int

	StrategicReserveAccounts []banktypes.Balance

	ConsensusParams *tmproto.ConsensusParams

	GenesisTime         time.Time
	NativeCoinMetadatas []banktypes.Metadata

	StakingParams      stakingtypes.Params
	MintParams         minttypes.Params
	DistributionParams distributiontypes.Params
	GovParams          govtypes.Params

	CrisisConstantFee sdk.Coin

	SlashingParams slashingtypes.Params
}

func MainnetGenesisParams() GenesisParams {
	genParams := GenesisParams{}

	genParams.GenesisTime = time.Date(2021, 6, 18, 17, 0, 0, 0, time.UTC)

	genParams.NativeCoinMetadatas = []banktypes.Metadata{
		{
			Description: "The native token of OLLO",
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    "uollo",
					Exponent: 0,
					Aliases:  nil,
				},
				{
					Denom:    "OLLO",
					Exponent: 6,
					Aliases:  nil,
				},
			},
			Base:    "uollo",
			Display: "OLLO",
		},
		{
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    "uuso",
					Exponent: 0,
					Aliases:  nil,
				},
				{
					Denom:    "USO",
					Exponent: 6,
					Aliases:  nil,
				},
			},
			Base:    "uuso",
			Display: "USO",
		},
	}

	genParams.StrategicReserveAccounts = []banktypes.Balance{}

	genParams.StakingParams = stakingtypes.DefaultParams()
	genParams.StakingParams.UnbondingTime = time.Hour * 24 * 7 * 2
	genParams.StakingParams.MaxValidators = 100
	genParams.StakingParams.BondDenom = genParams.NativeCoinMetadatas[0].Base
	genParams.StakingParams.MinCommissionRate = sdk.MustNewDecFromStr("0.05")

	genParams.MintParams = minttypes.DefaultParams()
	genParams.MintParams.MintDenom = genParams.NativeCoinMetadatas[0].Base

	genParams.DistributionParams = distributiontypes.DefaultParams()
	genParams.DistributionParams.BaseProposerReward = sdk.MustNewDecFromStr("0.01")
	genParams.DistributionParams.BonusProposerReward = sdk.MustNewDecFromStr("0.04")
	genParams.DistributionParams.CommunityTax = sdk.MustNewDecFromStr("0")
	genParams.DistributionParams.WithdrawAddrEnabled = true

	genParams.GovParams = govtypes.DefaultParams()
	genParams.GovParams.DepositParams.MaxDepositPeriod = time.Hour * 24 * 14
	genParams.GovParams.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		sdk.NewInt(2_500_000_000),
	))
	genParams.GovParams.TallyParams.Quorum = sdk.MustNewDecFromStr("0.2")
	genParams.GovParams.VotingParams.VotingPeriod = time.Hour * 24 * 3

	genParams.CrisisConstantFee = sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		sdk.NewInt(500_000_000_000),
	)

	genParams.SlashingParams = slashingtypes.DefaultParams()
	genParams.SlashingParams.SignedBlocksWindow = int64(30000)
	genParams.SlashingParams.MinSignedPerWindow = sdk.MustNewDecFromStr("0.05")
	genParams.SlashingParams.DowntimeJailDuration = time.Minute
	genParams.SlashingParams.SlashFractionDoubleSign = sdk.MustNewDecFromStr("0.05")
	genParams.SlashingParams.SlashFractionDowntime = sdk.ZeroDec()

	genParams.ConsensusParams = tmtypes.DefaultConsensusParams()
	genParams.ConsensusParams.Block.MaxBytes = 5 * 1024 * 1024
	genParams.ConsensusParams.Block.MaxGas = 6_000_000
	genParams.ConsensusParams.Evidence.MaxAgeDuration = genParams.StakingParams.UnbondingTime
	genParams.ConsensusParams.Evidence.MaxAgeNumBlocks = int64(
		genParams.StakingParams.UnbondingTime.Seconds(),
	) / 3
	genParams.ConsensusParams.Version.AppVersion = 1

	return genParams
}

func TestnetGenesisParams() GenesisParams {
	genParams := MainnetGenesisParams()

	genParams.GenesisTime = time.Now()

	genParams.StakingParams.UnbondingTime = time.Hour * 24 * 7 * 2

	genParams.GovParams.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		sdk.NewInt(1000000),
	))
	genParams.GovParams.TallyParams.Quorum = sdk.MustNewDecFromStr("0.0000000001")
	genParams.GovParams.VotingParams.VotingPeriod = time.Second * 300

	return genParams
}

const GentxFilename = "gentx.json"

type (
	// GentxInfo represents the basic info about gentx file
	GentxInfo struct {
		DelegatorAddress string
		PubKey           ed25519.PubKey
		SelfDelegation   sdk.Coin
		Memo             string
	}

	// StargateGentx represents the stargate gentx file
	StargateGentx struct {
		Body struct {
			Messages []struct {
				DelegatorAddress string `json:"delegator_address"`
				ValidatorAddress string `json:"validator_address"`
				PubKey           struct {
					Type string `json:"@type"`
					Key  string `json:"key"`
				} `json:"pubkey"`
				Value struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"value"`
			} `json:"messages"`
			Memo string `json:"memo"`
		} `json:"body"`
	}
)

// GentxFromPath returns GentxInfo from the json file
func GentxFromPath(path string) (info GentxInfo, gentx []byte, err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return info, gentx, errors.New("chain home folder is not initialized yet: " + path)
	}

	gentx, err = os.ReadFile(path)
	if err != nil {
		return info, gentx, err
	}
	return ParseGentx(gentx)
}

// ParseGentx returns GentxInfo and the gentx file in bytes
// TODO refector. no need to return file, it's already given as gentx in the argument.
func ParseGentx(gentx []byte) (info GentxInfo, file []byte, err error) {
	// Try parsing Stargate gentx
	var stargateGentx StargateGentx
	if err := json.Unmarshal(gentx, &stargateGentx); err != nil {
		return info, gentx, err
	}
	if stargateGentx.Body.Messages == nil {
		return info, gentx, errors.New("the gentx cannot be parsed")
	}

	// This is a stargate gentx
	if len(stargateGentx.Body.Messages) != 1 {
		return info, gentx, errors.New("add validator gentx must contain 1 message")
	}

	info.Memo = stargateGentx.Body.Memo
	info.DelegatorAddress = stargateGentx.Body.Messages[0].DelegatorAddress

	pb := stargateGentx.Body.Messages[0].PubKey.Key
	info.PubKey, err = base64.StdEncoding.DecodeString(pb)
	if err != nil {
		return info, gentx, fmt.Errorf("invalid validator public key %s", err.Error())
	}

	amount, ok := sdkmath.NewIntFromString(stargateGentx.Body.Messages[0].Value.Amount)
	if !ok {
		return info, gentx, errors.New("the self-delegation inside the gentx is invalid")
	}

	info.SelfDelegation = sdk.NewCoin(
		stargateGentx.Body.Messages[0].Value.Denom,
		amount,
	)

	return info, gentx, nil
}
