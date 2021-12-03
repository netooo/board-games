package model

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
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
	//db := config.Connect()
	//defer config.Close()

	// websocketからjson形式でメッセージを読み出す。
	// 読み込みは無限ループで実行される。
	for {
		var msg *Message
		if err := u.Socket.ReadJSON(&msg); err != nil {
			break
		} else {

		}
	}
	_ = u.Socket.Close()
}
