#!/bin/sh -l

time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT
ollod init $0 --chain-id "ollo-testnet-1"
