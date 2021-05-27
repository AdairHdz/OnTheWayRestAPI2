package validators

import (
	"regexp"
	validator "github.com/go-playground/validator"
)

var (
	validate *validator.Validate = validator.New()
)

func init(){
	validate.RegisterValidation("lettersAndSpaces", LettersAndSpaces)
}

func GetValidator() *validator.Validate{
	return validate
}


func LettersAndSpaces(fieldLevel validator.FieldLevel) bool {
	matches, _ := regexp.MatchString("^[A-z ]+$", fieldLevel.Field().String())
	return matches
}
