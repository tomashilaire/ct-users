package testsrv

import (
	"log"
	"test/internal/core/domain"
	"test/internal/core/ports"
)

type service struct {
	tr ports.TestRepository
}

func NewService(tr ports.TestRepository) *service {
	return &service{tr: tr}
}

func (s *service) ShowById(Id string) (domain.Test, error) {
	panic("unimplemented")
}

func (s *service) ShowAll() ([]*domain.Test, error) {
	tests, err := s.tr.ShowAll()
	if err != nil {
		log.Fatal("Unable to retrieve data", err)
	}
	log.Println(tests)

	return tests, nil
}

func (s *service) Delete(Id string) error {
	panic("unimplemented")
}

func (s *service) Update(t *domain.Test) (domain.Test, error) {
	panic("unimplemented")
}

func (s *service) Create(t *domain.Test) (domain.Test, error) {
	panic("unimplemented")
}
