package persistence

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type numeronPersistence struct {
	Conn *gorm.DB
}

// 立ち上がっているNumeronを格納した配列
var Numerons []*model.Numeron

func NewNumeronPersistence(conn *gorm.DB) repository.NumeronRepository {
	return &numeronPersistence{Conn: conn}
}

func (p numeronPersistence) GetNumerons() ([]*model.Numeron, error) {
	return Numerons, nil
}

func (p numeronPersistence) CreateNumeron(name string, user *model.User) (uint, error) {
	db := config.Connect()
	defer config.Close()

	// Numeron の部屋を作成
	numeron := model.Numeron{
		Name:    name,
		Owner:   user,
		OwnerId: user.ID,
		Status:  0,
		Join:    make(chan *model.User),
		Leave:   make(chan *model.User),
		Players: make(map[*model.User]bool),
	}
	if err := db.Omit("Join", "Leave", "Players").Create(&numeron).Error; err != nil {
		return 0, err
	}

	// 作成されたnumeronをsliceに格納
	Numerons = append(Numerons, &numeron)

	// SocketUsersからuserを取得
	index, err := model.SearchUser(SocketUsers, user.ID)
	if err != nil {
		return 0, err
	}
	user = SocketUsers[index]

	// Numeron の部屋を起動する
	go numeron.Run(user)

	return numeron.ID, nil
}

func (p numeronPersistence) EntryNumeron(numeronId uint, user *model.User) error {
	// Numeronsからnumeronを取得
	index, err := model.SearchNumeron(Numerons, numeronId)
	if err != nil {
		return err
	}
	numeron := Numerons[index]

	// 部屋の状態をチェック
	if numeron.Status != 0 {
		return errors.New("Numeron is not Ready")
	}

	if len(numeron.Players) > 4 {
		return errors.New("Limit User in Numeron")
	}

	// SocketUsersからuserを取得
	index, err = model.SearchUser(SocketUsers, user.ID)
	if err != nil {
		return err
	}
	user = SocketUsers[index]

	// 既に入室済みの場合は弾く
	for p := range numeron.Players {
		if p.ID == user.ID {
			return errors.New("Already Join the Numeron")
		}
	}

	// Numeron の部屋に入室する
	numeron.Join <- user

	return nil
}

func (p numeronPersistence) ShowNumeron(numeronId uint) (*model.Numeron, error) {
	// Numeronsからnumeronを取得
	index, err := model.SearchNumeron(Numerons, numeronId)
	if err != nil {
		return nil, err
	}
	numeron := Numerons[index]

	return numeron, nil
}

func (p numeronPersistence) StartNumeron(numeronId uint, user *model.User) error {
	// Numeronsからnumeronを取得
	index, err := model.SearchNumeron(Numerons, numeronId)
	if err != nil {
		return err
	}
	numeron := Numerons[index]

	// 部屋の状態をチェック
	if numeron.Status != 0 {
		return errors.New("Numeron is not Ready")
	}

	// 開始人数を上回っているかチェック
	if len(numeron.Players) != 2 {
		return errors.New("Inappropriate Number of Players")
	}

	// Request UserがNumeronsに存在しない場合は弾く
	var userIds []uint
	for k, _ := range numeron.Players {
		userIds = append(userIds, k.ID)
	}
	if !isContains(userIds, user.ID) {
		return errors.New("Invalid Request User")
	}

	return nil
}

func isContains(ids []uint, id uint) bool {
	for _, i := range ids {
		if i == id {
			return true
		}
	}
	return false
}
