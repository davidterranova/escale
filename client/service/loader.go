package service

import (
	"log"

	"github.com/davidterranova/homework/client/repository"
)

type LoaderService struct {
	reader repository.PortStreamReader
	writer repository.PortWriter

	loaded int
}

func NewLoaderService(reader repository.PortStreamReader, writer repository.PortWriter) *LoaderService {
	return &LoaderService{
		reader: reader,
		writer: writer,
	}
}

func (s *LoaderService) Load() {
	for port := range s.reader.Read() {
		err := s.writer.Save(port)
		if err != nil {
			log.Printf("failed to save port: %s", err)
		}
		s.loaded++
	}
}

func (s *LoaderService) NbLoaded() int {
	return s.loaded
}
