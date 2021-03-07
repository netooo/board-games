package persistence

import (
	"database/sql"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) Insert(DB *sql.DB, userID, name, email, password string) error {
	panic("implement me")
}

func (up userPersistence) GetByUserID(DB *sql.DB, userID string) (*model.User, error) {
	panic("implement me")
}
