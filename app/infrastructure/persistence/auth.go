package persistence

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type authPersistence struct {
	Conn *gorm.DB
}

func NewAuthPersistence(conn *gorm.DB) repository.AuthRepository {
	return &authPersistence{Conn: conn}
}

func (ap authPersistence) Signin(email, password string) (*model.User, error) {
	var user model.User

	db := config.Connect()
	defer config.Close()

	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invarit Email or Password")
	}

	return &user, nil
}
