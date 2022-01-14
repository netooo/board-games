package persistence

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
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

func (rp roomPersistence) GetRooms() ([]*model.Room, error) {
	return Rooms, nil
}

func (rp roomPersistence) CreateRoom(name string, user *model.User) (uint, error) {
	db := config.Connect()
	defer config.Close()

	// Room の部屋を作成
	room := model.Room{
		Name:    name,
		Owner:   user,
		OwnerId: user.ID,
		Status:  0,
		Join:    make(chan *model.User),
		Leave:   make(chan *model.User),
		Players: make(map[*model.User]bool),
	}
	if err := db.Omit("Join", "Leave", "Players").Create(&room).Error; err != nil {
		return 0, err
	}

	// 作成されたroomをsliceに格納
	Rooms = append(Rooms, &room)

	// SocketUsersからuserを取得
	index, err := model.SearchUser(SocketUsers, user.ID)
	if err != nil {
		return 0, err
	}
	user = SocketUsers[index]

	// Room の部屋を起動する
	go room.Run(user)

	return room.ID, nil
}

func (rp roomPersistence) JoinRoom(roomId uint, user *model.User) error {
	// Roomsからroomを取得
	index, err := model.SearchRoom(Rooms, roomId)
	if err != nil {
		return err
	}
	room := Rooms[index]

	// 部屋の状態をチェック
	if room.Status != 0 {
		return errors.New("Room is not Ready")
	}

	if len(room.Players) > 4 {
		return errors.New("Limit User in Room")
	}

	// SocketUsersからuserを取得
	index, err = model.SearchUser(SocketUsers, user.ID)
	if err != nil {
		return err
	}
	user = SocketUsers[index]

	// 既に入室済みの場合は弾く
	for p := range room.Players {
		if p.ID == user.ID {
			return errors.New("Already Join the Room")
		}
	}

	// Room の部屋に入室する
	room.Join <- user

	return nil
}

func (rp roomPersistence) ShowRoom(roomId uint) (*model.Room, error) {
	// Roomsからroomを取得
	index, err := model.SearchRoom(Rooms, roomId)
	if err != nil {
		return nil, err
	}
	room := Rooms[index]

	return room, nil
}
