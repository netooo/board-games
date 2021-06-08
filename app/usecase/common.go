package usecase

import (
	_ "github.com/go-playground/validator"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
	validators "github.com/netooo/board-games/app/interfaces/validators/users"
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
	ValidateRoom := &validators.CreateRoom{
		Game: game,
	}

	if err := validators.CreateRoomValidate(ValidateRoom); err != nil {
		return -1, err
	}

	roomId, err := cu.commonRepository.CreateRoom(user, game)
	if err != nil {
		return -1, err
	}
	return roomId, nil
}
