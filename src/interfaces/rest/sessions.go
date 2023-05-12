package rest

import "github.com/gin-gonic/gin"

func CreateSessionsRoutes(router *gin.Engine) {
	sessionRouter := router.Group("/sessions")

	sessionRouter.POST("/", func(ctx *gin.Context) {

	})
}
