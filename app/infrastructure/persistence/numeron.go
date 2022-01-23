package persistence

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/config"
	"github.com/netooo/board-games/app/domain/model"
	"github.com/netooo/board-games/app/domain/repository"
)

type numeronPersistence struct {
	Conn *gorm.DB
}

// 立ち上がっているNumeronを格納した配列
var Numerons = map[string]*model.Numeron{}

func NewNumeronPersistence(conn *gorm.DB) repository.NumeronRepository {
	return &numeronPersistence{Conn: conn}
}

func (p numeronPersistence) GetNumerons(userId string) ([]*model.Numeron, error) {
	// SocketUsersからuserを取得
	_, ok := SocketUsers[userId]
	if !ok {
		return nil, errors.New("Invalid Request User")
	}

	var numerons []*model.Numeron
	for _, v := range Numerons {
		numerons = append(numerons, v)
	}

	return numerons, nil
}

func (p numeronPersistence) CreateNumeron(name string, userId string) (string, error) {
	// SocketUsersからuserを取得
	user, ok := SocketUsers[userId]
	if !ok {
		return "", errors.New("Invalid Request User")
	}

	// Numeron のインスタンスを作成
	// DBへの登録はゲーム開始時に行う
	numeron := model.Numeron{
		DisplayId: generateDisplayId(),
		Name:      name,
		Owner:     user,
		OwnerId:   user.ID,
		Status:    0,
		Join:      make(chan *model.User),
		Leave:     make(chan *model.User),
		Players:   make(map[*model.User]bool),
	}

	// 作成されたnumeronをmapに格納
	Numerons[numeron.DisplayId] = &numeron

	// Numeron の部屋を起動する
	go numeron.Run(user)

	return numeron.DisplayId, nil
}

func (p numeronPersistence) EntryNumeron(id string, userId string) error {
	// Numeronsからnumeronを取得
	numeron, ok := Numerons[id]
	if !ok {
		return errors.New("Numeron Not Found")
	}

	// 部屋の状態をチェック
	if numeron.Status != 0 {
		return errors.New("Numeron is not Ready")
	}

	if len(numeron.Players) > 2 {
		return errors.New("Limit User in Numeron")
	}

	// SocketUsersからuserを取得
	user, ok := SocketUsers[userId]
	if !ok {
		return errors.New("Invalid Request User")
	}

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

func (p numeronPersistence) LeaveNumeron(id string, userId string) error {
	// Numeronsからnumeronを取得
	numeron, ok := Numerons[id]
	if !ok {
		return errors.New("Numeron Not Found")
	}

	// 部屋の状態をチェック
	if numeron.Status == 2 {
		return errors.New("Numeron was Finished")
	}

	// SocketUsersからuserを取得
	user, ok := SocketUsers[userId]
	if !ok {
		return errors.New("Invalid Request User")
	}

	// 既に退室済みの場合は弾く
	ok = false
	for p := range numeron.Players {
		if p.ID == user.ID {
			ok = true
			break
		}
	}

	if !ok {
		return errors.New("Already Leave the Numeron")
	}

	// Numeron の部屋から退室する
	numeron.Leave <- user

	return nil
}

func (p numeronPersistence) ShowNumeron(id string, userId string) (*model.Numeron, error) {
	// SocketUsersからuserを取得
	_, ok := SocketUsers[userId]
	if !ok {
		return nil, errors.New("Invalid Request User")
	}

	// Numeronsからnumeronを取得
	numeron, ok := Numerons[id]
	if !ok {
		return nil, errors.New("Numeron Not Found")
	}

	return numeron, nil
}

func (p numeronPersistence) StartNumeron(id string, userId string) error {
	// SocketUsersからuserを取得
	_, ok := SocketUsers[userId]
	if !ok {
		return errors.New("Invalid Request User")
	}

	// Numeronsからnumeronを取得
	numeron, ok := Numerons[id]
	if !ok {
		return errors.New("Numeron Not Found")
	}

	// 部屋の状態をチェック
	if numeron.Status != 0 {
		return errors.New("Numeron is not Ready")
	}

	// 人数が妥当かチェック
	if len(numeron.Players) != 2 {
		return errors.New("Inappropriate Number of Players")
	}

	// Request UserがNumeronsに存在しない場合は弾く
	var userIds []string
	for k, _ := range numeron.Players {
		userIds = append(userIds, k.UserId)
	}

	if !isContains(userIds, userId) {
		return errors.New("Invalid Request User")
	}

	// Play中に変更
	// TODO: numeron_playerなども作成
	numeron.Status = 1

	db := config.Connect()
	defer config.Close()
	if err := db.Omit("Join", "Leave", "Players").Create(&numeron).Error; err != nil {
		return err
	}

	return nil
}

func isContains(ids []string, id string) bool {
	for _, i := range ids {
		if i == id {
			return true
		}
	}
	return false
}

func generateDisplayId() string {
	db := config.Connect()
	defer config.Close()

	var id string

	for true {
		id = "NMR" + uuid.NewString()[0:8]

		// 起動中のゲームから同一のIDを検索
		if _, ok := Numerons[id]; ok {
			continue
		}

		// DBから同一のIDを検索
		numeron := model.Numeron{}
		if !db.First(&numeron, "Display_id = ?", id).RecordNotFound() {
			continue
		}

		break
	}

	return id
}
