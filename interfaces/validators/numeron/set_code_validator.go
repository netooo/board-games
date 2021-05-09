package validators

import (
	"github.com/go-playground/validator"
	"regexp"
)

// SetCode Validate Set Code
type SetCode struct {
	Code string `validate:"required,len=3,set_code_validation"`
}

// SetCodeValidate Validate validate
func SetCodeValidate(i interface{}) error {
	validate := validator.New()
	_ = validate.RegisterValidation("set_code_validation", SetCodeValidation)
	return validate.Struct(i)
}

// SetCodeValidation 3文字/0-9/重複不可
func SetCodeValidation(fl validator.FieldLevel) bool {
	code := fl.Field().String()

	return checkDigit(`[^0-9]`, code) && checkDuplication(code)
}

// [0-9]のみ => true
func checkDigit(reg, str string) bool {
	return !regexp.MustCompile(reg).Match([]byte(str))
}

// 重複無し => true
func checkDuplication(code string) bool {
	return code[0:1] != code[1:2] && code[1:2] != code[2:3]
}
