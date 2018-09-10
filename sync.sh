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


nohup $GETH --syncmode="$GETH_SYNC_MODE" --cache=$CACHE_SIZE \
--verbosity $LOGGING_VERBOSITY --datadir $DATA_DIRECTORY \
--wsapi admin,db,debug,eth,net,txpool,web3 --rpcapi admin,db,debug,eth,net,txpool,web3 --rpc > $LOG_DIRECTORY/geth_$NOW.log 2>&1 &
