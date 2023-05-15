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

// This function is generating a hash from a given password using the bcrypt algorithm. It takes a
// string password as input and returns a string hash and an error. If the hash generation is
// successful, the function returns the hash as a string. If there is an error, it returns an empty
// string and the error.
func (a *Auth) GenerateHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// This function is validating a given password against a hash using the bcrypt algorithm. It takes a
// string password and a string hash as input and returns a boolean value indicating whether the
// password matches the hash or not. If the password matches the hash, the function returns true,
// otherwise it returns false.
func (a *Auth) ValidatePassword(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

// This function is generating a JWT (JSON Web Token) using the HMAC-SHA256 algorithm. It signs the token using
// the salt key stored in the Auth struct and returns the signed token and any error that occurred
// during the process.
func (a *Auth) GenerateToken(data map[string]interface{}) (token string, err error) {
	claims := jwt.MapClaims{}
	for key, value := range data {
		claims[key] = value
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(a.saltKey))
	return
}

// This function is validating a JWT (JSON Web Token) using the HMAC-SHA256 algorithm. It takes a
// string token as input and returns a map of string keys and interface{} values representing the data
// included in the token's payload, as well as a boolean value indicating whether the token is valid or
// not.
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
