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
			// log.Printf("%s", e.Namespace())
			root := strings.Split(e.Namespace(), ".")[0]

			rawPath := strings.TrimPrefix(e.Namespace(), root+".")

			parts := strings.Split(rawPath, ".")

			for i, part := range parts {
				if strings.Contains("part", "[") {
					idx := strings.Index(part, "[")
					base := camelToSnake(part[:idx]) // => 0 đến trước [
					index := part[idx:]
					parts[i] = base + index
				} else {
					parts[i] = camelToSnake(part)
				}
			}

			fieldPart := strings.Join(parts, ".")

			switch e.Tag() {
			case "gt":
				err[fieldPart] = fmt.Sprintf("%s phải lớn hơn %s", fieldPart, e.Param())
			case "lt":
				err[fieldPart] = fmt.Sprintf("%s phải nhỏ hơn %s", fieldPart, e.Param())
			case "gte":
				err[fieldPart] = fmt.Sprintf("%s phải lớn hơn hoặc bằng %s", fieldPart, e.Param())
			case "lte":
				err[fieldPart] = fmt.Sprintf("%s phải nhỏ hơn hoặc bằng %s", fieldPart, e.Param())
			case "uuid":
				err[fieldPart] = fmt.Sprintf("%s phải là UUID hợp lệ", fieldPart)
			case "slug":
				err[fieldPart] = fmt.Sprintf("%s chỉ được chứa chữ thường, số, dấu gạch ngang hoặc dấu chấm", fieldPart)
			case "min":
				err[fieldPart] = fmt.Sprintf("%s phải nhiều hơn %s kí tự", fieldPart, e.Param())
			case "max":
				err[fieldPart] = fmt.Sprintf("%s phải ít hơn %s kí tự", fieldPart, e.Param())
			case "min_int":
				err[fieldPart] = fmt.Sprintf("%s phải có giá trị lớn hơn %s", fieldPart, e.Param())
			case "max_int":
				err[fieldPart] = fmt.Sprintf("%s phải có giá trị nhỏ hơn %s", fieldPart, e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				err[fieldPart] = fmt.Sprintf("%s phải là một trong các giá trị: %s", fieldPart, allowedValues)
			case "required":
				err[fieldPart] = fmt.Sprintf("%s là bắt buộc", fieldPart)
			case "search":
				err[fieldPart] = fmt.Sprintf("%s chỉ được chứa chữ thường, in hoa, số và khoảng trắng", fieldPart)
			case "email":
				err[fieldPart] = fmt.Sprintf("%s phải đúng định dạng là email", fieldPart)
			case "datetime":
				err[fieldPart] = fmt.Sprintf("%s phải đúng định dạng YYYY-MM-DD", fieldPart)
			case "file_ext":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				err[fieldPart] = fmt.Sprintf("%s chỉ cho phép file có extension: %s", fieldPart, allowedValues)
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
