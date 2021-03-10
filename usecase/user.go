package usecase

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
)

type UserUseCase interface {
	GetByUserID(DB *sql.DB, userID string) (model.User, error)
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

func (uu userUseCase) GetByUserID(DB *sql.DB, userID string) (*model.User, error) {
	user, err := uu.userRepository.GetByUserID(DB, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) Insert(DB *sql.DB, name, email, password string) error {
	// 各種パラメータのバリデーションを行う

	userID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	err = uu.userRepository.Insert(DB, userID.String(), name, email, password)
	if err != nil {
		return err
	}
	return nil
}
