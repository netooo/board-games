package persistence

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type roomPersistence struct {
	Conn *gorm.DB
}

func NewRoomPersistence(conn *gorm.DB) repository.RoomRepository {
	return &roomPersistence{Conn: conn}
}

func (np roomPersistence) CreateRoom(user *model.User, socket *websocket.Conn) (*model.Room, error) {
	db := config.Connect()
	defer config.Close()

	// Room の部屋を作成
	room := model.Room{
		Owner:   user,
		Status:  0,
		Join:    make(chan *model.User),
		Leave:   make(chan *model.User),
		Players: make(map[*model.User]bool),
	}
	if err := db.Omit("Join", "Leave", "Players").Create(&room).Error; err != nil {
		return nil, err
	}

	// 作成者のsocketをつなぐ
	user.Socket = socket

	// Room の部屋を起動する
	go room.Run(user)

	return &room, nil
}

func (np roomPersistence) JoinRoom(roomId string, user *model.User, socket *websocket.Conn) error {
	db := config.Connect()
	defer config.Close()

	// Room の部屋を取得
	var room model.Room
	if err := db.Omit("Join", "Leave", "Players").First(&room, roomId).Error; err != nil {
		return err
	}

	// 部屋の状態をチェック
	if room.Status != 0 {
		return errors.New("Room is not Ready")
	}

	if len(room.Players) > 1 {
		return errors.New("Limit User in Room Room")
	}

	for p := range room.Players {
		if p == user {
			return errors.New("Already Join the Room Room")
		}
	}

	// 作成者のsocketをつなぐ
	user.Socket = socket

	// Room の部屋に入室する
	room.Join <- user

	return nil
}
