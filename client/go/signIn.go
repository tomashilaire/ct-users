package client

import (
	"context"
	"github.com/tomashilaire/ct-users/client/go/pb"
	"log"
	"time"
)

// SignIn calls sign up RPC
func (grpcClient *GrpcClient) SignIn(email string, password string) (user *pb.User, token string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := new(pb.SignInRequest)
	req.Email = email
	req.Password = password
	resp, err := grpcClient.authService.SignIn(ctx, req)
	if err != nil {
		log.Println("cannot sign in, error: ", err)
		return user, "", err
	}
	return resp.User, resp.Token, nil
}
