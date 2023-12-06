package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate(obj any, err error) []string {
	var errBadRequests []string

	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("%s field is %s", fieldErr.Field(), fieldErr.ActualTag())
			errBadRequests = append(errBadRequests, errMsg)
		}
	}

	return errBadRequests
}
