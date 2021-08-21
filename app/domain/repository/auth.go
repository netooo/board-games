package repository

import (
	"github.com/netooo/board-games/app/domain/model"
)

type AuthRepository interface {
	Signin(email, password string) (*model.User, error)
}
