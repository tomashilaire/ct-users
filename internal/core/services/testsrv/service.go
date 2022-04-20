package testsrv

import (
	"test/internal/core/domain"
	"test/internal/core/ports"
	"test/pkg/apperrors"
	"test/pkg/errors"
	"test/pkg/uidgen"
)

type service struct {
	tr     ports.TestRepository
	uidGen uidgen.UIDGen
}

func NewService(tr ports.TestRepository, uidGen uidgen.UIDGen) *service {
	return &service{tr: tr, uidGen: uidGen}
}

func (s *service) ShowById(id string) (*domain.Test, error) {
	test, err := s.tr.ShowById(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return &domain.Test{}, errors.New(apperrors.NotFound, err, "entity not found", "")
		}
		return &domain.Test{}, errors.New(apperrors.Internal, err, "get entity from repository has failed", "")
	}
	return test, nil
}

func (s *service) ShowAll() ([]*domain.Test, error) {
	tests, err := s.tr.ShowAll()
	if err != nil {
		return []*domain.Test{}, errors.New(apperrors.Internal, err, "show all entities from repository has failed", "")
	}

	return tests, nil
}

func (s *service) Delete(id string) error {
	err := s.tr.Delete(id)
	if err != nil {
		return errors.New(apperrors.Internal, err, "delete entity from repository has failed", "")
	}
	return nil
}

func (s *service) Update(id string, name string) (*domain.Test, error) {
	test, err := s.ShowById(id)
	if err != nil {
		return test, err
	}
	test.Name = name
	if err := s.tr.Update(test); err != nil {
		return &domain.Test{}, errors.New(apperrors.Internal, err, "Create entity into repository failed", "")
	}

	return test, nil

}

func (s *service) Create(name string) (*domain.Test, error) {
	test := domain.NewTest(s.uidGen.New(), name)

	if err := s.tr.Create(test); err != nil {
		return &domain.Test{}, errors.New(apperrors.Internal, err, "Create entity into repository failed", "")
	}

	return test, nil
}
