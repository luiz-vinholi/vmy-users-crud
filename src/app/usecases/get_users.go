package usecases

import (
	"vmytest/src/domain/entities"
)

func GetUsers() ([]entities.User, error) {
	users, err := usersRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	var eusers []entities.User
	for _, user := range users {
		euser := entities.User{
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
		eusers = append(eusers, euser)
	}
	return eusers, nil
}
