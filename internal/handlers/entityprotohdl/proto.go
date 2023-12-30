package entityprotohdl

import (
	"context"
	"root/internal/core/ports"
	"root/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type protoHandler struct {
	ts ports.EntityService
}

func NewProtoHandler(ts ports.EntityService) *protoHandler {
	return &protoHandler{ts: ts}
}

// GetEntity implements pb.EntityServer
func (ph *protoHandler) GetEntity(ctx context.Context, req *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {
	entity, err := ph.ts.ShowById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.GetEntityResponse{Id: entity.Id, Name: entity.Name}, nil
}

// GetAllEntities implements pb.Entity.Server
func (ph *protoHandler) GetAllEntities(req *pb.GetAllEntitiesRequest, stream pb.Entity_GetAllEntitiesServer) error {
	entities, err := ph.ts.ShowAll()
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, entity := range entities {
		err := stream.Send(&pb.GetEntityResponse{
			Id:   entity.Id,
			Name: entity.Name,
		})
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

// DeleteEntity implements pb.EntityServer
func (ph *protoHandler) DeleteEntity(ctx context.Context, req *pb.DeleteEntityRequest) (*pb.DeleteEntityResponse, error) {
	err := ph.ts.Delete(req.Id)
	if err != nil {
		return &pb.DeleteEntityResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.DeleteEntityResponse{}, nil
}

// PostEntity implements pb.EntityServer
func (ph *protoHandler) PostEntity(ctx context.Context, req *pb.PostEntityRequest) (*pb.PostEntityResponse, error) {
	nEntity, err := ph.ts.Create(req.Name, req.Action)
	if err != nil {
		return &pb.PostEntityResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.PostEntityResponse{Id: nEntity.Id}, nil
}

// PutEntity implements pb.EntityServer
func (ph *protoHandler) PutEntity(ctx context.Context, req *pb.PutEntityRequest) (*pb.PutEntityResponse, error) {
	uEntity, err := ph.ts.Update(req.Id, req.Name, req.Action)
	if err != nil {
		return &pb.PutEntityResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.PutEntityResponse{Id: uEntity.Id}, nil
}
