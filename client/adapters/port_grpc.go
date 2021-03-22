package adapters

import (
	"context"

	v1 "github.com/davidterranova/homework/api/protobuf/v1"
	"github.com/davidterranova/homework/domain"
)

type PortGRPCClient struct {
	client v1.PortServiceClient
}

func NewPortGRPCClient(client v1.PortServiceClient) *PortGRPCClient {
	return &PortGRPCClient{
		client: client,
	}
}

func (s PortGRPCClient) Save(port *domain.Port) error {
	_, err := s.client.Save(
		context.Background(),
		&v1.Port{
			Id:          port.ID,
			Name:        port.Name,
			City:        port.City,
			Country:     port.Country,
			Alias:       port.Alias,
			Regions:     port.Regions,
			Province:    port.Province,
			Coordinates: port.Coords,
			Timezone:    port.Timezone,
			Unlocs:      port.Unlocs,
			Code:        port.Code,
		},
	)
	return err
}

func (s PortGRPCClient) Search(ctx context.Context, name string) ([]*domain.Port, error) {
	resp, err := s.client.Search(
		ctx,
		&v1.PortQuery{
			Name: name,
		},
	)

	if err != nil {
		return nil, err
	}

	var ports = make([]*domain.Port, 0, len(resp.Ports))
	for _, p := range resp.Ports {
		ports = append(
			ports,
			&domain.Port{
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
			},
		)
	}

	return ports, nil
}
