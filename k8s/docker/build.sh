#!/bin/bash

cd k8s/docker
cp ../../cmd/grpcserver/grpcsvr .

docker build -t basic_api_hex -f Dockerfile.grpc .
docker inspect basic_api_hex

cd ../..