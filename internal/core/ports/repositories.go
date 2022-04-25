package ports

import (
	"bytes"
	"root/internal/core/domain"
)

type EntityRepository interface {
	Insert(t *domain.Entity) error
	Set(t *domain.Entity) error
	SelectAll() (t []*domain.Entity, err error)
	SelectById(id string) (t *domain.Entity, err error)
	Delete(id string) error
}

type FileRepository interface {
	Save(f *domain.FileInfo) (*domain.FileInfo, error)
	Load(p string) (*bytes.Buffer, error)
}
