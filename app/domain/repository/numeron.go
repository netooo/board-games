package repository

import (
	"github.com/netooo/board-games/app/domain/model"
)

type NumeronRepository interface {
	GetNumerons() ([]*model.Numeron, error)
	CreateNumeron(name string, user *model.User) (uint, error)
	ShowNumeron(roomId uint) (*model.Numeron, error)
	EntryNumeron(roomId uint, user *model.User) error
	StartNumeron(roomId uint, user *model.User) error
}
