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
	Game     *interface{}    `json:"-"`
	Socket   *websocket.Conn `json:"-"`
}

type Message struct {
	Id     int
	Game   string
	Action string
	Value  string
}

type PushMessage struct {
	RoomId uint
}

/*
対象ユーザから送られてきたsocketメッセージを受け取る
*/
func (user *User) Read() {
	// websocketからjson形式でメッセージを読み出す。
	// 読み込みは無限ループで実行される。
	for {
		var msg *Message
		if err := user.Socket.ReadJSON(&msg); err != nil {
			break
		} else {

		}
	}
	_ = user.Socket.Close()
}
