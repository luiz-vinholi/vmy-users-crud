package middlewares

import (
	"net/http"
	"vmytest/src/app/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationErrResponse struct {
	Field   string      `json:"field"`
	Tag     string      `json:"tag"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

func ErrorHandler(codes map[string]int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.Last()
		if err == nil {
			return
		}
		validationErrs, isValidationErrs := err.Err.(validator.ValidationErrors)
		if isValidationErrs {
			errors := handleValidationErrors(validationErrs)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		customErr, isCustomErr := err.Err.(*errors.CustomError)
		if isCustomErr {
			for key, value := range codes {
				if customErr.Code == key {
					ctx.AbortWithStatusJSON(value, gin.H{"error": customErr})
					return
				}
			}
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}
}

func handleValidationErrors(verrs validator.ValidationErrors) []*ValidationErrResponse {
	var errors []*ValidationErrResponse
	for _, verr := range verrs {
		errorResponse := &ValidationErrResponse{
			Field:   verr.Field(),
			Tag:     verr.Tag(),
			Value:   verr.Value(),
			Message: verr.Error(),
		}
		errors = append(errors, errorResponse)
	}
	return errors
}
