package client

import (
	"bitbucket.org/agroproag/am_authentication/client/go/pb"
	"context"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

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
