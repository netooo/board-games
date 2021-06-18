package repository

import "github.com/netooo/board-games/app/domain/model"

type NumeronRepository interface {
	CreateRoom(user *model.User) (*model.Numeron, error)
}
