package client

import (
	"context"
	"github.com/tomashilaire/ct-users/client/go/pb"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

// Authenticate calls sign up RPC
func (grpcClient *GrpcClient) Authenticate(token string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", token)
	resp, err := grpcClient.authService.Authenticate(newCtx, new(pb.AuthenticateRequest))
	if err != nil {
		log.Println("cannot authenticate, error: ", err)
		return "", err
	}
	return resp.AuthId, nil
}
