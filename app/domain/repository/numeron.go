package repository

import (
	"github.com/netooo/board-games/app/domain/model"
)

type NumeronRepository interface {
	GetNumerons() ([]*model.Numeron, error)
	CreateNumeron(name string, user *model.User) (uint, error)
	ShowNumeron(id string) (*model.Numeron, error)
	EntryNumeron(id string, user *model.User) error
	StartNumeron(id string, user *model.User) error
}
