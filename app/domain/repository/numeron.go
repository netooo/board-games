package repository

import (
	"github.com/netooo/board-games/app/domain/model"
)

type NumeronRepository interface {
	GetNumerons(userId string) ([]*model.Numeron, error)
	CreateNumeron(name string, userId string) (string, error)
	ShowNumeron(id string, userId string) (*model.Numeron, error)
	EntryNumeron(id string, userId string) error
	LeaveNumeron(id string, userId string) error
	StartNumeron(id string, userId string) error
}
