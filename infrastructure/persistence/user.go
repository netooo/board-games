package persistence

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
)

type userPersistence struct {
	Conn *gorm.DB
}

func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	// TODO: 後で直す
	return &userPersistence{Conn: conn}
}

func (up userPersistence) Insert(userId, name, email, password string) error {
	user := model.User{UserId: userId, Name: name, Email: email, Password: password}

	// DB接続確認
	if err := up.Conn.Take(&user).Error; err != nil {
		return err
	}

	db := config.Connect()
	defer config.Close()

	db.Create(&user)

	return nil
}

func (up userPersistence) GetByUserId(DB *sql.DB, userId string) (*model.User, error) {
	user := model.User{}

	// DB接続確認
	if err := up.Conn.Take(&user).Error; err != nil {
		return nil, err
	}

	db := config.Connect()
	defer config.Close()

	db.Where("UserId = ?", userId).Find(&user)

	return &user, nil
}
