package utils

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		err := make(map[string]string)

		for _, e := range validationError {
			switch e.Tag() {
			case "gt":
				err[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu"
			case "lt":
				err[e.Field()] = e.Field() + " phải nhỏ hơn giá trị tối đa"
			case "gte":
				err[e.Field()] = e.Field() + " phải lớn hơn hoặc bằng giá trị tối thiểu"
			case "lte":
				err[e.Field()] = e.Field() + " phải nhỏ hơn hoặc bằng giá trị tối đa"
			case "uuid":
				err[e.Field()] = e.Field() + " phải là UUID hợp lệ"
			case "slug":
				err[e.Field()] = e.Field() + " chỉ được chứa chữ thường, số, dấu gạch ngang hoặc dấu chấm"
			case "min":
				err[e.Field()] = fmt.Sprintf("%s phải nhiều hơn %s kí tự", e.Field(), e.Param())
			case "max":
				err[e.Field()] = fmt.Sprintf("%s phải ít hơn %s kí tự", e.Field(), e.Param())
			case "min_int":
				err[e.Field()] = fmt.Sprintf("%s phải có giá trị lớn hơn %s", e.Field(), e.Param())
			case "max_int":
				err[e.Field()] = fmt.Sprintf("%s phải có giá trị nhỏ hơn %s", e.Field(), e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				err[e.Field()] = fmt.Sprintf("%s phải là một trong các giá trị: %s", e.Field(), allowedValues)
			case "required":
				err[e.Field()] = e.Field() + " là bắt buộc"
			case "search":
				err[e.Field()] = e.Field() + " chỉ được chứa chữ thường, in hoa, số và khoảng trắng"
			case "email":
				err[e.Field()] = e.Field() + " phải đúng định dạng là email"
			case "datetime":
				err[e.Field()] = e.Field() + " phải đúng định dạng YYYY-MM-DD"
			case "file_ext":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				err[e.Field()] = fmt.Sprintf("%s chỉ cho phép file có extension: %s", e.Field(), allowedValues)
			}
		}
		return gin.H{"error": err}
	}
	return gin.H{"error": "Yêu cầu không hợp lệ " + err.Error()}
}

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("Failed to get validator engine")
	}

	var slugRegex = regexp.MustCompile(`[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	var searchRegex = regexp.MustCompile(`[a-zA-Z0-9\s]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	})

	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		// Base
		// 10 : Decimal
		// 16 : Hệ thập lục phân Hex VD: "FF": 255
		// 2 : Binary
		minVal, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() >= minVal
	})

	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		// Base
		// 10 : Decimal
		// 16 : Hệ thập lục phân Hex VD: "FF": 255
		// 2 : Binary
		maxVal, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() <= maxVal
	})

	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()

		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}

		allowedExt := strings.Fields(allowedStr)
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")
		for _, allowed := range allowedExt {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}

		return false
	})

	return nil
}
