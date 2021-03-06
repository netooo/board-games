package repository

import (
	"database/sql"
	"github.com/netooo/board-games/domain"
)

type UserRepository interface {
	Insert(DB *sql.DB, userID, name, email, password string) error
	GetByUserID(DB *sql.DB, userID string) (*domain.User, error)
}
