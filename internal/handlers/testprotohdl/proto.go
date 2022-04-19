package testprotohdl

import (
	"context"
	"log"
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
func (ph *protoHandler) GetTest(context.Context, *pb.GetTestRequest) (*pb.GetTestResponse, error) {
	panic("unimplemented")
}

// GetAllTests implements pb.Test.Server
func (ph *protoHandler) GetAllTests(req *pb.GetAllTestsRequest, stream pb.Test_GetAllTestsServer) error {
	tests, err := ph.ts.ShowAll()
	if err != nil {
		log.Fatal("Unable to retrieve data", err)
		return err
	}
	log.Println(tests)

	for _, test := range tests {
		err := stream.Send(&pb.GetTestResponse{
			Id:   test.Id,
			Name: test.Name,
		})
		if err != nil {
			log.Fatal("Unable to sent to protobuffer", err)
			return err
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
