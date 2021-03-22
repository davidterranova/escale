package inmemory

import (
	"sync"

	"github.com/davidterranova/homework/domain"
	"github.com/google/uuid"
)

type PortRepository struct {
	byID map[string]*domain.Port
	mtx  sync.RWMutex
}

func NewPortRepository() *PortRepository {
	return &PortRepository{
		byID: make(map[string]*domain.Port),
	}
}

func (r *PortRepository) Find(id string) (*domain.Port, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	port, ok := r.byID[id]

	if !ok {
		return nil, domain.ErrPortNotFound
	}
	return port, nil
}

func (r *PortRepository) FindByName(name string) (*domain.Port, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	for _, port := range r.byID {
		if port.Name == name {
			return port, nil
		}
	}

	return nil, domain.ErrPortNotFound
}

func (r *PortRepository) Save(port *domain.Port) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if port.ID == "" {
		uuid := uuid.New().String()
		port.ID = uuid
	}
	r.byID[port.ID] = port

	return nil
}
