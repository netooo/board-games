package validators

import (
	"github.com/go-playground/validator"
	"regexp"
)

// Validate User Struct
type InsertUser struct {
	Name     string `validate:"required,max=12"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,password_validation"`
}

// Validate validate
func UserValidate(i interface{}) error {
	validate := validator.New()
	_ = validate.RegisterValidation("password_validation", PasswordValidation)
	return validate.Struct(i)
}

// 計8文字以上/数字1文字以上/小文字1文字以上/大文字1文字以上
func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}
	if !checkRegexp(`[0-9]`, password) {
		return false
	}
	if !checkRegexp(`[a-z]`, password) {
		return false
	}
	if !checkRegexp(`[A-Z]`, password) {
		return false
	}

	return true
}

func checkRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
