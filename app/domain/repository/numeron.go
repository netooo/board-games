package repository

import "github.com/netooo/board-games/app/domain/model"

type NumeronRepository interface {
	CreateRoom() (*model.Numeron, error)
}
