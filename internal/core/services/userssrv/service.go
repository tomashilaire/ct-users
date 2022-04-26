package userssrv

import (
	"root/internal/core/ports"
	"root/pkg/uidgen"
)

type service struct {
	ur     ports.UsersRepository
	uidGen uidgen.UIDGen
}

func NewService(ur ports.UsersRepository, uidGen uidgen.UIDGen) *service {
	return &service{ur: ur, uidGen: uidGen}
}
