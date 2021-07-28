package repository

import (
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/domain/model"
)

type NumeronRepository interface {
	CreateRoom(user *model.User, socket *websocket.Conn) (*model.Numeron, error)
	GameStart(user *model.User, socket *websocket.Conn, numeronId int) (interface{}, error)
}
