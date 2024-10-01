package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors"`
}

func GetValidationErrorMsg(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldError.Field())
	case "email":
		return fmt.Sprintf("Invalid %s address", fieldError.Field())
	case "min":
		return fmt.Sprintf("Your %s must be have at least %s characters long", fieldError.Field(), fieldError.Param())
	case "gt":
		return fmt.Sprintf("Should be greater than %s", fieldError.Param())
	}
	return "Unknown error"
}

func ValidationErrorsToMap(errors validator.ValidationErrors) map[string]string {
	res := make(map[string]string)
	for _, v := range errors {
		res[strings.ToLower(v.Field())] = GetValidationErrorMsg(v)
	}
	return res
}

func ExecHandlerValidation[T interface{}](ctx *gin.Context, req *T) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validationError validator.ValidationErrors
		errors.As(err, &validationError)
		errorMap := ValidationErrorsToMap(validationError)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errorMap,
		})
		return err
	}

	return nil
}
