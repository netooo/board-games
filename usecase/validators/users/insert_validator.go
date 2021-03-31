package validators

import "github.com/go-playground/validator"

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

func PasswordValidation(fl validator.FieldLevel) bool {
	// TODO: パスワードバリデーター作成
	// 		 小文字1文字以上/大文字1文字以上/数字1文字以上/8文字以上
	return true
}
