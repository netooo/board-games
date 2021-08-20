package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/netooo/board-games/app/config"
)

type Numeron struct {
	gorm.Model
	Status  int            `json:"status"`
	OwnerId uint           `json:"owner_id"`
	Owner   *User          `json:"-"`
	Join    chan *User     `json:"-"`
	Leave   chan *User     `json:"-"`
	Players map[*User]bool `json:"-"`
}

type NumeronAction struct {
	Action string `json:"action"`
}

type Status int

const (
	Ready  Status = iota // Ready  == 0
	Play                 // Play   == 1
	Finish               // Finish == 2
)

type StartOrder struct {
	First  string
	Second string
}

func (s Status) String() string {
	switch s {
	case Ready:
		return "準備中"
	case Play:
		return "実行中"
	case Finish:
		return "終了"
	}
	return "未定義"
}

/*
ヌメロンルームを起動する
*/
func (n *Numeron) Run(user *User) {
	// 作成者を入室させる
	n.Players[user] = true

	for {
		// チャネルの動きを監視し、処理を決定する
		select {
		/* Joinチャネルに動きがあった場合(ユーザの入室) */
		case player := <-n.Join:
			for p := range n.Players {
				if err := p.Socket.WriteJSON(player); err != nil {
					delete(n.Players, p)
				}
			}
			n.Players[player] = true

		/* Leaveチャネルに動きがあった場合(ユーザの退室) */
		case player := <-n.Leave:
			// Player mapから対象ユーザを削除する
			delete(n.Players, player)
		}
	}
}

func (n *Numeron) Read(user *User, action string, value string) {
	db := config.Connect()
	defer config.Close()

	switch action {
	case "start":
		if n.Status != 0 {
			return
		}

		if user != n.Owner {
			return
		}

		var order StartOrder
		if err := json.Unmarshal([]byte(value), &order); err != nil {
			return
		}
		if len(n.Players) != 2 {
			return
		}

		var first string = order.First
		var second string = order.Second
		if first == "" || second == "" {
			//TODO: 順番の指定がない場合はランダムにしたい
			return
		}

		var firstUser User
		var secondUser User
		if err := db.Omit("Socket").First(&firstUser, first).Error; err != nil {
			return
		}
		if err := db.Omit("Socket").First(&secondUser, first).Error; err != nil {
			return
		}

		firstPlayer := NumeronPlayer{
			Numeron: n,
			User:    &firstUser,
			Order:   0,
		}
		if err := db.Create(firstPlayer).Error; err != nil {
			return
		}

		secondPlayer := NumeronPlayer{
			Numeron: n,
			User:    &secondUser,
			Order:   1,
		}
		if err := db.Create(secondPlayer).Error; err != nil {
			return
		}

		if err := db.Omit("Join", "Leave", "Players").Model(&n).Update("Status", 1).Error; err != nil {
			return
		}

		action := NumeronAction{
			Action: "start",
		}

		for p := range n.Players {
			if err := p.Socket.WriteJSON(action); err != nil {
				delete(n.Players, p)
			}
		}
	default:

	}
}
