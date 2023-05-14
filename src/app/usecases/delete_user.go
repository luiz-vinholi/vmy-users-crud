package usecases

import "github.com/luiz-vinholi/vmy-users-crud/src/app/errors"

func DeleteUser(id string) (err error) {
	isExists, _ := checkIfUserExists(id)
	if !isExists {
		return errors.UserNotFound()
	}
	err = usersRepo.DeleteUser(id)
	return
}
