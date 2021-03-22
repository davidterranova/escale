package repository

import "github.com/davidterranova/homework/domain"

type PortStreamReader interface {
	Read() chan *domain.Port
}

type PortWriter interface {
	Save(port *domain.Port) error
}
