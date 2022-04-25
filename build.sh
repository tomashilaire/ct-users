#!/bin/bash

#go clean --cache && go test -v -cover entity/...

### GRPC SERVER ###
go build -o cmd/grpcserver/grpcsvr cmd/grpcserver/main.go

### HTTP SERVER ###
go build -o cmd/httpserver/httpsvr cmd/httpserver/main.go