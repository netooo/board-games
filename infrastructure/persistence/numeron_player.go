package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
)

type numeronPlayerPersistence struct {
	Conn *gorm.DB
}

func NewNumeronPlayerPersistence(conn *gorm.DB) repository.NumeronPlayerRepository {
	return &numeronPlayerPersistence{Conn: conn}
}

func (npp numeronPlayerPersistence) SetCode(code string) (*model.NumeronPlayer, error) {
	numeronPlayer := model.NumeronPlayer{
		Code: code,
	}

	db := config.Connect()
	defer config.Close()

	db.UpdateColumn(&numeronPlayer)

	return &numeronPlayer, nil
}
