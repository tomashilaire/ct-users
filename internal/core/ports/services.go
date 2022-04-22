package ports

import (
	"entity/internal/core/domain"
)

type EntityService interface {
	ShowById(id string) (*domain.Entity, error)
	ShowAll() ([]*domain.Entity, error)
	Update(id string, name string, action string) (*domain.Entity, error)
	Create(name string, action string) (*domain.Entity, error)
	Delete(id string) error
}
