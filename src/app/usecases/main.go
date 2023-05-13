package usecases

import "vmytest/src/infra/repositories"

type AddressData struct {
	Street  string
	City    string
	State   string
	Country string
}

type UserData struct {
	Name      string
	Email     string
	BirthDate string
	Address   *AddressData
}

var usersRepo *repositories.UsersRepository

func Load() {
	usersRepo = repositories.NewUsersRepository()
}

func checkIfUserExists(id string) (bool, error) {
	user, err := usersRepo.GetUser(id)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	} else {
		return true, nil
	}
}
