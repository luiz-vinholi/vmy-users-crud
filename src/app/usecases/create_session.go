package usecases

import (
	"vmytest/src/app/errors"
	"vmytest/src/infra/services"

	"golang.org/x/crypto/bcrypt"
)

type Session struct {
	Email    string
	Password string
}

func CreateSession(session Session) (token string, err error) {
	user, err := usersRepo.GetUserByEmail(session.Email)
	if err != nil {
		return
	}
	if user == nil {
		err = errors.InvalidCredentials()
		return
	}
	if user.Password == "" {
		err = errors.InvalidCredentials()
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(session.Password)); err != nil {
		err = errors.InvalidCredentials()
		return
	}
	auth := services.NewAuth()
	payload := map[string]interface{}{"id": user.Id.Hex()}
	token, err = auth.GenerateToken(payload)
	return
}
