package errors

func UserNotFound() error {
	return &CustomError{
		Code:    "user-not-found",
		Message: "User not found",
	}
}
