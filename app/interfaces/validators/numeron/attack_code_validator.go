package validators

import (
	"github.com/go-playground/validator"
)

type AttackCode struct {
	Code string `validate:"required,len=3,attack_code_validation"`
}

func AttackCodeValidate(i interface{}) error {
	validate := validator.New()
	_ = validate.RegisterValidation("attack_code_validation", AttackCodeValidation)
	return validate.Struct(i)
}

// AttackCodeValidation 0-9/重複有り
func AttackCodeValidation(fl validator.FieldLevel) bool {
	code := fl.Field().String()

	return checkDigit(`[^0-9]`, code)
}
