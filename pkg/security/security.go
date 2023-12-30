package security

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strings"
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
	signBytes := []byte(strings.ReplaceAll(os.Getenv("JWT_PRIV_KEY"), "\\n", "\n"))

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal("Parse sign KEY failed: ", err)
	}
	verifyBytes := []byte(strings.ReplaceAll(os.Getenv("JWT_PUB_KEY"), "\\n", "\n"))
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal("Parse verify KEY failed: ", err)
	}
	return &security{
		verifyKey: verifyKey,
		signKey:   signKey,
	}
}
