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
var SocketUsers []*model.User

func NewSocketPersistence(conn *gorm.DB) repository.SocketRepository {
	return &socketPersistence{Conn: conn}
}

func (sp socketPersistence) ConnectSocket(user *model.User, socket *websocket.Conn) error {
	// WebSocket接続中だった場合は前のコネクションを切断しルーム情報を引き継ぐ
	// TODO: indexを見ているのでマルチスレッドに対応できていない気がする
	index, _ := model.SearchUser(SocketUsers, user.ID)
	if index != -1 {
		oldUser := SocketUsers[index]
		if room := oldUser.Room; room != nil {
			user.Room = room
			room.Leave <- oldUser
		}
		_ = oldUser.Socket.Close()
		SocketUsers[index] = SocketUsers[len(SocketUsers)-1]
		SocketUsers = SocketUsers[:len(SocketUsers)-1]
	}

	// Socketに紐付いたユーザをsliceに格納
	user.Socket = socket
	SocketUsers = append(SocketUsers, user)

	// Messageを受け取るゴルーチンを起動する
	go user.Read()

	return nil
}

func (sp socketPersistence) DisconnectSocket(user *model.User, socket *websocket.Conn) error {
	index, _ := model.SearchUser(SocketUsers, user.ID)
	if index == -1 {
		return nil
	}

	oldUser := SocketUsers[index]
	if room := oldUser.Room; room != nil {
		room.Leave <- oldUser
	}
	_ = oldUser.Socket.Close()
	SocketUsers[index] = SocketUsers[len(SocketUsers)-1]
	SocketUsers = SocketUsers[:len(SocketUsers)-1]

	return nil
}
