package repository

import (
	"github.com/netooo/board-games/app/domain/model"
)

type UserRepository interface {
	Insert(userId, name, email, password string) (*model.User, error)
	GetByUserId(userId string) (*model.User, error)
}
