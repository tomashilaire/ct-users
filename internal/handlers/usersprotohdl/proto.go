package usersprotohdl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"root/internal/core/ports"
	"root/pb"
)

type protoHandler struct {
	us ports.UsersService
}

func NewProtoHandler(us ports.UsersService) *protoHandler {
	return &protoHandler{us: us}
}

// SingUp implements pb.UsersServer
func (ph *protoHandler) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	nUser, err := ph.us.SignUp(req.Name, req.Email, req.Password, req.ConfirmPassword, req.Type)
	if err != nil {
		return &pb.SignUpResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.SignUpResponse{User: &pb.User{
		Id:      nUser.Id,
		Name:    nUser.Name,
		Email:   nUser.Email,
		Created: nUser.Created.String(),
		Updated: nUser.Updated.String(),
	}}, nil
}
