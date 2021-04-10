package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/config"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
)

type userPersistence struct {
	Conn *gorm.DB
}

func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

func (up userPersistence) Insert(userId, name, email, password string) error {
	user := model.User{
		UserId:   userId,
		Name:     name,
		Email:    email,
		Password: password,
	}

	db := config.Connect()
	defer config.Close()

	db.Create(&user)
	// return new session

	return nil
}

func (up userPersistence) GetByUserId(userId string) (*model.User, error) {
	// check the session
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
