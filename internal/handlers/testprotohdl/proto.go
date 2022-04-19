package testprotohdl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"test/internal/core/ports"
	"test/pb"
)

type protoHandler struct {
	ts ports.TestService
}

func NewProtoHandler(ts ports.TestService) *protoHandler {
	return &protoHandler{ts: ts}
}

// GetTest implements pb.TestServer
func (ph *protoHandler) GetTest(ctx context.Context, req *pb.GetTestRequest) (*pb.GetTestResponse, error) {
	test, err := ph.ts.ShowById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.GetTestResponse{Id: test.Id, Name: test.Name}, nil
}

// GetAllTests implements pb.Test.Server
func (ph *protoHandler) GetAllTests(req *pb.GetAllTestsRequest, stream pb.Test_GetAllTestsServer) error {
	tests, err := ph.ts.ShowAll()
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, test := range tests {
		err := stream.Send(&pb.GetTestResponse{
			Id:   test.Id,
			Name: test.Name,
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

// DeleteTest implements pb.TestServer
func (ph *protoHandler) DeleteTest(context.Context, *pb.DeleteTestRequest) (*pb.DeleteTestResponse, error) {
	panic("unimplemented")
}

// PostTest implements pb.TestServer
func (ph *protoHandler) PostTest(context.Context, *pb.PostTestRequest) (*pb.PostTestResponse, error) {
	panic("unimplemented")
}

// PutTest implements pb.TestServer
func (ph *protoHandler) PutTest(context.Context, *pb.PutTestRequest) (*pb.PutTestResponse, error) {
	panic("unimplemented")
}
