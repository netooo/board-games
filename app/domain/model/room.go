package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Room struct {
	gorm.Model
	Name    string         `json:"name"`
	Status  int            `json:"status"`
	OwnerId uint           `json:"owner_id"`
	Owner   *User          `json:"-"`
	Join    chan *User     `json:"-"`
	Leave   chan *User     `json:"-"`
	Players map[*User]bool `json:"-"`
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

func (r *Room) Run(user *User) {
	// 作成者を入室させる
	r.Players[user] = true

	for {
		// チャネルの動きを監視し、処理を決定する
		select {
		/* Joinチャネルに動きがあった場合(ユーザの入室) */
		case player := <-r.Join:
			for p := range r.Players {
				if err := p.Socket.WriteJSON(player); err != nil {
					delete(r.Players, p)
				}
			}
			r.Players[player] = true

		/* Leaveチャネルに動きがあった場合(ユーザの退室) */
		case player := <-r.Leave:
			// Player mapから対象ユーザを削除する
			delete(r.Players, player)
		}
	}
}

func SearchRoom(rooms []*Room, roomId uint) (int, error) {
	for i, r := range rooms {
		if r.ID == roomId {
			return i, nil
		}
	}
	return -1, errors.New("Room Not found")
}
