package repository

import (
	"github.com/netooo/board-games/domain/model"
)

type NumeronPlayerRepository interface {
	SetCode(code string) (*model.NumeronPlayer, error)
}
