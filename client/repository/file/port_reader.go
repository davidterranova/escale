package file

import (
	"encoding/json"
	"io"
	"log"

	"github.com/davidterranova/homework/domain"
)

type JSONPortsLoader struct {
	dec *json.Decoder
}

type portData map[string]domain.Port

func NewJSONPortLoader(reader io.Reader) *JSONPortsLoader {
	return &JSONPortsLoader{
		dec: json.NewDecoder(reader),
	}
}

// TODO : handle errors properly
func (r JSONPortsLoader) Read() chan *domain.Port {
	portChan := make(chan *domain.Port)

	go func(ch chan *domain.Port) {
		defer func() {
			close(ch)
		}()
		/*_, err := r.dec.Token()
		if err != nil {
			log.Fatalf("failed to read token: %s", err)
		}*/

		for r.dec.More() {
			var pd portData
			err := r.dec.Decode(&pd)
			if err != nil {
				log.Fatalf("failed to read port: %s", err)
			}

			for _, p := range pd {
				ch <- &p
			}
		}

		/*_, err = r.dec.Token()
		if err != nil {
			log.Fatalf("failed to read token: %s", err)
		}*/
	}(portChan)

	return portChan
}
