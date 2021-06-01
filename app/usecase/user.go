package usecase

import (
	_ "github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
	validators "github.com/netooo/board-games/app/interfaces/validators/users"
)

type UserUseCase interface {
	GetByUserId(userId string) (*model.User, error)
	Insert(name, email, password string) (*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) GetByUserId(userId string) (*model.User, error) {
	user, err := uu.userRepository.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) Insert(name, email, password string) (*model.User, error) {
	// リクエストパラメータのバリデーション
	ValidateUser := &validators.InsertUser{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := validators.InsertUserValidate(ValidateUser); err != nil {
		return nil, err
	}

	userId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	user, err := uu.userRepository.Insert(userId.String(), name, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
