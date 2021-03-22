package inmemory

import (
	"testing"

	"github.com/davidterranova/homework/domain"
	fuzz "github.com/google/gofuzz"
	"github.com/google/uuid"
	"github.com/matryer/is"
)

// TODO : add concurrent saving test
func TestSavePortRepository(t *testing.T) {
	is := is.New(t)

	repo := NewPortRepository()
	ports := make([]*domain.Port, 50)
	for i := 0; i < 50; i++ {
		ports[i] = portFuzzer(t)
		err := repo.Save(ports[i])
		is.NoErr(err)
	}

	for i := 0; i < 50; i++ {
		port, ok := repo.byID[ports[i].ID]
		is.True(ok)
		is.Equal(port, ports[i])
	}
}

func TestPortRepositoryReader(t *testing.T) {
	is := is.New(t)

	ports := make(map[string]*domain.Port)
	for i := 0; i < 50; i++ {
		p := portFuzzer(t)
		ports[p.ID] = p
	}

	t.Run("findByID", func(t *testing.T) {
		repo := NewPortRepository()
		repo.byID = ports

		for id, port := range ports {
			portFromRepo, err := repo.Find(id)
			is.NoErr(err)
			is.Equal(port, portFromRepo)
		}
	})

	t.Run("findByIDWithNonExistingEntry", func(t *testing.T) {
		repo := NewPortRepository()

		port, err := repo.Find("unknown_id")
		is.Equal(err, domain.ErrPortNotFound)
		is.Equal(port, nil)
	})

	t.Run("findByName", func(t *testing.T) {
		repo := NewPortRepository()
		repo.byID = ports

		for _, port := range ports {
			portFromRepo, err := repo.FindByName(port.Name)
			is.NoErr(err)
			is.Equal(port, portFromRepo)
		}
	})

	t.Run("findByNameWithNonExistingEntry", func(t *testing.T) {
		repo := NewPortRepository()

		port, err := repo.FindByName("unknown_name")
		is.Equal(err, domain.ErrPortNotFound)
		is.Equal(port, nil)
	})
}

// TODO : improve port fuzzing with more accurate values
func portFuzzer(t *testing.T) *domain.Port {
	t.Helper()
	f := fuzz.New().NilChance(0)

	var p domain.Port
	f.Fuzz(&p)
	p.ID = uuid.New().String()
	p.Name = "name_" + uuid.New().String()

	return &p
}
