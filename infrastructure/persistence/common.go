package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
)

type commonPersistence struct {
	Conn *gorm.DB
}

type Game int

const (
	Numeron Game = iota
)

func (g Game) String() string {
	switch g {
	case Numeron:
		return "Numeron"
	default:
		return "Unknown"
	}
}

func NewCommonPersistence(conn *gorm.DB) repository.CommonRepository {
	return &commonPersistence{Conn: conn}
}

func (cp commonPersistence) CreateRoom(user model.User, game string) (int, error) {
	db := config.Connect()
	defer config.Close()

	var room interface{}
	switch game {
	case Numeron.String():
		room = model.Numeron{
			Owner:  user,
			Status: "ready",
		}
	}

	db.Create(&room)

	return room.id, nil
}
