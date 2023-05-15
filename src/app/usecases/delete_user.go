package usecases

import "github.com/luiz-vinholi/vmy-users-crud/src/app/errors"

// The function deletes a user by their ID if they exist, otherwise it returns an
// user not found error.
func DeleteUser(id string) (err error) {
	isExists, _ := checkIfUserExists(id)
	if !isExists {
		return errors.UserNotFound()
	}
	err = usersRepo.DeleteUser(id)
	return
}
