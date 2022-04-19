package ports

import (
	"test/internal/core/domain"
)

type TestRepository interface {
	Create(t *domain.Test) error
	Update(t *domain.Test) error
	ShowAll() (t []*domain.Test, err error)
	ShowById(id string) (t *domain.Test, err error)
	Delete(id string) error
}
