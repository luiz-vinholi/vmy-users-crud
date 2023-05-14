package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luiz-vinholi/vmy-users-crud/src/app/usecases"
)

type Session struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}

type Password struct {
	Password string `json:"password" validate:"required,min=6,max=72"`
}

func CreateSessionRoutes(router *gin.Engine) {
	sessionRouter := router.Group("/sessions")
	sessionRouter.POST("/", createSession)
	sessionRouter.PUT("/users/:id/passwords", setSessionUserPassword)
}

func createSession(ctx *gin.Context) {
	var body Session
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(err)
		return
	}

	if err := validate.Struct(body); err != nil {
		ctx.Error(err)
		return
	}

	session := usecases.Session{
		Email:    body.Email,
		Password: body.Password,
	}
	token, err := usecases.CreateSession(session)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func setSessionUserPassword(ctx *gin.Context) {
	var body Password
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(err)
		return
	}

	if err := validate.Struct(body); err != nil {
		ctx.Error(err)
		return
	}

	userId := ctx.Param("id")
	pass := body.Password
	if err := usecases.SetSessionPassword(userId, pass); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
