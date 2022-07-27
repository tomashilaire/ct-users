package security

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"root/pkg/apperrors"
	"strings"
	"time"
)

func (s *security) NewToken(userId string) (string, error) {

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    userId,
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(s.signKey)
}

func (s *security) parseJwtCallback(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return s.verifyKey, nil
}

func (s *security) ExtractToken(r *http.Request) (string, error) {
	// Authorization => Bearer Token...
	header := strings.TrimSpace(r.Header.Get("Authorization"))
	splitted := strings.Split(header, " ")
	if len(splitted) != 2 {
		log.Println("error on extract token from header:", header)
		return "", apperrors.ErrInvalidToken
	}
	return splitted[1], nil
}

func (s *security) ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, s.parseJwtCallback)
}

type TokenPayload struct {
	UserId    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (s *security) NewTokenPayload(tokenString string) (*TokenPayload, error) {
	token, err := s.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, apperrors.ErrInvalidToken
	}
	id, _ := claims["iss"].(string)
	createdAt, _ := claims["iat"].(float64)
	expiresAt, _ := claims["exp"].(float64)
	return &TokenPayload{
		UserId:    id,
		CreatedAt: time.Unix(int64(createdAt), 0),
		ExpiresAt: time.Unix(int64(expiresAt), 0),
	}, nil
}
