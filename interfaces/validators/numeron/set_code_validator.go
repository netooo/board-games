package validators

import (
	"github.com/go-playground/validator"
	"regexp"
	"strings"
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

	if checkDigit(`[^0-9]`, code) || !checkDuplication(code) {
		return false
	}

	return true
}

// 0-9チェック
func checkDigit(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

// 重複チェック
// TODO: チェック処理がイケてない。もっとスマートに書きたい。
func checkDuplication(code string) bool {
	first := code[0:1]
	second := code[1:2]
	return strings.Count(code, first) == 1 && strings.Count(code, second) == 1
}
