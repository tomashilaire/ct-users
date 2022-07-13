package userssrv

import (
	"log"
	"root/internal/core/domain"
	"root/internal/core/ports"
	"root/pkg/apperrors"
	"root/pkg/errors"
	"root/pkg/security"
	"root/pkg/uidgen"
	"root/pkg/validators"
	"time"
)

type service struct {
	ur     ports.UsersRepository
	uidGen uidgen.UIDGen
	v      validators.Validators
	sec    security.Security
}

func NewService(ur ports.UsersRepository, uidGen uidgen.UIDGen,
	v validators.Validators, sec security.Security) *service {
	return &service{ur: ur, uidGen: uidGen, v: v, sec: sec}
}

func (s *service) SignUp(name string, lastName string, email string, password string,
	confirmPassword string, userType string) (*domain.User, error) {
	err := s.v.ValidateSingUp(name, email, password,
		confirmPassword, userType)
	if err != nil {
		return &domain.User{}, errors.LogError(errors.New(apperrors.InvalidInput,
			err, "Invalid sing up data", ""))
	}

	password, err = s.sec.EncryptPassword(password)
	if err != nil {
		return &domain.User{}, errors.LogError(errors.New(apperrors.Internal,
			err, "Error processing password", ""))
	}

	email = s.v.NormalizeEmail(email)

	_, err = s.ur.GetByEmail(email)
	if err == apperrors.NotFound {
		nUser := domain.NewUser(
			s.uidGen.New(),
			name,
			lastName,
			email,
			userType,
			password,
			time.Now().UTC(),
		)
		err = s.ur.Save(nUser)
		if err != nil {
			return &domain.User{}, errors.LogError(errors.New(apperrors.Internal,
				err, "Error saving user", ""))
		}
		return nUser, nil

	}
	return &domain.User{}, errors.LogError(errors.New(apperrors.InvalidInput,
		nil, "User already exists", ""))

}

func (s *service) SignIn(email string, password string) (u *domain.User, token string, err error) {
	email = s.v.NormalizeEmail(email)

	usr, err := s.ur.GetByEmail(email)
	if err == apperrors.NotFound {
		return &domain.User{}, "", errors.LogError(errors.New(apperrors.InvalidInput,
			err, "Wrong email or password", ""))
	} else if err != nil {
		return &domain.User{}, "", errors.LogError(errors.New(apperrors.Internal,
			err, "Error signin", ""))
	}
	err = s.sec.VerifyPassword(usr.Password, password)
	if err != nil {
		return &domain.User{}, "", errors.LogError(errors.New(apperrors.InvalidInput,
			err, "Wrong email or password", ""))

	}

	token, err = s.sec.NewToken(usr.Id)
	if err != nil {
		log.Println("signin failed:", err)
		return &domain.User{}, "", errors.LogError(errors.New(apperrors.ErrInvalidToken,
			err, "signin failed", ""))
	}

	return usr, token, nil
}
