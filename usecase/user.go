package usecase

import (
	"database/sql"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
)

type UserUseCase interface {
	GetByUesrID(DB *sql.DB, userID string) (model.User, error)
	Insert(DB *sql.DB, userID, name, email, password string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) *userUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}
