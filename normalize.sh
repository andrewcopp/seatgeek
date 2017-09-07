#!/bin/sh

SCRIPT_PATH=${0%/*}
echo $SCRIPT_PATH
if [ "$0" != "$SCRIPT_PATH" ] && [ "$SCRIPT_PATH" != "" ]; then
    cd $SCRIPT_PATH
fi
go run normalize.go "$0"
