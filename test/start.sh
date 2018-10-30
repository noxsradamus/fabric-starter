#!/usr/bin/env bash
FOLDER=$2
if [ -z $FOLDER ]; then
FOLDER=test
fi

CONFIG_FILE=$1
if [ -z $CONFIG_FILE ]; then
CONFIG_FILE=config.json
fi

export BENCHMARK_FOLDER=$FOLDER
export CONFIG_FILE=$CONFIG_FILE
docker-compose build
docker-compose up