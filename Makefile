check_swagger:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger
	swagger generate spec -o ./swagger.yaml --scan-models

protos:
# con buf (permite generar protos + documentacion, requiere yaml)
#	cd protos; buf generate;
# con protoc
	protoc -I=./internal/handlers/usersprotohdl --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./internal/handlers/usersprotohdl/*.proto

	protoc -I=./internal/handlers/usersprotohdl --go_out=./client/go --go-grpc_out=require_unimplemented_servers=false:./client/go ./internal/handlers/usersprotohdl/*.proto

	protoc -I=./internal/handlers/usersprotohdl users.proto --js_out=import_style=commonjs:./client/node --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client/node