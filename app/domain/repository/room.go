package repository

import (
	"github.com/netooo/board-games/app/domain/model"
)

type RoomRepository interface {
	GetRooms() ([]*model.Room, error)
	CreateRoom(name string, user *model.User) (uint, error)
	JoinRoom(roomId uint, user *model.User) error
}
