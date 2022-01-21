package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type socketPersistence struct {
	Conn *gorm.DB
}

// Websocket接続中のユーザを格納した配列
var SocketUsers = map[string]*model.User{}

func NewSocketPersistence(conn *gorm.DB) repository.SocketRepository {
	return &socketPersistence{Conn: conn}
}

func (sp socketPersistence) ConnectSocket(user *model.User, socket *websocket.Conn) error {
	oldUser, ok := SocketUsers[user.UserId]
	if ok {
		_ = oldUser.Socket.Close()
		oldUser.Socket = socket
	} else {
		user.Socket = socket
		SocketUsers[user.UserId] = user
	}

	return nil
}

func (sp socketPersistence) DisconnectSocket(user *model.User) error {
	oldUser, ok := SocketUsers[user.UserId]
	if !ok {
		return nil
	}

	// Join中のゲームから退出させる。
	// TODO: User#Gameをstringで持っているが、interfaceにしてポインタ参照させてswitchを辞めたい。
	if game := oldUser.Game; game != "" {
		switch game[:3] {
		case "NMR":
			if numeron := Numerons[game]; numeron != nil {
				numeron.Leave <- oldUser
			}
			return nil
		}
	}
	_ = oldUser.Socket.Close()
	delete(SocketUsers, oldUser.UserId)

	return nil
}
