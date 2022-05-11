package validate

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func NotEmpty(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

func NotNegative(fl validator.FieldLevel) bool {
	return fl.Field().Float() >= 0
}
