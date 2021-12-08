package usecase

import (
	"errors"
	_ "github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
	"strconv"
)

type RoomUseCase interface {
	GetRooms() ([]*model.Room, error)
	CreateRoom(user *model.User, socket *websocket.Conn) (*model.Room, error)
	JoinRoom(id string, user *model.User, socket *websocket.Conn) error
}

type roomUseCase struct {
	roomRepository repository.RoomRepository
}

func NewRoomUseCase(rr repository.RoomRepository) RoomUseCase {
	return &roomUseCase{
		roomRepository: rr,
	}
}

func (ru roomUseCase) GetRooms() ([]*model.Room, error) {

}

func (ru roomUseCase) CreateRoom(user *model.User, socket *websocket.Conn) (*model.Room, error) {
	room, err := ru.roomRepository.CreateRoom(user, socket)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (ru roomUseCase) JoinRoom(id string, user *model.User, socket *websocket.Conn) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	roomId_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("Invalid ID")
	}

	var roomId uint = uint(roomId_)

	err = ru.roomRepository.JoinRoom(roomId, user, socket)
	if err != nil {
		return err
	}

	return nil
}
