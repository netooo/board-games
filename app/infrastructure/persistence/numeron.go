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

type numeronPersistence struct {
	Conn *gorm.DB
}

func NewNumeronPersistence(conn *gorm.DB) repository.NumeronRepository {
	return &numeronPersistence{Conn: conn}
}

func (np numeronPersistence) CreateRoom(user *model.User, socket *websocket.Conn) (*model.Numeron, error) {
	db := config.Connect()
	defer config.Close()

	// Numeron の部屋を作成
	numeron := model.Numeron{
		Owner:   user,
		Status:  0,
		Join:    make(chan *model.User),
		Leave:   make(chan *model.User),
		Players: make(map[*model.User]bool),
	}
	if err := db.Omit("Join", "Leave", "Players").Create(&numeron).Error; err != nil {
		return nil, err
	}

	// 作成者のsocketをつなぐ
	user.Socket = socket

	// Numeron の部屋を起動する
	go numeron.Run(user)

	return &numeron, nil
}

func (np numeronPersistence) JoinRoom(numeronId string, user *model.User, socket *websocket.Conn) error {
	db := config.Connect()
	defer config.Close()

	// Numeron の部屋を取得
	var numeron model.Numeron
	if err := db.Omit("Join", "Leave", "Players").First(&numeron, numeronId).Error; err != nil {
		return err
	}

	// 作成者のsocketをつなぐ
	user.Socket = socket

	// Numeron の部屋に入室する
	if len(numeron.Players) > 1 {
		return errors.New("Limit User in Numeron Room")
	}
	numeron.Join <- user

	return nil
}
