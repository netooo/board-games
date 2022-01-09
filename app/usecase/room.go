package usecase

import (
	"errors"
	_ "github.com/go-playground/validator"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
	"strconv"
)

type RoomUseCase interface {
	GetRooms() ([]*model.Room, error)
	CreateRoom(user *model.User) (uint, error)
	JoinRoom(id string, user *model.User) error
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
	rooms, err := ru.roomRepository.GetRooms()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (ru roomUseCase) CreateRoom(user *model.User) (uint, error) {
	roomId, err := ru.roomRepository.CreateRoom(user)
	if err != nil {
		return 0, err
	}

	return roomId, nil
}

func (ru roomUseCase) JoinRoom(id string, user *model.User) error {
	if id == "" {
		return errors.New("ID Not Found")
	}

	roomId_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("Invalid ID")
	}

	var roomId uint = uint(roomId_)

	err = ru.roomRepository.JoinRoom(roomId, user)
	if err != nil {
		return err
	}

	return nil
}
