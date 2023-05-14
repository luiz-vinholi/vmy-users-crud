package services

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	saltKey string
}

func NewAuth() *Auth {
	return &Auth{saltKey: os.Getenv("JWT_SALT_KEY")}
}

func (a *Auth) GenerateHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (a *Auth) ValidatePassword(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func (a *Auth) GenerateToken(data map[string]interface{}) (token string, err error) {
	claims := jwt.MapClaims{}
	for key, value := range data {
		claims[key] = value
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(a.saltKey))
	return
}

func (a *Auth) ValidateToken(token string) (data map[string]interface{}, isValid bool) {
	claims := jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, a.getTokenKeyFunc())
	if err != nil {
		return nil, false
	}
	if !t.Valid {
		return nil, false
	}
	data = make(map[string]interface{})
	for key, value := range claims {
		data[key] = value
	}
	return data, true
}

func (a *Auth) getTokenKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(a.saltKey), nil
	}
}
