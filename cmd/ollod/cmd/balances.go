package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"

	liquiditytypes "github.com/ollo-station/ollo/x/liquidity/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const FlagSelectPoolIds = "breakdown-by-pool-ids"
const FlagMinimumStakeAmount = "minimum-stake-amount"

type DeriveSnapshot struct {
	NumberAccounts uint64                    `json:"num_accounts"`
	Accounts       map[string]DerivedAccount `json:"accounts"`
}

type DerivedAccount struct {
	Address             string               `json:"address"`
	Staked              sdk.Int              `json:"staked"`
	UnbondingStake      sdk.Int              `json:"unbonding_stake"`
	Bonded              sdk.Coins            `json:"bonded"`
	BondedBySelectPools map[uint64]sdk.Coins `json:"bonded_by_select_pools"`
	TotalBalances       sdk.Coins            `json:"total_balances"`
}

func newDerivedAccount(address string) DerivedAccount {
	return DerivedAccount{
		Address:        address,
		Staked:         sdk.ZeroInt(),
		UnbondingStake: sdk.ZeroInt(),
		Bonded:         sdk.Coins{},
	}
}

func underlyingCoins(originCoins sdk.Coins, pools map[string]liquiditytypes.Pool) sdk.Coins {
	balances := sdk.Coins{}
	for _, coin := range originCoins {
		balances = balances.Add(coin)
	}
	return balances
}

func getGenStateFromPath(genesisFilePath string) (map[string]json.RawMessage, error) {
	genState := make(map[string]json.RawMessage)

	genesisFile, err := os.Open(filepath.Clean(genesisFilePath))
	if err != nil {
		return genState, err
	}
	defer genesisFile.Close()

	byteValue, _ := io.ReadAll(genesisFile)

	var doc tmtypes.GenesisDoc
	err = tmjson.Unmarshal(byteValue, &doc)
	if err != nil {
		return genState, err
	}

	err = json.Unmarshal(doc.AppState, &genState)
	if err != nil {
		panic(err)
	}
	return genState, nil
}

func ExportBalancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance-export [input-genesis] [output-snapshot]",
		Short: "Export balances from exported genesis file",
		Long: `Export a balances from exported genesis file
Example:
	ollod balance-export ./genesis.json ./snapshot.json
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			genesisFile := args[0]
			genState, err := getGenStateFromPath(genesisFile)
			if err != nil {
				return err
			}
			snapshotOutput := args[1]

			snapshotAccs := make(map[string]DerivedAccount)

			bankGenesis := banktypes.GenesisState{}
			if len(genState["bank"]) > 0 {
				clientCtx.Codec.MustUnmarshalJSON(genState["bank"], &bankGenesis)
			}
			for _, balance := range bankGenesis.Balances {
				address := balance.Address
				acc, ok := snapshotAccs[address]
				if !ok {
					acc = newDerivedAccount(address)
				}

				snapshotAccs[address] = acc
			}

			stakingGenesis := stakingtypes.GenesisState{}
			if len(genState["staking"]) > 0 {
				clientCtx.Codec.MustUnmarshalJSON(genState["staking"], &stakingGenesis)
			}
			for _, unbonding := range stakingGenesis.UnbondingDelegations {
				address := unbonding.DelegatorAddress
				acc, ok := snapshotAccs[address]
				if !ok {
					acc = newDerivedAccount(address)
				}

				unbondingOsmos := sdk.NewInt(0)
				for _, entry := range unbonding.Entries {
					unbondingOsmos = unbondingOsmos.Add(entry.Balance)
				}

				acc.UnbondingStake = acc.UnbondingStake.Add(unbondingOsmos)

				snapshotAccs[address] = acc
			}

			validators := make(map[string]stakingtypes.Validator)
			for _, validator := range stakingGenesis.Validators {
				validators[validator.OperatorAddress] = validator
			}

			for _, delegation := range stakingGenesis.Delegations {
				address := delegation.DelegatorAddress

				acc, ok := snapshotAccs[address]
				if !ok {
					acc = newDerivedAccount(address)
				}

				val := validators[delegation.ValidatorAddress]
				stakedOsmos := delegation.Shares.MulInt(val.Tokens).Quo(val.DelegatorShares).RoundInt()

				acc.Staked = acc.Staked.Add(stakedOsmos)

				snapshotAccs[address] = acc
			}

			for addr, account := range snapshotAccs {
				account.TotalBalances = sdk.NewCoins().
					Add(sdk.NewCoin("uollo", account.Staked)).
					Add(sdk.NewCoin("uollo", account.UnbondingStake)).
					Add(account.Bonded...)
				snapshotAccs[addr] = account
			}
			snapshot := DeriveSnapshot{
				NumberAccounts: uint64(len(snapshotAccs)),
				Accounts:       snapshotAccs,
			}
			fmt.Printf("# accounts: %d\n", len(snapshotAccs))
			snapshotJSON, err := json.MarshalIndent(snapshot, "", "    ")
			if err != nil {
				return fmt.Errorf("failed to marshal snapshot: %w", err)
			}

			err = os.WriteFile(snapshotOutput, snapshotJSON, 0o644)
			return err
		},
	}

	cmd.Flags().String(FlagSelectPoolIds, "",
		"Output a special breakdown for amount LP'd to the provided pools. Usage --breakdown-by-pool-ids=1,2,605")

	return cmd
}

func StakedCSVCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staked-csv [input-balances-file] [output-airdrop-csv]",
		Short: "Export a csv from a exported genesis file",
		Long: `Export a csv from a exported genesis file
Example:
	ollod staked-csv ./balances.json ./airdrop.csv
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			balancesFile := args[0]
			snapshotOutput := args[1]
			minStakeAmount, _ := cmd.Flags().GetInt64(FlagMinimumStakeAmount)

			var deriveSnapshot DeriveSnapshot

			sourceFile, err := os.Open(balancesFile)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			if err := json.NewDecoder(sourceFile).Decode(&deriveSnapshot); err != nil {
				return err
			}

			outputFile, err := os.Create(snapshotOutput)
			if err != nil {
				return err
			}
			defer outputFile.Close()

			writer := csv.NewWriter(outputFile)
			defer writer.Flush()

			header := []string{"address", "staked"}
			if err := writer.Write(header); err != nil {
				return err
			}

			for _, r := range deriveSnapshot.Accounts {
				var csvRow []string
				if r.Staked.GT(sdk.NewInt(minStakeAmount)) {
					csvRow = append(csvRow, r.Address, r.Staked.String())
					if err := writer.Write(csvRow); err != nil {
						return err
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().Int64(FlagMinimumStakeAmount, 0, "Specify minimum amount (non inclusive) accounts that must stake to be included (default: 0)")

	return cmd
}
