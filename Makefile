protos:
# con buf (permite generar protos + documentacion, requiere yaml)
#	cd protos; buf generate;
# con protoc
	protoc -I=./internal/handlers/testprotohdl --go_out=plugins=grpc:. ./internal/handlers/testprotohdl/*.proto