package usecase

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
)

type UserUseCase interface {
	GetByUserId(DB *sql.DB, userId string) (model.User, error)
	Insert(DB *sql.DB, userId, name, email, password string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) *userUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) GetByUserId(DB *sql.DB, userId string) (*model.User, error) {
	user, err := uu.userRepository.GetByUserId(DB, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) Insert(DB *sql.DB, name, email, password string) error {
	// 各種パラメータのバリデーションを行う

	userId, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	err = uu.userRepository.Insert(DB, userId.String(), name, email, password)
	if err != nil {
		return err
	}
	return nil
}
