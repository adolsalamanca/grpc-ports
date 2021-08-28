package _interface

import (
	context "context"
	"fmt"
	"log"
	"net"

	"github.com/adolsalamanca/grpc-ports/server/internal/domain/entity"
	"github.com/adolsalamanca/grpc-ports/server/internal/domain/repository"
	"github.com/adolsalamanca/grpc-ports/server/internal/infrastructure/api"
	"github.com/adolsalamanca/grpc-ports/server/pkg"
	"google.golang.org/grpc"
)

const (
	SuccessCreateMessage = "Successfully added ports to the list"
	InternalServerError  = pkg.Error("Port storage error")
	NotFoundError        = pkg.Error("Port not found error")
)

type Server struct {
	gRpcServ       *grpc.Server
	portRepository repository.PortRepository
}

func NewServer(gRpcServ *grpc.Server, portRepository repository.PortRepository) *Server {
	return &Server{
		gRpcServ:       gRpcServ,
		portRepository: portRepository,
	}
}

func (s *Server) Serve(port uint) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error while trying to listen on tcp connections, %s", err)
	}
	log.Printf("Server listening in port %d", port)
	api.RegisterPortServiceServer(s.gRpcServ, s)

	err = s.gRpcServ.Serve(l)
	if err != nil {
		log.Fatalf("could not accept incomming connections, %s", err)
	}
}

func (s *Server) StorePorts(ctx context.Context, input *api.MultiplePortsInput) (*api.AddPortEntryResponse, error) {
	var ports []entity.PortInfo

	for _, p := range input.Ports {
		ports = append(ports, entity.PortInfo{
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
	err := s.portRepository.StorePorts(ports)
	if err != nil {
		return nil, InternalServerError
	}

	return &api.AddPortEntryResponse{Message: SuccessCreateMessage}, nil
}

func (s *Server) GetAllPorts(ctx context.Context, params *api.EmptyParams) (*api.GetAllPortsResponse, error) {
	r := &api.GetAllPortsResponse{}

	ports, err := s.portRepository.GetPorts()
	if err != nil {
		return nil, NotFoundError
	}

	for _, p := range ports {
		r.Ports = append(r.Ports, &api.Port{
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

	return r, nil
}
