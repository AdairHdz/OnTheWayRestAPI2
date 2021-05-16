package validators

import (
	"regexp"
	validator "github.com/go-playground/validator"
)

func LettersAndSpaces(fieldLevel validator.FieldLevel) bool {
	matches, _ := regexp.MatchString("^[A-z ]+$", fieldLevel.Field().String())
	return matches
}