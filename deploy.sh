#!/bin/bash
./build.sh
./k8s/docker/build.sh

STAGE=prod
TAG=1.0.0

docker tag am-authentication:latest *******.dkr.ecr.us-east-1.amazonaws.com/${STAGE}/am-authentication:${TAG} || exit 1
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin *******.dkr.ecr.us-east-1.amazonaws.com
docker push *******.dkr.ecr.us-east-1.amazonaws.com/${STAGE}/am-authentication:${TAG} || exit 1

docker rmi am-authentication:latest
docker rmi *******.dkr.ecr.us-east-1.amazonaws.com/${STAGE}/am-authentication:${TAG}
docker builder prune -f
