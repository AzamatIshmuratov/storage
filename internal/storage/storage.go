package storage

import (
	"fmt"

	"github.com/AzamatIshmuratov/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error) {
	file, err := file.NewFile(fileName, blob)

	if err != nil {
		return nil, err
	}

	s.files[file.ID] = file

	return file, nil

}

func (s *Storage) GetByID(id uuid.UUID) (*file.File, error) {
	file, ok := s.files[id]

	if !ok {
		return nil, fmt.Errorf("file %s not found in storage", id.String())
	}

	return file, nil
}
