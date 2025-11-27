package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		err := make(map[string]string)
		for _, e := range validationError {

			switch e.Tag() {
			case "gt":
				err[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu"
			case "uuid":
				err[e.Field()] = e.Field() + " phải là UUID hợp lệ"
			}
		}
		return gin.H{"error": err}
	}
	return gin.H{"error": "Yêu cầu không hợp lệ "}
}
