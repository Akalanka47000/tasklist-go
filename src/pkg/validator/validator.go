package validator

import (
	"strings"
	"unicode"

	v "github.com/go-playground/validator/v10"
	"github.com/samber/lo"
)

type (
	ValidationErrors       = v.ValidationErrors
	FieldError             = v.FieldError
	FieldLevel             = v.FieldLevel
	Validator              = v.Validate
	InvalidValidationError = v.InvalidValidationError
)

func New() *v.Validate {
	validator := v.New()
	lo.Must0(validator.RegisterValidation("password", password))
	lo.Must0(validator.RegisterValidation("objectid", objectId))
	return validator
}

func password(fl v.FieldLevel) bool {
	field := fl.Field().String()
	if field == "" {
		return true
	}

	if len(field) < 8 || len(field) > 30 {
		return false
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, ch := range field {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case strings.ContainsRune("#$@!%&*?", ch):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

func objectId(fl v.FieldLevel) bool {
	field := fl.Field().String()
	if field == "" {
		return true
	}

	if len(field) != 24 {
		return false
	}

	for _, ch := range field {
		if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')) {
			return false
		}
	}

	return true
}
