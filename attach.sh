#!/bin/sh

GETH=~/ethereum/go/bin/geth
GETH_SYNC_MODE=$1
DATA_DIRECTORY=~/ethereum/$GETH_SYNC_MODE/data


echo "Attaching to geth in $GETH_SYNC_MODE @ $DATA_DIRECTORY"

$GETH attach ipc:$DATA_DIRECTORY/geth.ipc
