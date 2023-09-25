#!/bin/bash
rm -rf $HOME/.ollo/

GENESIS=$HOME/od/x/nx/genesis.json


# make four osmosis directories
mkdir $HOME/.ollo
mkdir $HOME/.ollo/validator1
mkdir $HOME/.ollo/validator2
mkdir $HOME/.ollo/validator3

# init all three validators
ollod init --chain-id=testing validator1 --home=$HOME./.ollo/validator1
ollod init --chain-id=testing validator2 --home=$HOME./.ollo/validator2
ollod init --chain-id=testing validator3 --home=$HOME./.ollo/validator3
# create keys for all three validators
ollod keys add validator1 --keyring-backend=test --home=$HOME./.ollo/validator1
ollod keys add validator2 --keyring-backend=test --home=$HOME./.ollo/validator2
ollod keys add validator3 --keyring-backend=test --home=$HOME./.ollo/validator3

cp -r $GENESIS $HOME/.ollo/validator1/config/
cp -r $GENESIS $HOME/.ollo/validator2/config/
cp -r $GENESIS $HOME/.ollo/validator3/config/

ollod add-genesis-account $(ollod keys show validator1 -a --keyring-backend=test --home=$HOME./.ollo/validator1) 1000000000000uollo,1000000000utwise --home=$HOME/.ollo/validator1
ollod add-genesis-account $(ollod keys show validator2 -a --keyring-backend=test --home=$HOME./.ollo/validator2) 1000000000000uollo,1000000000utwise --home=$HOME/.ollo/validator2
ollod add-genesis-account $(ollod keys show validator3 -a --keyring-backend=test --home=$HOME./.ollo/validator3) 1000000000000uollo,1000000000utwise --home=$HOME/.ollo/validator3

ollod gentx validator1 900000000000uollo --keyring-backend=test --home=$HOME./.ollo/validator1 --chain-id=testing
ollod gentx validator2 900000000000uollo --keyring-backend=test --home=$HOME./.ollo/validator2 --chain-id=testing
ollod gentx validator3 900000000000uollo --keyring-backend=test --home=$HOME./.ollo/validator3 --chain-id=testing

mkdir $HOME/.ollo/gentx
cp -r  $HOME/.ollo/validator1/config/gentx/* $HOME/.ollo/gentx/
cp -r  $HOME/.ollo/validator2/config/gentx/* $HOME/.ollo/gentx/
cp -r  $HOME/.ollo/validator3/config/gentx/* $HOME/.ollo/gentx/

cp -r  $HOME/.ollo/gentx/* $HOME/.ollo/validator1/config/gentx/
cp -r  $HOME/.ollo/gentx/* $HOME/.ollo/validator2/config/gentx/
cp -r  $HOME/.ollo/gentx/* $HOME/.ollo/validator3/config/gentx/

ollod collect-gentxs --home=$HOME/.ollo/validator1
ollod collect-gentxs --home=$HOME/.ollo/validator2
ollod collect-gentxs --home=$HOME/.ollo/validator3


# port key (validator1 uses default ports)

# change app.toml values

# validator2
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1316|g' $HOME/.ollo/validator2/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9088|g' $HOME/.ollo/validator2/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9089|g' $HOME/.ollo/validator2/config/app.toml

# validator3
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1315|g' $HOME/.ollo/validator3/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $HOME/.ollo/validator3/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $HOME/.ollo/validator3/config/app.toml

# change config.toml values

# validator1
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.ollo/validator1/config/config.toml
# validator2
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26655|g' $HOME/.ollo/validator2/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26654|g' $HOME/.ollo/validator2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26653|g' $HOME/.ollo/validator2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.ollo/validator3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.ollo/validator2/config/config.toml
# validator3
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26652|g' $HOME/.ollo/validator3/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26651|g' $HOME/.ollo/validator3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.ollo/validator3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.ollo/validator3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.ollo/validator3/config/config.toml



# copy tendermint node id of validator1 to persistent peers of validator2-3
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(ollod tendermint show-node-id --home=$HOME./.ollo/validator1)@$(curl -4 icanhazip.com):26656\"|g" $HOME/.ollo/validator2/config/config.toml
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(ollod tendermint show-node-id --home=$HOME./.ollo/validator1)@$(curl -4 icanhazip.com):26656\"|g" $HOME/.ollo/validator3/config/config.toml


# start all three validators
tmux new -s validator1 -d ollod start --home=$HOME./.ollo/validator1
tmux new -s validator2 -d ollod start --home=$HOME./.ollo/validator2
tmux new -s validator3 -d ollod start --home=$HOME./.ollo/validator3


# send uollo from first validator to second validator
# sleep 7
# ollod tx bank send validator1 $(ollod keys show validator2 -a --keyring-backend=test --home=$HOME./.ollo/validator2) 500000000uosmo --keyring-backend=test --home=$HOME/.osmosisd/validator1 --chain-id=testing --yes
# sleep 7
# ollod tx bank send validator1 $(ollod keys show validator3 -a --keyring-backend=test --home=$HOME./.ollo/validator3) 400000000uosmo --keyring-backend=test --home=$HOME/.osmosisd/validator1 --chain-id=testing --yes

# create second validator
# sleep 7
# ollod tx staking create-validator --amount=500000000uosmo --from=validator2 --pubkey=$(ollod tendermint show-validator --home=$HOME./.ollo/validator2) --moniker="validator2" --chain-id="testing" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="500000000" --keyring-backend=test --home=$HOME/.osmosisd/validator2 --yes
# sleep 7
# ollod tx staking create-validator --amount=400000000uosmo --from=validator3 --pubkey=$(ollod tendermint show-validator --home=$HOME./.ollo/validator3) --moniker="validator3" --chain-id="testing" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="400000000" --keyring-backend=test --home=$HOME/.osmosisd/validator3 --yes

