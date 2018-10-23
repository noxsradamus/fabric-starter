#!/usr/bin/env bash
VERSION=$(printenv FABRIC_VERSION)
BENCHMARK_FOLDER=$(printenv BENCHMARK_FOLDER)
npm i grpc@1.10.1 fabric-ca-client@${VERSION} fabric-client@${VERSION}
npm i
npm rebuild
node benchmark/${BENCHMARK_FOLDER}/main
cp $(ls | grep report) ./reports