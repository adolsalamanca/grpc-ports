package repository

import "github.com/adolsalamanca/ports/client/domain/entity"

type PortRepository interface {
	StorePorts([]entity.PortInfo) error
	GetAllPorts() ([]entity.PortInfo, error)
}
