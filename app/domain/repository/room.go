package repository

import (
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/domain/model"
)

type RoomRepository interface {
	GetRooms() ([]*model.Room, error)
	CreateRoom(user *model.User, socket *websocket.Conn) error
	JoinRoom(roomId uint, user *model.User, socket *websocket.Conn) error
}
