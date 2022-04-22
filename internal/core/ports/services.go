package ports

import (
	"bytes"
	"test/internal/core/domain"
)

type TestService interface {
	ShowById(id string) (*domain.Test, error)
	ShowAll() ([]*domain.Test, error)
	Update(id string, name string) (*domain.Test, error)
	Create(name string, action string) (*domain.Test, error)
	Delete(id string) error
}

type FilesService interface {
	Upload(filePath string, fileType string, fileData bytes.Buffer) (string, error)
	Download(filePath string, id string, fileType string) (*domain.FileInfo, error)
}
