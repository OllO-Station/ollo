#!/bin/bash

export PATH="$PATH:/root/go/bin"

# Function to validate a numeric input
validate_numeric() {
  re='^[0-9]+$'
  if ! [[ $1 =~ $re ]]; then
    echo "Invalid input. Please enter a numeric value."
    exit 1
  fi
}

# Function to validate a floating-point input
validate_float() {
  re='^[0-9]+([.][0-9]+)?$'
  if ! [[ $1 =~ $re ]]; then
    echo "Invalid input. Please enter a numeric or floating-point value."
    exit 1
  fi
}

# Function to validate an email address
validate_email() {
  if [[ ! $1 =~ ^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$ ]]; then
    echo "Invalid email address. Please enter a valid email."
    exit 1
  fi
}

# Prompt the user for input
read -p "Enter amount (e.g., 1000000000000uollo(must be greater than min self delegation)): " amount
read -p "Enter commission rate (e.g., 0.10): " commission_rate
read -p "Enter commission max rate (e.g., 0.20): " commission_max_rate
read -p "Enter commission max change rate (e.g., 0.05): " commission_max_change_rate
read -p "Enter min self delegation (e.g., 1000000): " min_self_delegation
read -p "Enter your website: " website
read -p "Enter details: " details
read -p "Enter security contact email: " security_contact
read -p "Enter identity (e.g., KEYBASE PGP): " identity


# Validate numeric inputs
validate_numeric "$min_self_delegation"

# Validate floating-point inputs
validate_float "$commission_rate"
validate_float "$commission_max_rate"
validate_float "$commission_max_change_rate"

# Validate email address
# validate_email "$security_contact"


VALIDATOR="validator2"
CHAINID="ollo-testnet-2"
MONIKER="ollo-testnet-seed"
MAINNODE_RPC="https://rpc.ollo.zone"

KEYRING="test"
CONFIG="$HOME/.ollo/config/config.toml"
APPCONFIG="$HOME/.ollo/config/app.toml"

# install chain binary file
make install

# Set moniker and chain-id for chain (Moniker can be anything, chain-id must be same mainnode)
rm -rf $HOME/.ollo
ollod init $MONIKER --chain-id=$CHAINID --overwrite 

# Fetch genesis.json from genesis node
curl $MAINNODE_RPC/genesis? | jq ".result.genesis" > $HOME/.ollo/config/genesis.json

ollod validate-genesis

# Use curl to make the HTTP request and capture the response
response=$(curl -s "$MAINNODE_RPC/status")

# Check if the curl command was successful
if [ $? -eq 0 ]; then
  # Extract the seed ID from the response using tools like awk, grep, or jq
  seed_id=$(echo "$response" | jq -r '.result.node_info.id')

  # Check if the seed ID is not empty
  if [ -n "$seed_id" ]; then
    echo "Seed ID: $seed_id"
  else
    echo "Failed to retrieve the seed ID."
  fi
else
  echo "HTTP request failed."
fi

MAINNODE_ID="$seed_id@73.14.46.216:26656"

# # set seed to main node's id manually
sed -i 's/persistent_peers = ""/persistent_peers = "'$MAINNODE_ID'"/g' ~/.ollo/config/config.toml

# add for rpc
sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/g' "$CONFIG"
sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' "$CONFIG"
sed -i '/\[api\]/,+3 s/enable = false/enable = true/' "$APPCONFIG"
sed -i '/\[api\]/,+3 s/swagger = false/swagger = true/' "$APPCONFIG"
sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g'  "$APPCONFIG"
sed -i 's/api = "eth,net,web3"/api = "eth,txpool,personal,net,debug,web3"/g' "$APPCONFIG"

# add account for validator in the node
ollod keys add $VALIDATOR --recover

sudo tee /etc/systemd/system/ollo.service > /dev/null <<EOF  
[Unit]
Description=OLLO Daemon
After=network-online.target
[Service]
User=$USER
ExecStart=$(which ollod) start
Restart=always
RestartSec=3
LimitNOFILE=8192
[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload
sudo systemctl enable ollo

# Start the service
# sudo systemctl start ollo > /var/log/ollo-log.log 2>&1

sudo systemctl start ollo

# Check log output
# journalctl -fu ollo -o cat


sleep 7

# Run the create-validator command with user-provided input
ollod tx staking create-validator \
  --amount "$amount" \
  --pubkey "$(ollod tendermint show-validator)" \
  --moniker "$MONIKER" \
  --chain-id "$CHAINID" \
  --commission-rate "$commission_rate" \
  --commission-max-rate "$commission_max_rate" \
  --commission-max-change-rate "$commission_max_change_rate" \
  --min-self-delegation "$min_self_delegation" \
  --gas "auto" \
  --gas-adjustment "1.5" \
  --from "$VALIDATOR" \
  --website "$website" \
  --details "$details" \
  --security-contact "$security_contact" \
  --identity "$identity"