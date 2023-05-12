package usecases

import "vmytest/src/app/errors"

func DeleteUser(id string) (err error) {
	isExists, _ := checkIfUserExists(id)
	if !isExists {
		return errors.UserNotFound()
	}
	err = usersRepo.DeleteUser(id)
	return
}
