package middlewares

import (
	"net/http"
	"strings"
	"vmytest/src/app/errors"
	"vmytest/src/infra/repositories"
	"vmytest/src/infra/services"

	"github.com/gin-gonic/gin"
)

var invalidTokenErr = &errors.CustomError{
	Code:    "invalid-authorization-token",
	Message: "You must provide a valid JWT token in Authorization header",
}

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := extractTokenFromHeader(ctx)
		if token == "" {
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{"error": "You must provide a JWT token in Authorization header"})
			return
		}

		auth := services.NewAuth()
		data, isValid := auth.ValidateToken(token)
		if !isValid {
			ctx.Error(invalidTokenErr)
			return
		}

		userId, _ := data["id"].(string)
		usersRepo := repositories.NewUsersRepository()
		user, err := usersRepo.GetUser(userId)
		if err != nil {
			ctx.Error(err)
			return
		}
		if user == nil {
			ctx.Error(invalidTokenErr)
			return
		}
	}
}

func extractTokenFromHeader(ctx *gin.Context) (token string) {
	header := ctx.Request.Header["Authorization"][0]
	token = strings.Split(header, " ")[1]
	return
}
