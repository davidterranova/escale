package adapters

import (
	"context"

	v1 "github.com/davidterranova/homework/api/protobuf/v1"
	"github.com/davidterranova/homework/domain"
	"github.com/davidterranova/homework/store/repository"
)

// GRPCServer implements the server interface defined by our proto service
type GRPCServer struct {
	portRepository repository.PortRepository
	v1.UnimplementedPortServiceServer
}

func NewGRPCServer(portRepository repository.PortRepository) *GRPCServer {
	return &GRPCServer{
		portRepository: portRepository,
	}
}

func (s GRPCServer) Save(ctx context.Context, p *v1.Port) (*v1.EmptyResponse, error) {
	err := s.portRepository.Save(fromGrpc(p))
	if err != nil {
		return nil, err
	}
	return &v1.EmptyResponse{}, nil
}

func (s GRPCServer) Search(ctx context.Context, query *v1.PortQuery) (*v1.PortList, error) {
	port, err := s.portRepository.FindByName(query.Name)
	if err != nil {
		return nil, err
	}

	portsList := &v1.PortList{
		Ports: []*v1.Port{toGrpc(port)},
	}
	return portsList, nil
}

func fromGrpc(p *v1.Port) *domain.Port {
	return &domain.Port{
		ID:       p.Id,
		Name:     p.Name,
		City:     p.City,
		Country:  p.Country,
		Alias:    p.Alias,
		Regions:  p.Regions,
		Province: p.Province,
		Coords:   p.Coordinates,
		Timezone: p.Timezone,
		Unlocs:   p.Unlocs,
		Code:     p.Code,
	}
}

func toGrpc(p *domain.Port) *v1.Port {
	return &v1.Port{
		Id:          p.ID,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Province:    p.Province,
		Coordinates: p.Coords,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}
