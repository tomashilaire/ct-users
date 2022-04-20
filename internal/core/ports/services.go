package ports

import (
	"test/internal/core/domain"
)

type TestService interface {
	ShowById(id string) (*domain.Test, error)
	ShowAll() ([]*domain.Test, error)
	Update(id string, name string) (*domain.Test, error)
	Create(name string, action string) (*domain.Test, error)
	Delete(id string) error
}
