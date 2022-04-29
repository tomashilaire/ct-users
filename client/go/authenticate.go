package client

import (
	"context"
	"google.golang.org/grpc/metadata"
	"log"
	"root/pb"
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
