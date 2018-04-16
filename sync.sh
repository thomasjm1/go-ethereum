#!/bin/sh

NOW=$(date +"%s")

echo "Building directories"

mkdir ~/eth
mkdir ~/eth_logs

echo "Starting geth"

nohup ~/go/bin/geth --syncmode "full" --cache=512 --verbosity 3 --datadir ~/eth --rpc --rpcapi admin,db,debug,eth,net,web3 > ~/eth_logs/geth_$NOW.log 2>&1 &
