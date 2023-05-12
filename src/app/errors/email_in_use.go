package errors

func EmailInUse() error {
	return &CustomError{
		Code:    "email-in-use",
		Message: "This email is already in use",
	}
}
