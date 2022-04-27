check_swagger:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger
	swagger generate spec -o ./swagger.yaml --scan-models

protos:
# con buf (permite generar protos + documentacion, requiere yaml)
#	cd protos; buf generate;
# con protoc
	protoc -I=./internal/handlers/filesprotohdl --go_out=plugins=grpc:. ./internal/handlers/filesprotohdl/*.proto
	protoc -I=./internal/handlers/entityprotohdl --go_out=plugins=grpc:. ./internal/handlers/entityprotohdl/*.proto
	protoc -I=./internal/handlers/usersprotohdl --go_out=plugins=grpc:. ./internal/handlers/usersprotohdl/*.proto
