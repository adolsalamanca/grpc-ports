package repository

import (
	"github.com/adolsalamanca/grpc-ports/server/internal/domain/entity"
)

type PortRepository interface {
	StorePorts([]entity.PortInfo) error
	GetPorts() ([]entity.PortInfo, error)
}
