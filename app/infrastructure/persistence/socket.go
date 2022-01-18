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
var SocketUsers map[string]*model.User

func NewSocketPersistence(conn *gorm.DB) repository.SocketRepository {
	return &socketPersistence{Conn: conn}
}

func (sp socketPersistence) ConnectSocket(user *model.User, socket *websocket.Conn) error {
	// WebSocket接続中だった場合は前のコネクションを切断しルーム情報を引き継ぐ
	// 新しいSocketにルーム情報を引き継がせる方が良いかも
	oldUser, ok := SocketUsers[user.UserId]
	if !ok {
		if game := oldUser.Game; game != nil {
			user.Game = game
			game.Leave <- oldUser
		}
		_ = oldUser.Socket.Close()
		delete(SocketUsers, oldUser.UserId)
	}

	// Socketに紐付いたユーザをmapに格納
	user.Socket = socket
	SocketUsers[user.UserId] = user

	// Messageを受け取るゴルーチンを起動する
	go user.Read()

	return nil
}

func (sp socketPersistence) DisconnectSocket(user *model.User, socket *websocket.Conn) error {
	oldUser, ok := SocketUsers[user.UserId]
	if !ok {
		return nil
	}

	if game := oldUser.Game; game != nil {
		game.Leave <- oldUser
	}
	_ = oldUser.Socket.Close()
	delete(SocketUsers, oldUser.UserId)

	return nil
}
