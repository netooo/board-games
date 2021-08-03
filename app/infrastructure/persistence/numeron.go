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
		Status: 0,
	}
	if err := db.Create(&numeron).Error; err != nil {
		return nil, err
	}

	// 作成者のプレイヤー情報を作成
	player := model.NumeronPlayer{
		Numeron: &numeron,
		User:    user,
		Socket:  socket,
	}
	if err := db.Create(&player).Error; err != nil {
		return nil, err
	}

	// Numeron の部屋を起動する
	go numeron.Run()

	// 作成者を入室させる
	numeron.Join <- &player
	defer func() {
		numeron.Leave <- &player
	}()

	return &numeron, nil
}

func (np numeronPersistence) GameStart(user *model.User, socket *websocket.Conn, numeronId int, orders interface{}) error {
	db := config.Connect()
	defer config.Close()

	var numeron model.Numeron
	if err := db.First(&numeron, "NumeronId=?", numeronId).Error; err != nil {
		return err
	}

	if numeron.Status != 0 {
		return errors.New("Invalid room status")
	}

	if err := db.Model(&numeron).Update("Status", 1).Error; err != nil {
		return err
	}

	return nil
}
