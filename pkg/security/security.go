package security

import "golang.org/x/crypto/bcrypt"

type Security interface {
	EncryptPassword(password string) (string, error)
	VerifyPassword(hashed, password string) error
}

type security struct{}

func NewSecurity() Security {
	return &security{}
}

func (s *security) EncryptPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (s *security) VerifyPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
