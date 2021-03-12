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

func (up userPersistence) Insert(DB *sql.DB, userId, name, email, password string) error {
	stmt, err := DB.Prepare("INSERT INTO user(user_id, name, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userId, name, email, password)
	return err
}

func (up userPersistence) GetByUserId(DB *sql.DB, userId string) (*model.User, error) {
	row := DB.QueryRow("SELECT * FROM user WHERE user_id = ?", userId)
	return convertToUser(row)
}

// row型をuser型に紐づける
func convertToUser(row *sql.Row) (*model.User, error) {
	user := model.User{}
	err := row.Scan(&user.UserId, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
