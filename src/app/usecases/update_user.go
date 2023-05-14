package usecases

import (
	"github.com/luiz-vinholi/vmy-users-crud/src/app/errors"
	"github.com/luiz-vinholi/vmy-users-crud/src/infra/models"
)

func UpdateUser(id string, userData UserData) (err error) {
	isExists, _ := checkIfUserExists(id)
	if !isExists {
		err = errors.UserNotFound()
		return
	}

	payload := models.User{
		Name:      userData.Name,
		Email:     userData.Email,
		BirthDate: userData.BirthDate,
	}
	if userData.Address != nil {
		payload.Address = &models.Address{
			Street:  userData.Address.Street,
			City:    userData.Address.City,
			State:   userData.Address.State,
			Country: userData.Address.Country,
		}
	}
	err = usersRepo.UpdateUser(id, payload)
	return
}
