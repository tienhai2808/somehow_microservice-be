package common

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func HandleValidationError(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				return fmt.Sprintf("%s là bắt buộc", e.Field())
			case "email":
				return fmt.Sprintf("%s không phải là email hợp lệ", e.Field())
			case "min":
				return fmt.Sprintf("%s phải có ít nhất %s ký tự", e.Field(), e.Param())
			case "max":
				return fmt.Sprintf("%s không được vượt quá %s ký tự", e.Field(), e.Param())
			case "len":
				return fmt.Sprintf("%s phải có chính xác %s ký tự", e.Field(), e.Param())
			case "numeric":
				return fmt.Sprintf("%s phải là số", e.Field())
			case "uuid4":
				return fmt.Sprintf("%s phải là UUID phiên bản 4 hợp lệ", e.Field())
			case "oneof":
				return fmt.Sprintf("%s phải có giá trị là: %s", e.Field(), e.Param())
			default:
				return fmt.Sprintf("%s không hợp lệ", e.Field())
			}
		}
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return fmt.Sprintf("%s phải là kiểu %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String())
	}

	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		return fmt.Sprintf("JSON không hợp lệ tại byte %d", syntaxError.Offset)
	}

	if err != nil {
		return err.Error()
	}

	return "Dữ liệu không hợp lệ"
}
