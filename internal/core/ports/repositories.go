package ports

import (
	"entity/internal/core/domain"
)

type EntityRepository interface {
	Create(t *domain.Entity) error
	Update(t *domain.Entity) error
	ShowAll() (t []*domain.Entity, err error)
	ShowById(id string) (t *domain.Entity, err error)
	Delete(id string) error
}
