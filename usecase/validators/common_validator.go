package validators

import "github.com/go-playground/validator"

// Validate User Struct
type User struct {
	Name     string `validate:"required,max=12"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,password_validation"`
}

// CustomValidator
type CustomValidator struct {
	validator *validator.Validate
}

func MyValidator() Validator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
	_ = cv.validator.RegisterValidation("password_validation", PasswordValidation)
	return cv.validator.Struct(i)
}

func PasswordValidation(fl validator.FieldLevel) bool {
	// TODO: パスワードバリデーター作成
	// 		 小文字1文字以上/大文字1文字以上/数字1文字以上/8文字以上
	return true
}
