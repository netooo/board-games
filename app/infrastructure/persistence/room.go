package persistence

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type roomPersistence struct {
	Conn *gorm.DB
}

// 立ち上がっているRoomを格納した配列
var Rooms []*model.Room

func NewRoomPersistence(conn *gorm.DB) repository.RoomRepository {
	return &roomPersistence{Conn: conn}
}

func (rp roomPersistence) GetRooms() ([]*model.ResponseRoom, error) {
	var activeRooms []*model.ResponseRoom

	for _, r_ := range Rooms {
		r := model.ResponseRoom{
			Id:      r_.ID,
			Name:    r_.Name,
			Owner:   r_.Owner.Name,
			Players: len(r_.Players),
		}
		activeRooms = append(activeRooms, &r)
	}

	return activeRooms, nil
}

func (rp roomPersistence) CreateRoom(user *model.User, socket *websocket.Conn) (*model.Room, error) {
	db := config.Connect()
	defer config.Close()

	// Room の部屋を作成
	room := model.Room{
		OwnerId: user.ID,
		Status:  0,
		Join:    make(chan *model.User),
		Leave:   make(chan *model.User),
		Players: make(map[*model.User]bool),
	}
	if err := db.Omit("Join", "Leave", "Players").Create(&room).Error; err != nil {
		return nil, err
	}

	// 作成されたroomをsliceに格納
	Rooms = append(Rooms, &room)

	// 作成者のsocketをつなぐ
	user.Socket = socket

	// Room の部屋を起動する
	go room.Run(user)

	return &room, nil
}

func (rp roomPersistence) JoinRoom(roomId uint, user *model.User, socket *websocket.Conn) error {
	// Room の部屋を取得
	index, err := searchRooms(Rooms, roomId)
	if err != nil {
		return err
	}

	var room *model.Room = Rooms[index]

	// 部屋の状態をチェック
	if room.Status != 0 {
		return errors.New("Room is not Ready")
	}

	if len(room.Players) > 4 {
		return errors.New("Limit User in Room")
	}

	for p := range room.Players {
		if p.ID == user.ID {
			return errors.New("Already Join the Room")
		}
	}

	// 作成者のsocketをつなぐ
	user.Socket = socket

	// Room の部屋に入室する
	room.Join <- user

	return nil
}

func searchRooms(rooms []*model.Room, roomId uint) (int, error) {
	for i, r := range rooms {
		if r.ID == roomId {
			return i, nil
		}
	}
	return -1, errors.New("Room Not found")
}
