package client

import (
	"bitbucket.org/agroproag/am_authentication/client/go/pb"
	"context"
	"log"
	"time"
)

// SignUp calls sign up RPC
func (grpcClient *GrpcClient) SignUp(name string, email string,
	password string, confirmPassword string, userType string) (id string, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := new(pb.SignUpRequest)
	req.Name = name
	req.Email = email
	req.Password = password
	req.ConfirmPassword = confirmPassword
	req.Type = userType
	resp, err := grpcClient.authService.SignUp(ctx, req)
	if err != nil {
		log.Println("cannot sign up, error: ", err)
		return "", err
	}
	return resp.User.Id, nil
}
