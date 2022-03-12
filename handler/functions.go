package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateErrorResponse(err string) gin.H {
	return gin.H{"errors": map[string]string{"code": "parse_link_error"}}
}

func GetValidationErrorResponse(validateError error) gin.H {
	errors := map[string]string{}

	if validateError != nil {
		if _, ok := validateError.(*validator.InvalidValidationError); ok {
			fmt.Println(validateError)
			return CreateErrorResponse(validateError.Error())
		}

		for _, vErr := range validateError.(validator.ValidationErrors) {
			errors[vErr.Field()] = vErr.Error()
		}

		return gin.H{"errors": errors}
	}

	return nil
}
