package services

import (
	"os"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthValidatePassword(t *testing.T) {
	assert := assert.New(t)

	pass := "Han Solo"
	hash := generateHash(pass)

	auth := NewAuth()
	isValid := auth.ValidatePassword(pass, hash)

	assert.True(isValid)
}

func TestAuthValidatePasswordInvalid(t *testing.T) {
	assert := assert.New(t)

	pass := "Han Solo"
	hash := "Leia Organa"

	auth := NewAuth()
	isValid := auth.ValidatePassword(pass, hash)

	assert.False(isValid)
}

func TestAuthGenerateHash(t *testing.T) {
	assert := assert.New(t)

	pass := "Qui-gon Jinn"
	auth := NewAuth()
	hash, err := auth.GenerateHash(pass)

	isValid := validatePass(pass, hash)

	assert.Nil(err)
	assert.True(isValid)
}

func TestAuthGenerateHashInvalid(t *testing.T) {
	assert := assert.New(t)

	pass := strings.Repeat("X", 73)
	auth := NewAuth()
	hash, err := auth.GenerateHash(pass)

	assert.NotNil(err)
	assert.Equal("", hash)
}

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

func generateHash(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash)
}

func validatePass(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
