package security

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
)

type Security interface {
	EncryptPassword(password string) (string, error)
	VerifyPassword(hashed, password string) error

	NewToken(userId string) (string, error)
	NewTokenPayload(tokenString string) (*TokenPayload, error)
}

type security struct {
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
}

func NewSecurity() Security {
	signBytes := []byte(os.Getenv("JWT_PRIV_KEY"))

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}
	verifyBytes := []byte(os.Getenv("JWT_PUB_KEY"))

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}
	return &security{
		verifyKey: verifyKey,
		signKey:   signKey,
	}
}
