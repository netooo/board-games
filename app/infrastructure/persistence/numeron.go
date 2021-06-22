package persistence

import (
	_ "github.com/go-sql-driver/mysql"
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

func (np numeronPersistence) CreateRoom(user *model.User) (*model.Numeron, error) {
	db := config.Connect()
	defer config.Close()

	numeron := model.Numeron{
		Status: 0,
	}

	db.Create(&numeron)

	// Numeron の部屋を起動する
	go Run(&numeron)

	return &numeron, nil
}

func Run(room *model.Numeron) {
	// TODO: Numeron Room を起動
	// Web Socket の開通
}
