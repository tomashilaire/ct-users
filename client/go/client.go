package client

import (
	"github.com/tomashilaire/ct-users/client/go/pb"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	authService pb.AuthenticationClient
}

// NewGrpcClient returns a new grpc client
func NewGrpcClient(cc *grpc.ClientConn) *GrpcClient {
	authService := pb.NewAuthenticationClient(cc)
	return &GrpcClient{authService: authService}
}
