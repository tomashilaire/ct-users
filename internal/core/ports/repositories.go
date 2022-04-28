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

type UsersRepository interface {
	Save(user *domain.User) error
	GetById(id string) (user *domain.User, err error)
	GetByEmail(email string) (user *domain.User, err error)
	GetAll() (users []*domain.User, err error)
	Update(user *domain.User) error
	Delete(id string) error

	Disconnect()
}
