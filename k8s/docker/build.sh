#!/bin/bash

REPO_NAME=${PWD##*/}

cd k8s/docker
cp ../../cmd/grpcserver/grpcsvr .

docker build -t ${REPO_NAME//_/-} -f Dockerfile.grpc .
docker inspect ${REPO_NAME//_/-}

cd ../..