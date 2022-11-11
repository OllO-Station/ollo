#!/bin/bash
rm -rf ~/.ollo
make gp
make install
ollod init --chain-id "ollot" "test"
ollod keys add --keyring-backend test "testkey" 
ollod add-genesis-account "testkey" 1000000000000stake --keyring-backend test
ollod gentx "testkey" 1000000000000stake --pubkey "$(ollod tendermint show-validator)" --keyring-backend test
ollod collect-gentxs
ollod unsafe-reset-all
ollod collect-gentxs
ollod start
