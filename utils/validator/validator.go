package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ValidatorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Help    string `json:"help"`
}

func ValidateReq(req interface{}) interface{} {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return errorMessage(err, req)
	}

	return nil
}

func errorMessage(err error, req interface{}) interface{} {
	errorValidation := make([]ValidatorResponse, 0)
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			field, _ := reflect.TypeOf(req).Elem().FieldByName(err.Field())

			var key string
			for _, source := range []string{"json", "param", "query", "form"} {
				key = field.Tag.Get(source)
				if len(key) > 0 && key != "-" {
					break
				}
			}

			switch err.Tag() {
			case "required", "required_if":
				validation := ValidatorResponse{
					Error:   "BAD_REQUEST",
					Message: key + " is required",
					Detail:  "Please check request data is in accordance to our API docs",
				}
				errorValidation = append(errorValidation, validation)
			}
		}
	}

	return errorValidation
}
