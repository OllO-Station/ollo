#!/bin/bash

VALIDATOR="validator2"
CHAINID="ollo-testnet-2"
MONIKER="ollo-testnet-seed"
MAINNODE_RPC="https://rpc.ollo.zone"
MAINNODE_ID="[main node's seed_id]@73.14.46.216:26656"
KEYRING="test"
CONFIG="$HOME/.ollo/config/config.toml"
APPCONFIG="$HOME/.ollo/config/app.toml"

# install chain binary file
make install

# Set moniker and chain-id for chain (Moniker can be anything, chain-id must be same mainnode)
ollod init $MONIKER --chain-id=$CHAINID

# Fetch genesis.json from genesis node
curl $MAINNODE_RPC/genesis? | jq ".result.genesis" > ~/.ollod/config/genesis.json

ollod validate-genesis

# set seed to main node's id manually
# sed -i 's/seeds = ""/seeds = "'$MAINNODE_ID'"/g' ~/.ollod/config/config.toml

# add for rpc
sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/g' "$CONFIG"
sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' "$CONFIG"
sed -i '/\[api\]/,+3 s/enable = false/enable = true/' "$APPCONFIG"
sed -i '/\[api\]/,+3 s/swagger = false/swagger = true/' "$APPCONFIG"
sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g'  "$APPCONFIG"
sed -i 's/api = "eth,net,web3"/api = "eth,txpool,personal,net,debug,web3"/g' "$APPCONFIG"

# add account for validator in the node
ollod keys add $VALIDATOR --keyring-backend $KEYRING

# run node
ollod start --rpc.laddr tcp://0.0.0.0:26657 --pruning=nothing