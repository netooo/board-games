package repository

import (
	"github.com/netooo/board-games/domain/model"
)

type UserRepository interface {
	Insert(userId, name, email, password string) error
	GetByUserId(userId string) (*model.User, error)
}
