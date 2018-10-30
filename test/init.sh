#!/usr/bin/env bash
echo "Clone hyperledger/caliper..."
git clone https://gitlab.altoros.com/intprojects/Internal-Blockchain-Projects/caliper.git

cd caliper
git fetch
echo "Switching to the benchmark branch"
git checkout -b benchmark origin/benchmark
echo "Pulling changes"
git pull
cd ..

echo "Copy ./docker/. to ./caliper"
cp -ra ./docker/. ./caliper

FOLDER=$1
if [ -z $FOLDER ]; then
FOLDER=test
fi
# copy into benchmark
# network-config.json -> fabric.json
# copy into network/fabric/<folder> from ./artifacts/channel & /crypto-config
echo "Copy ../artifacts/channel to ./network/fabric/$FOLDER/channel"
rm -rf ./network/fabric/${FOLDER}/channel
cp -ra ../artifacts/channel ./network/fabric/${FOLDER}/channel
echo "Copy ../artifacts/crypto-config to ./network/fabric/$FOLDER/crypto-config"
rm -rf ./network/fabric/${FOLDER}/crypto-config
cp -ra ../artifacts/crypto-config ./network/fabric/${FOLDER}/crypto-config
# copy chaincode into src/contract/fabric/<folder> from ./chaincode
echo "Copy ../chaincode to ./src/contract/fabric/$FOLDER"
cp -ra ../chaincode ./src/contract/fabric/${FOLDER}

echo "Rename certs"
source update_certs.sh
