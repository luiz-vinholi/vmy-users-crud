package usecases

import (
	"vmytest/src/infra/models"

	"golang.org/x/crypto/bcrypt"
)

func SetSessionPassword(userId string, pass string) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	data := models.User{Password: string(hash)}
	err = usersRepo.UpdateUser(userId, data)
	return
}
