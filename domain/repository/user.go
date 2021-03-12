package repository

import (
	"database/sql"
	"github.com/netooo/board-games/domain/model"
)

type UserRepository interface {
	Insert(DB *sql.DB, userId, name, email, password string) error
	GetByUserId(DB *sql.DB, userId string) (*model.User, error)
}
