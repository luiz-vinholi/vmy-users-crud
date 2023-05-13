package services

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	saltKey string
}

func NewAuth() *Auth {
	return &Auth{saltKey: os.Getenv("JWT_SALT_KEY")}
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
	t, err := jwt.ParseWithClaims(token, claims, a.getKeyFunc())
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

func (a *Auth) getKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(a.saltKey), nil
	}
}
