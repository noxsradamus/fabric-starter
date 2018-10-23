#!/usr/bin/env bash
export BENCHMARK_FOLDER=$1
docker-compose build
docker-compose up