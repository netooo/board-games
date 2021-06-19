package repository

import "github.com/netooo/board-games/app/domain/model"

type NumeronPlayerRepository interface {
	SetCode(user *model.User, code string) error
}
