package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vmytest/src/interfaces/rest/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Run() {
	router := gin.Default()
	router.Use(
		middlewares.ErrorHandler(
			map[string]int{
				"user-not-found":              http.StatusNotFound,
				"email-in-use":                http.StatusBadRequest,
				"invalid-credentials":         http.StatusUnauthorized,
				"invalid-authorization-token": http.StatusUnauthorized,
			},
		),
	)
	validate = validator.New()
	CreateUserRoutes(router)
	CreateSessionsRoutes(router)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf("Listening on port %s", port)
	router.Run(port)
}
