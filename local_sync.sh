#!/bin/sh

GETH=~/ethereum/go/bin/geth
GETH_SYNC_MODE=$1
DATA_DIRECTORY=~/ethereum/$GETH_SYNC_MODE/data
LOG_DIRECTORY=~/ethereum/$GETH_SYNC_MODE/logs

CACHE_SIZE=512
LOGGING_VERBOSITY=3

mkdir ~/ethereum/$GETH_SYNC_MODE
mkdir $DATA_DIRECTORY
mkdir $LOG_DIRECTORY
NOW=$(date +"%s")

echo "Starting geth in $GETH_SYNC_MODE @ $DATA_DIRECTORY"

NODES=enode://fd609e0de7edecf10c2dddeff399ccbf1c4f487c62f9659bb581a6f82c7a3ac16cf6e5e306257c174fa7a2b290d8178fd7e30eb1c67c401a20f8f1a3e0309094@:30303

nohup $GETH --syncmode="$GETH_SYNC_MODE" --nodiscover --cache=$CACHE_SIZE --verbosity $LOGGING_VERBOSITY --datadir $DATA_DIRECTORY > $LOG_DIRECTORY/geth_$NOW.log 2>&1 &
sleep 5
$GETH attach ipc:$DATA_DIRECTORY/geth.ipc --exec 'admin.addPeer("enode://fd609e0de7edecf10c2dddeff399ccbf1c4f487c62f9659bb581a6f82c7a3ac16cf6e5e306257c174fa7a2b290d8178fd7e30eb1c67c401a20f8f1a3e0309094@128.237.153.99:30303")'
