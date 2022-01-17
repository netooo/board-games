package model

import (
	"github.com/jinzhu/gorm"
)

type Numeron struct {
	gorm.Model
	DisplayId string         `json:"display_id"`
	Name      string         `json:"name"`
	Status    int            `json:"status"`
	OwnerId   uint           `json:"owner_id"`
	Owner     *User          `json:"-"`
	Join      chan *User     `json:"-"`
	Leave     chan *User     `json:"-"`
	Players   map[*User]bool `json:"-"`
}

type StartOrder struct {
	First  string
	Second string
}

type Status int

const (
	Ready  Status = iota // Ready  == 0
	Play                 // Play   == 1
	Finish               // Finish == 2
)

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
