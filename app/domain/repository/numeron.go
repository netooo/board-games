package repository

import (
	"github.com/netooo/board-games/app/domain/model"
)

type NumeronRepository interface {
	GetNumerons() ([]*model.Numeron, error)
	CreateNumeron(name string, userId string) (string, error)
	ShowNumeron(id string, userId string) (*model.Numeron, error)
	EntryNumeron(id string, userId string) error
	LeaveNumeron(id string, userId string) error
	StartNumeron(id string, userId string, firstId string, secondId string) error
	SetNumeron(id string, userId string, code string) error
	AttackNumeron(id string, userId string, code string) error
}
