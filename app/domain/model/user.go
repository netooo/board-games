package model

import (
	"errors"
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

func SearchUser(users []*User, userId uint) (int, error) {
	for i, u := range users {
		if u.ID == userId {
			return i, nil
		}
	}
	return -1, errors.New("User Not found")
}
