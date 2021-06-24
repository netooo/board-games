package persistence

import (
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
		Status: 0,
	}
	db.Create(&numeron)

	// 作成者を部屋に入室させる
	player := model.NumeronPlayer{
		Numeron: &numeron,
		User:    user,
		Socket:  socket,
	}
	db.Create(&player)

	// Numeron の部屋を起動する
	go numeron.Run()

	return &numeron, nil
}
