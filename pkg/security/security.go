package security

type Security interface {
	EncryptPassword(password string) (string, error)
	VerifyPassword(hashed, password string) error

	NewToken(userId string) (string, error)
}

type security struct{}

func NewSecurity() Security {
	return &security{}
}
