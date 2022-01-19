package usecase

import (
	_ "github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type SocketUseCase interface {
	ConnectSocket(user *model.User, socket *websocket.Conn) error
	DisconnectSocket(user *model.User) error
}

type socketUseCase struct {
	socketRepository repository.SocketRepository
}

func NewSocketUseCase(sr repository.SocketRepository) SocketUseCase {
	return &socketUseCase{
		socketRepository: sr,
	}
}

func (su socketUseCase) ConnectSocket(user *model.User, socket *websocket.Conn) error {
	if err := su.socketRepository.ConnectSocket(user, socket); err != nil {
		return err
	}

	return nil
}

func (su socketUseCase) DisconnectSocket(user *model.User) error {
	if err := su.socketRepository.DisconnectSocket(user); err != nil {
		return err
	}

	return nil
}
