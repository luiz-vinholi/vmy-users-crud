package errors

func InvalidCredentials() error {
	return &CustomError{
		Code:    "invalid-credentials",
		Message: "Cannot create a session because the credentials are invalid",
	}
}
