package services

import (
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestAuthGenerateToken(t *testing.T) {
	assert := assert.New(t)
	data := map[string]interface{}{
		"id":   "123",
		"name": "Luke",
	}
	auth := NewAuth()
	token, err := auth.GenerateToken(data)

	assert.NotNil(token)
	assert.Nil(err)
}

func TestAuthValidateToken(t *testing.T) {
	assert := assert.New(t)

	token, _ := generateToken(os.Getenv("JWT_SALT_KEY"))

	auth := NewAuth()
	data, isValid := auth.ValidateToken(token)

	assert.True(isValid)
	assert.Equal(data["id"], "123")
}

func TestAuthInvalidToken(t *testing.T) {
	assert := assert.New(t)

	token, _ := generateToken("Sabini")

	auth := NewAuth()
	data, isValid := auth.ValidateToken(token)

	assert.False(isValid)
	assert.Nil(data)
}

func TestAuthNotAToken(t *testing.T) {
	assert := assert.New(t)

	auth := NewAuth()
	data, isValid := auth.ValidateToken("not-a-token")

	assert.False(isValid)
	assert.Nil(data)
}

func generateToken(key string) (string, error) {
	claims := jwt.MapClaims{"id": "123"}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(key))
	return token, err
}
