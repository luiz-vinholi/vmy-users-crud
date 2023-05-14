package usecases

import (
	"github.com/luiz-vinholi/vmy-users-crud/src/app/errors"
	"github.com/luiz-vinholi/vmy-users-crud/src/infra/services"
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

	passHashed := user.Password
	auth := services.NewAuth()
	if isValid := auth.ValidatePassword(session.Password, passHashed); !isValid {
		err = errors.InvalidCredentials()
		return
	}
	payload := map[string]interface{}{"id": user.Id.Hex()}
	token, err = auth.GenerateToken(payload)
	return
}
