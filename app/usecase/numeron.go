package usecase

import (
	_ "github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type NumeronUseCase interface {
	CreateRoom(user *model.User, socket *websocket.Conn) (*model.Numeron, error)
	GameStart(userId uint, socket *websocket.Conn, numeronId int, orders interface{}) error
}

type numeronUseCase struct {
	numeronRepository repository.NumeronRepository
}

func NewNumeronUseCase(nr repository.NumeronRepository) NumeronUseCase {
	return &numeronUseCase{
		numeronRepository: nr,
	}
}

func (nu numeronUseCase) CreateRoom(user *model.User, socket *websocket.Conn) (*model.Numeron, error) {
	room, err := nu.numeronRepository.CreateRoom(user, socket)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (nu numeronUseCase) GameStart(userId uint, socket *websocket.Conn, numeronId int, orders interface{}) error {
	err := nu.numeronRepository.GameStart(userId, socket, numeronId, orders)
	if err != nil {
		return err
	}

	return nil
}
