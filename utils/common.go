package utils

import (
	"fmt"

	"github.com/go-playground/validator"
)

func Validate(obj interface{}) []string {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		var errs = make([]string, 0)
		for _, currErr := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("%s field is %s", currErr.Field(), currErr.ActualTag())
			errs = append(errs, errMsg)
		}
		return errs
	}
	return nil
}
