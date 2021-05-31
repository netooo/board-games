package validators

import (
	"github.com/go-playground/validator"
	"regexp"
)

// InsertUser Validate User Struct
type InsertUser struct {
	Name     string `validate:"required,max=12"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=255,password_validation"`
}

// InsertUserValidate Validate validate
func InsertUserValidate(i interface{}) error {
	validate := validator.New()
	_ = validate.RegisterValidation("password_validation", PasswordValidation)
	return validate.Struct(i)
}

// PasswordValidation 計8文字以上/数字1文字以上/小文字1文字以上/大文字1文字以上
func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if !checkExists(`[0-9]`, password) {
		return false
	}
	if !checkExists(`[a-z]`, password) {
		return false
	}
	if !checkExists(`[A-Z]`, password) {
		return false
	}

	return true
}

func checkExists(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
