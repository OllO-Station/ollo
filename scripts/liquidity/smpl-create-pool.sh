#!/bin/bash

# Set localnet configuration
# Reference localnet script to see which $DENOM2s are given to the user accounts in genesis state
BINARY=ollod
CHAIN_ID=ollo-testnet-2
CHAIN_DIR=$HOME/.ollo
USER_1_ADDRESS=ollo1zaavvzxez0elundtn32qnk9lkm8kmcsz0yetz9
USER_2_ADDRESS=ollo1mzgucqnfr2l8cj5apvdpllhzt4zeuh2carhct4

DENOM1=uollo
DENOM2=uwise
DENOM3=umollo

# Ensure jq is installed
if [[ ! -x "$(which jq)" ]]; then
  echo "jq (a tool for parsing json in the command line) is required..."
  echo "https://stedolan.github.io/jq/download/"
  exit 1
fi

# Ensure liquidityd is installed
if ! [ -x "$(which $BINARY)" ]; then
  echo "Error: liquidityd is not installed. Try building $BINARY by 'make install'" >&2
  exit 1
fi

# Ensure localnet is running
if [[ "$(pgrep $BINARY)" == "" ]];then
    echo "Error: localnet is not running. Try running localnet by 'make localnet" 
    exit 1
fi

# liquidityd q bank balances cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu --home ./data/localnet --output json | jq
echo "-> Checking user1 account balances..."
$BINARY q bank balances $USER_1_ADDRESS \
--home "$CHAIN_DIR" \
--output json | jq

# liquidityd q bank balances cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny --home ./data/localnet --output json | jq
echo "-> Checking user2 account balances..."
$BINARY q bank balances $USER_2_ADDRESS \
--home "$CHAIN_DIR" \
--output json | jq

# liquidityd tx liquidity create-pool 1 100000000$DENOM1,100000000$DENOM2 --home ./data/localnet --chain-id localnet --from user1 --keyring-backend test --yes
echo "-> Creating liquidity pool 1..."
$BINARY tx liquidity create-pool 1 100000000$DENOM1,100000000$DENOM2 \
--home "$CHAIN_DIR" \
--chain-id $CHAIN_ID \
--from $USER_1_ADDRESS \
--keyring-backend test \
--yes

sleep 2

# liquidityd tx liquidity create-pool 1 100000000$DENOM1,100000000$DENOM3 --home ./data/localnet --chain-id localnet --from user2 --keyring-backend test --yes
echo "-> Creating liquidity pool 2..."
$BINARY tx liquidity create-pool 1 100000000$DENOM1,100000000$DENOM3 \
--home "$CHAIN_DIR" \
--chain-id $CHAIN_ID \
--from $USER_2_ADDRESS \
--keyring-backend test \
--yes

sleep 2

# liquidityd q bank balances cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu --home ./data/localnet --output json | jq
echo "-> Checking user1 account balances after..."
$BINARY q bank balances $USER_1_ADDRESS \
--home "$CHAIN_DIR" \
--output json | jq

# liquidityd q bank balances cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny --home ./data/localnet --output json | jq
echo "-> Checking user2 account balances after..."
$BINARY q bank balances $USER_2_ADDRESS \
--home "$CHAIN_DIR" \
--output json | jq

# liquidityd q liquidity pools --home ./data/localnet --output json | jq
echo "-> Querying liquidity pools..."
$BINARY q liquidity pools \
--home "$CHAIN_DIR" \
--output json | jq
