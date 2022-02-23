package repository

import (
	"github.com/netooo/board-games/app/domain/model"
)

type UserRepository interface {
	GetByUserId(userId string) (*model.User, error)
	Insert(userId, name, email, password string) (*model.User, error)
	BasicSignin(email, password string) (*model.User, error)
}
