package adapters

import (
	"testing"

	"github.com/davidterranova/homework/domain"
	"github.com/google/uuid"
	"github.com/matryer/is"
)

func TestFromToGrpc(t *testing.T) {
	is := is.New(t)

	initial := &domain.Port{
		ID:      uuid.New().String(),
		Name:    "Ajman",
		City:    "Ajman",
		Country: "United Arab Emirates",
		Alias: []string{
			"some_aliad",
		},
		Regions: []string{},
		Coords: []float32{
			55.5136433,
			25.4052165,
		},
		Province: "Ajman",
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AEAJM",
		},
		Code: "52000",
	}

	copy := fromGrpc(toGrpc(initial))
	is.Equal(*initial, *copy)
}
