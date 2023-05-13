package rest

import (
	"net/http"
	"vmytest/src/app/usecases"
	"vmytest/src/interfaces/rest/middlewares"

	"github.com/gin-gonic/gin"
)

type CreateAdressData struct {
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	Country string `json:"country"`
}

type CreateUserData struct {
	Name      string            `json:"name" validate:"required,min=2"`
	Email     string            `json:"email" validate:"required,email"`
	BirthDate string            `json:"birthDate" validate:"required,datetime=2006-01-02"`
	Address   *CreateAdressData `json:"address" validate:"required,dive"`
}

type UpdateAddressData struct {
	Street  string `json:"street" validate:"omitempty"`
	City    string `json:"city" validate:"omitempty"`
	State   string `json:"state" validate:"omitempty"`
	Country string `json:"country" validate:"omitempty"`
}

type UpdateUserData struct {
	Name      string             `json:"name" validate:"omitempty,min=2"`
	Email     string             `json:"email,omitempty" validate:"omitempty,email"`
	BirthDate string             `json:"birthDate" validate:"omitempty,datetime=2006-01-02"`
	Address   *UpdateAddressData `json:"address" validate:"omitempty,dive"`
}

func CreateUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users", middlewares.ValidateToken())

	userRouter.GET("/", func(ctx *gin.Context) {
		users, err := usecases.GetUsers()
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, users)
	})

	userRouter.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		user, err := usecases.GetUser(id)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

	userRouter.POST("/", func(ctx *gin.Context) {
		var body CreateUserData
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.Error(err)
			return
		}

		if err := validate.Struct(body); err != nil {
			ctx.Error(err)
			return
		}

		addressData := &usecases.AddressData{
			Street:  body.Address.Street,
			City:    body.Address.City,
			State:   body.Address.State,
			Country: body.Address.Country,
		}
		userData := usecases.UserData{
			Name:      body.Name,
			Email:     body.Email,
			BirthDate: body.BirthDate,
			Address:   addressData,
		}
		id, err := usecases.CreateUser(userData)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	})

	userRouter.PATCH("/:id", func(ctx *gin.Context) {
		var body UpdateUserData
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.Error(err)
			return
		}

		if err := validate.Struct(body); err != nil {
			ctx.Error(err)
			return
		}

		id := ctx.Param("id")
		userData := usecases.UserData{
			Name:      body.Name,
			Email:     body.Email,
			BirthDate: body.BirthDate,
		}
		if body.Address != nil {
			userData.Address = &usecases.AddressData{
				Street:  body.Address.Street,
				City:    body.Address.City,
				State:   body.Address.State,
				Country: body.Address.Country,
			}
		}
		err := usecases.UpdateUser(id, userData)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{})
	})

	userRouter.DELETE("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := usecases.DeleteUser(id)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{})
	})
}
