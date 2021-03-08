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
	stmt, err := DB.Prepare("INSERT INTO user(user_id, name, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, name, email, password)
	return err
}

func (up userPersistence) GetByUserID(DB *sql.DB, userID string) (*model.User, error) {
	panic("implement me")
}
