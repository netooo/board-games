package repository

import (
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/domain/model"
)

type SocketRepository interface {
	ConnectSocket(user *model.User, socket *websocket.Conn) error
	DisconnectSocket(user *model.User, socket *websocket.Conn) error
}
