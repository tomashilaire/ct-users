package usersprotohdl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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

// SignUp implements pb.AuthenticationServer
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

// SignIn implements pb.AuthenticationServer
func (ph *protoHandler) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	lUser, token, err := ph.us.SignIn(req.Email, req.Password)
	if err != nil {
		return &pb.SignInResponse{}, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.SignInResponse{
		User: &pb.User{
			Id:      lUser.Id,
			Name:    lUser.Name,
			Email:   lUser.Email,
			Created: lUser.Created.String(),
			Updated: lUser.Updated.String(),
		},
		Token: token,
	}, nil
}

// Authenticate implements pb.AuthenticationServer
func (ph *protoHandler) Authenticate(ctx context.Context,
	req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.AuthenticateResponse{}, status.Errorf(codes.Internal, "ERROR parsing request")
	}
	authId := md["authid"][0]
	return &pb.AuthenticateResponse{AuthId: authId}, nil
}
