package usecase

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
	"github.com/netooo/board-games/usecase/validators"
)

type UserUseCase interface {
	GetByUserId(userId string) (*model.User, error)
	Insert(name, email, password string) error
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

func (uu userUseCase) Insert(name, email, password string) error {
	ValidateUser := &validators.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	validate := validator.New()
	err := validate.Struct(ValidateUser)

	//v := validators.MyValidator()

	//if err := v(ValidateUser); err != nil {
	//	return err
	//}

	userId, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	err = uu.userRepository.Insert(userId.String(), name, email, password)
	if err != nil {
		return err
	}
	return nil
}
