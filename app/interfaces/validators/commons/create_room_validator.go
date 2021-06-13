package validators

import (
	"github.com/go-playground/validator"
)

// Validate User Struct
type CreateRoom struct {
	Game string `validate:"required,game_validation"`
}

// Validate validate
func CreateRoomValidate(i interface{}) error {
	validate := validator.New()
	_ = validate.RegisterValidation("game_validation", GameValidation)
	return validate.Struct(i)
}

// 指定されたGameが存在するか
func GameValidation(fl validator.FieldLevel) bool {
	game := fl.Field().String()

	if !checkGame(game) {
		return false
	}

	return true
}

func checkGame(game string) bool {
	switch game {
	case "Numeron":
		return true
	default:
		return false
	}
}
