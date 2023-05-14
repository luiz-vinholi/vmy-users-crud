package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/luiz-vinholi/vmy-users-crud/src/interfaces/rest/middlewares"
)

var validate *validator.Validate

func setupRouter() *gin.Engine {
	validate = validator.New()
	router := gin.Default()
	router.Use(
		middlewares.ErrorHandler(
			map[string]int{
				"user-not-found":      http.StatusNotFound,
				"email-in-use":        http.StatusBadRequest,
				"invalid-credentials": http.StatusUnauthorized,
			},
		),
	)
	CreateUserRoutes(router)
	CreateSessionRoutes(router)
	return router
}

func Run() *gin.Engine {
	router := setupRouter()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf("Listening on port %s", port)
	router.Run(port)
	return router
}
