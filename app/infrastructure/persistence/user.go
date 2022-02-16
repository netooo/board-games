package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type userPersistence struct {
	Conn *gorm.DB
}

func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

func (up userPersistence) Insert(userId, name, email, password string) (*model.User, error) {
	user := model.User{
		DisplayId: userId,
		Name:      name,
		Email:     email,
		Password:  password,
	}

	db := config.Connect()
	defer config.Close()

	if err := db.Omit("Game").Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (up userPersistence) GetByUserId(userId string) (*model.User, error) {
	// check the session
	var user model.User

	db := config.Connect()
	defer config.Close()

	if err := db.First(&user, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
