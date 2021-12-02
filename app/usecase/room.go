package usecase

import (
	_ "github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type RoomUseCase interface {
	CreateRoom(user *model.User, socket *websocket.Conn) (*model.Room, error)
	JoinRoom(roomId uint, user *model.User, socket *websocket.Conn) error
}

type roomUseCase struct {
	roomRepository repository.RoomRepository
}

func NewRoomUseCase(rr repository.RoomRepository) RoomUseCase {
	return &roomUseCase{
		roomRepository: rr,
	}
}

func (ru roomUseCase) CreateRoom(user *model.User, socket *websocket.Conn) (*model.Room, error) {
	room, err := ru.roomRepository.CreateRoom(user, socket)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (ru roomUseCase) JoinRoom(roomId uint, user *model.User, socket *websocket.Conn) error {
	err := ru.roomRepository.JoinRoom(roomId, user, socket)
	if err != nil {
		return err
	}

	return nil
}
