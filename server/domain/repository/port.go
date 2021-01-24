package repository

import "github.com/adolsalamanca/ports/server/domain/entity"

type PortRepository interface {
	StorePorts([]entity.PortInfo) error
	GetPorts() ([]entity.PortInfo, error)
}
