package filesrv

import (
	"bytes"
	"fmt"
	"test/internal/core/domain"
	"test/internal/core/ports"
	"test/pkg/apperrors"
	"test/pkg/errors"
	"test/pkg/uidgen"
)

type service struct {
	fr     ports.FileRepository
	uidGen uidgen.UIDGen
}

func NewService(fr ports.FileRepository, uidGen uidgen.UIDGen) *service {
	return &service{fr: fr, uidGen: uidGen}
}

func (s *service) Upload(filePath string, fileType string, fileData bytes.Buffer) (string, error) {
	file := domain.NewFile(filePath, s.uidGen.New(), fileType, fileData)
	sFile, err := s.fr.Save(file)
	if err != nil {
		return "", errors.LogError(errors.New(apperrors.Internal,
			err, "Saving file into repository failed", ""))
	}
	return sFile.Id, nil
}

func (s *service) Download(filePath string, id string, fileType string) (*domain.FileInfo, error) {
	file, err := s.fr.Load(fmt.Sprintf("%s%s%s", filePath, id, fileType))
	if err != nil {
		return &domain.FileInfo{}, errors.LogError(errors.New(apperrors.Internal,
			err, "Downloading file from repository failed", ""))
	}
	return &domain.FileInfo{
		Id:       id,
		FileData: *file,
		Path:     filePath,
		Type:     fileType,
	}, nil
}
