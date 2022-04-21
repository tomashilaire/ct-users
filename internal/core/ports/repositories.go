package ports

import (
	"entity/internal/core/domain"
)

type EntityRepository interface {
	Insert(t *domain.Entity) error
	Set(t *domain.Entity) error
	SelectAll() (t []*domain.Entity, err error)
	SelectById(id string) (t *domain.Entity, err error)
	Delete(id string) error
}
