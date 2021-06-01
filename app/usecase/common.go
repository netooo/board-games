package usecase

import (
	_ "github.com/go-playground/validator"
	"github.com/netooo/board-games/domain/model"
	"github.com/netooo/board-games/domain/repository"
)

type CommonUseCase interface {
	CreateRoom(user model.User, game string) (int, error)
}

type commonUseCase struct {
	commonRepository repository.CommonRepository
}

func NewCommonUseCase(cr repository.CommonRepository) CommonUseCase {
	return &commonUseCase{
		commonRepository: cr,
	}
}

func (cu commonUseCase) CreateRoom(user model.User, game string) (int, error) {
	// リクエストパラメータのバリデーション
	// TBD
}
