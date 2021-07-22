package repository

import "github.com/netooo/board-games/app/domain/model"

type NumeronPlayerRepository interface {
	SetCode(user *model.User, id string, code string) error
	JoinRoom(user *model.User, id string) error
}
