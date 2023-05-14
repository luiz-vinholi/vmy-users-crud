package usecases

import (
	"github.com/luiz-vinholi/vmy-users-crud/src/app/errors"
	"github.com/luiz-vinholi/vmy-users-crud/src/domain/entities"
)

func GetUser(id string) (*entities.User, error) {
	user, _ := usersRepo.GetUser(id)
	if user == nil {
		return nil, errors.UserNotFound()
	}

	euser := &entities.User{
		Id:        user.Id.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate,
	}
	if user.Address != nil {
		euser.Address = &entities.Address{
			Street:  user.Address.Street,
			City:    user.Address.City,
			State:   user.Address.State,
			Country: user.Address.Country,
		}
	}
	euser.SetAge()
	return euser, nil
}
