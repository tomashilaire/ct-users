package ports

import (
	"bytes"
	"root/internal/core/domain"
)

type EntityService interface {
	ShowById(id string) (*domain.Entity, error)
	ShowAll() ([]*domain.Entity, error)
	Update(id string, name string, action string) (*domain.Entity, error)
	Create(name string, action string) (*domain.Entity, error)
	Delete(id string) error
}

type FilesService interface {
	Upload(filePath string, fileType string, fileData bytes.Buffer) (string, error)
	Download(filePath string, id string, fileType string) (*domain.FileInfo, error)
}
