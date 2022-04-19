package ports

import (
	"test/internal/core/domain"
)

type TestService interface {
	ShowById(Id string) (domain.Test, error)
	ShowAll() ([]*domain.Test, error)
	Update(*domain.Test) (domain.Test, error)
	Create(*domain.Test) (domain.Test, error)
	Delete(Id string) error
}
