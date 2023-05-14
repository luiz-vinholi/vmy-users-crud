package usecases

import (
	"github.com/luiz-vinholi/vmy-users-crud/src/app/errors"
	"github.com/luiz-vinholi/vmy-users-crud/src/infra/models"
	"github.com/luiz-vinholi/vmy-users-crud/src/infra/services"
)

func SetSessionPassword(userId string, pass string) (err error) {
	isExists, _ := checkIfUserExists(userId)
	if !isExists {
		err = errors.UserNotFound()
		return
	}
	auth := services.NewAuth()
	hash, err := auth.GenerateHash(pass)
	if err != nil {
		return
	}
	data := models.User{Password: string(hash)}
	err = usersRepo.UpdateUser(userId, data)
	return
}
