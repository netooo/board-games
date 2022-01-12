package validators

import (
	"github.com/go-playground/validator"
)

type CreateRoom struct {
	Name string `validate:"required,max=50"`
}

func CreateRoomValidate(i interface{}) error {
	validate := validator.New()
	return validate.Struct(i)
}
