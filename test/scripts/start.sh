#!/usr/bin/env bash
VERSION=$(printenv FABRIC_VERSION)
GRPC_VERSION=$(printenv GRPC_VERSION)
BENCHMARK_FOLDER=$(printenv BENCHMARK_FOLDER)
npm i grpc@${GRPC_VERSION} fabric-ca-client@${VERSION} fabric-client@${VERSION}
npm i
npm rebuild
node benchmark/${BENCHMARK_FOLDER}/main
cp $(ls | grep report) ./reports