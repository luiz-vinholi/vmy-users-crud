package usecases

import (
	"github.com/luiz-vinholi/vmy-users-crud/src/app/errors"
	"github.com/luiz-vinholi/vmy-users-crud/src/infra/models"
)

func CreateUser(userData UserData) (id string, err error) {
	isExists, err := checkIfEmailExists(userData.Email)
	if err != nil {
		return
	}
	if isExists {
		err = errors.EmailInUse()
		return
	}

	payload := models.User{
		Name:      userData.Name,
		Email:     userData.Email,
		BirthDate: userData.BirthDate,
		Address: &models.Address{
			Street:  userData.Address.Street,
			City:    userData.Address.City,
			State:   userData.Address.State,
			Country: userData.Address.Country,
		},
	}
	id, err = usersRepo.CreateUser(payload)
	return
}

func checkIfEmailExists(email string) (bool, error) {
	user, err := usersRepo.GetUserByEmail(email)
	if err != nil {
		return true, err
	}
	if user == nil {
		return false, nil
	} else {
		return true, nil
	}
}
