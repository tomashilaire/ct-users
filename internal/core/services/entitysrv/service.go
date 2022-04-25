package entitysrv

import (
	"entity/internal/core/domain"
	"entity/internal/core/ports"
	"entity/pkg/apperrors"
	"entity/pkg/errors"
	"entity/pkg/uidgen"
)

type service struct {
	tr     ports.EntityRepository
	uidGen uidgen.UIDGen
}

func NewService(tr ports.EntityRepository, uidGen uidgen.UIDGen) *service {
	return &service{tr: tr, uidGen: uidGen}
}

func (s *service) ShowById(id string) (*domain.Entity, error) {
	entity, err := s.tr.SelectById(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return &domain.Entity{}, errors.New(apperrors.NotFound, err, "entity not found", "")
		}
		return &domain.Entity{}, errors.New(apperrors.Internal, err, "get entity from repository has failed", "")
	}
	return entity, nil
}

func (s *service) ShowAll() ([]*domain.Entity, error) {
	entities, err := s.tr.SelectAll()
	if err != nil {
		return []*domain.Entity{}, errors.New(apperrors.Internal, err, "show all entities from repository has failed", "")
	}

	return entities, nil
}

func (s *service) Delete(id string) error {
	err := s.tr.Delete(id)
	if err != nil {
		return errors.New(apperrors.Internal, err, "delete entity from repository has failed", "")
	}
	return nil
}

func (s *service) Update(id string, name string, action string) (*domain.Entity, error) {
	entity, err := s.ShowById(id)
	if err != nil {
		return entity, err
	}
	entity.Name = name
	entity.Action = action
	if err := s.tr.Set(entity); err != nil {
		return &domain.Entity{}, errors.New(apperrors.Internal, err, "Update entity into repository failed", "")
	}

	return entity, nil

}

func (s *service) Create(name string, action string) (*domain.Entity, error) {
	entity := domain.NewEntity(s.uidGen.New(), name, action)

	if err := s.tr.Insert(entity); err != nil {
		return &domain.Entity{}, errors.New(apperrors.Internal, err, "Create entity into repository failed", "")
	}

	return entity, nil
}
