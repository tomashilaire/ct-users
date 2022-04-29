#!/bin/bash

REPO_NAME=${PWD##*/}

cd k8s/docker
cp ../../cmd/grpcserver/grpcsvr .

docker build -t ${REPO_NAME} -f Dockerfile.grpc .
docker inspect ${REPO_NAME}

cd ../..