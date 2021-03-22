package repository

import "github.com/davidterranova/homework/domain"

type PortReader interface {
	Find(id string) (*domain.Port, error)
	FindByName(name string) (*domain.Port, error)
}

type PortWriter interface {
	Save(port *domain.Port) error
}

type PortRepository interface {
	PortReader
	PortWriter
}
