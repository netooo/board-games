package repository

import (
	"github.com/netooo/board-games/domain/model"
)

type CommonRepository interface {
	CreateRoom(user model.User, game string) (int, error)
}
