package model

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/config"
)

type User struct {
	gorm.Model
	UserId   string          `json:"user_id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	RoomId   int             `json:"room_id"`
	Room     *Room           `json:"-"`
	Socket   *websocket.Conn `json:"-"`
}

type Message struct {
	Id     int
	Game   string
	Action string
	Value  string
}

/*
対象ユーザから送られてきたsocketメッセージを受け取る
*/
func (u *User) Read() {
	db := config.Connect()
	defer config.Close()

	// websocketからjson形式でメッセージを読み出す。
	// 読み込みは無限ループで実行される。
	for {
		var msg *Message
		if err := u.Socket.ReadJSON(&msg); err != nil {
			break
		} else {
			switch msg.Game {
			case "numeron":
				var numeron Numeron
				if err := db.Omit("Join", "Leave", "Players").First(&numeron, msg.Id).Error; err != nil {
					break
				}
				numeron.Read(u, msg.Action, msg.Value)
			}
		}
	}
	_ = u.Socket.Close()
}
