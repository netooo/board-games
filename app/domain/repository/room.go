package repository

import (
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/infrastructure/persistence"
)

type RoomRepository interface {
	GetRooms() ([]*persistence.RespRoom, error)
	CreateRoom(user *model.User, socket *websocket.Conn) (*model.Room, error)
	JoinRoom(roomId uint, user *model.User, socket *websocket.Conn) error
}
