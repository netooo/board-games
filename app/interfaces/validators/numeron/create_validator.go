package validators

import (
	"github.com/go-playground/validator"
)

type CreateNumeron struct {
	Name string `validate:"required,max=50"`
}

func CreateNumeronValidate(i interface{}) error {
	validate := validator.New()
	return validate.Struct(i)
}
