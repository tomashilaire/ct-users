package testprotohdl

import (
	"context"
	"test/internal/core/ports"
	"test/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
func (ph *protoHandler) DeleteTest(ctx context.Context, req *pb.DeleteTestRequest) (*pb.DeleteTestResponse, error) {
	err := ph.ts.Delete(req.Id)
	if err != nil {
		return &pb.DeleteTestResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteTestResponse{}, nil
}

// PostTest implements pb.TestServer
func (ph *protoHandler) PostTest(ctx context.Context, req *pb.PostTestRequest) (*pb.PostTestResponse, error) {
	nEntity, err := ph.ts.Create(req.Name, req.Action)
	if err != nil {
		return &pb.PostTestResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.PostTestResponse{Id: nEntity.Id}, nil
}

// PutTest implements pb.TestServer
func (ph *protoHandler) PutTest(ctx context.Context, req *pb.PutTestRequest) (*pb.PutTestResponse, error) {
	uEntity, err := ph.ts.Update(req.Id, req.Name)
	if err != nil {
		return &pb.PutTestResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.PutTestResponse{Id: uEntity.Id}, nil
}
