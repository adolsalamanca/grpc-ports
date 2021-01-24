package persistence

import (
	"github.com/adolsalamanca/ports/server/domain/entity"
	"github.com/adolsalamanca/ports/server/infrastructure/common"
)

const (
	EmptyRepositoryErr = common.Error("Empty repository error")
)

type PortsMemoryRepository struct {
	ports map[string]entity.PortInfo
}

// NewPortsMemoryRepository instantiates a new PortRepository in memory to get/store port data
func NewPortsMemoryRepository() *PortsMemoryRepository {
	return &PortsMemoryRepository{
		ports: make(map[string]entity.PortInfo),
	}
}

func (r *PortsMemoryRepository) StorePorts(ports []entity.PortInfo) error {
	//TODO: It would be good to confirm that unlocs always stores the key of the port to be stored.
	for _, p := range ports {
		r.ports[p.Unlocs[0]] = p
	}
	return nil
}

func (r *PortsMemoryRepository) GetPorts() ([]entity.PortInfo, error) {
	if len(r.ports) == 0 {
		return nil, EmptyRepositoryErr
	}

	var allports []entity.PortInfo
	for _,p := range r.ports {
		allports = append(allports, p)
	}

	return allports, nil
}
