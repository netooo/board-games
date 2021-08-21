package usecase

import (
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type AuthUseCase interface {
	Signin(email, password string) (*model.User, error)
}

type authUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(ar repository.AuthRepository) AuthUseCase {
	return &authUseCase{
		authRepository: ar,
	}
}

func (au authUseCase) Signin(email, password string) (*model.User, error) {
	// domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	user, err := au.authRepository.Signin(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
