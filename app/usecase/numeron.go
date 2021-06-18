package usecase

import (
	_ "github.com/go-playground/validator"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type NumeronUseCase interface {
	CreateRoom(user *model.User) (*model.Numeron, error)
}

type numeronUseCase struct {
	numeronRepository repository.NumeronRepository
}

func NewNumeronUseCase(nr repository.NumeronRepository) NumeronUseCase {
	return &numeronUseCase{
		numeronRepository: nr,
	}
}

func (nu numeronUseCase) CreateRoom(user *model.User) (*model.Numeron, error) {
	room, err := nu.numeronRepository.CreateRoom(user)
	if err != nil {
		return nil, err
	}

	return room, nil
}
