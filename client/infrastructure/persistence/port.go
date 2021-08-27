package persistence

import (
	"context"
	"fmt"
	"log"

	"github.com/adolsalamanca/grpc-ports/client/domain/entity"
	"github.com/adolsalamanca/grpc-ports/client/infrastructure/api"
	"github.com/adolsalamanca/grpc-ports/client/infrastructure/common"
	"google.golang.org/grpc"
)

const (
	PersistenceError = common.Error("there was an error retrieving data from persistence")
)

type PortgRpcPersistence struct {
	serviceClient api.PortServiceClient
}

func NewPortgRpcPersistence(host string, port uint) *PortgRpcPersistence {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not dial server, %s", err)
	}

	return &PortgRpcPersistence{
		serviceClient: api.NewPortServiceClient(conn),
	}
}

func (p PortgRpcPersistence) StorePorts(ports []entity.PortInfo) error {
	var allPorts []*api.Port
	for _, p := range ports {
		allPorts = append(allPorts, &api.Port{
			Name:        p.Name,
			City:        p.City,
			Country:     p.Country,
			Alias:       p.Alias,
			Regions:     p.Regions,
			Coordinates: p.Coordinates,
			Province:    *p.Province,
			Timezone:    *p.Timezone,
			Unlocs:      p.Unlocs,
			Code:        *p.Code,
		})
	}
	p.serviceClient.StorePorts(context.Background(), &api.MultiplePortsInput{
		Ports: allPorts,
	})

	return nil
}

func (p PortgRpcPersistence) GetAllPorts() ([]entity.PortInfo, error) {
	ports, err := p.serviceClient.GetAllPorts(context.Background(), &api.EmptyParams{})
	if err != nil {
		log.Printf("Error getting list, %s", err)
		return nil, PersistenceError
	}

	var allPorts []entity.PortInfo
	for _, p := range ports.Ports {
		allPorts = append(allPorts, entity.PortInfo{
			Name:        p.Name,
			City:        p.City,
			Country:     p.Country,
			Alias:       p.Alias,
			Regions:     p.Regions,
			Coordinates: p.Coordinates,
			Province:    &p.Province,
			Timezone:    &p.Timezone,
			Unlocs:      p.Unlocs,
			Code:        &p.Code,
		})
	}

	return allPorts, nil
}
