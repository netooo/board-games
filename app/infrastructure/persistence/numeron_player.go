package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type numeronPlayerPersistence struct {
	Conn *gorm.DB
}

func NewNumeronPlayerPersistence(conn *gorm.DB) repository.NumeronPlayerRepository {
	return &numeronPlayerPersistence{Conn: conn}
}

func (npp numeronPlayerPersistence) SetCode(user *model.User, id string, code string) error {
	db := config.Connect()
	defer config.Close()

	//TODO: web socket から NumeronPlayer を特定
	var player model.NumeronPlayer
	db.First(&player, "NumeronId=? AND UserId=?", id, user.ID)
	db.Model(&player).Update("Code", code)

	return nil
}

func (npp numeronPlayerPersistence) JoinRoom(user *model.User, id string) error {
	db := config.Connect()
	defer config.Close()

	var numeron model.Numeron
	if err := db.First(&numeron, "Id=?", id).Error; err != nil {
		return err
	}

	player := model.NumeronPlayer{
		Numeron: &numeron,
		User:    user,
	}
	db.Create(&player)

	// Numeron の部屋に入室する
	numeron.Join <- &player
	defer func() {
		numeron.Leave <- &player
	}()

	return nil
}
