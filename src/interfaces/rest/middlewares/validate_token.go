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
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "You must provide a JWT token in Authorization header"})
			return
		}

		auth := services.NewAuth()
		data, isValid := auth.ValidateToken(token)
		if !isValid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": invalidTokenErr})
			return
		}

		userId, _ := data["id"].(string)
		usersRepo := repositories.NewUsersRepository()
		user, err := usersRepo.GetUser(userId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
			return
		}
		if user == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": invalidTokenErr})
			return
		}
	}
}

func extractTokenFromHeader(ctx *gin.Context) (token string) {
	header := ctx.Request.Header["Authorization"]
	if header == nil {
		return
	}
	token = strings.Split(header[0], " ")[1]
	return
}
