package model

import (
	"github.com/jinzhu/gorm"
)

type Numeron struct {
	gorm.Model
	DisplayId string                 `json:"display_id"`
	Name      string                 `json:"name"`
	Status    int                    `json:"status"`
	OwnerId   uint                   `json:"owner_id"`
	Owner     *User                  `json:"-"`
	Turn      int                    `json:"turn"`
	Join      chan *User             `json:"-"`
	Leave     chan *User             `json:"-"`
	Users     map[*User]bool         `json:"-"`
	Players   map[int]*NumeronPlayer `json:"-"`
}

type Message struct {
	Action string
	Value  string
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

func (n *Numeron) Run(owner *User) {
	// 作成者を入室させる
	owner.Game = n.DisplayId
	n.Users[owner] = true

	for {
		// チャネルの動きを監視し、処理を決定する
		select {
		/* Joinチャネルに動きがあった場合(ユーザの入室) */
		case user := <-n.Join:
			msg := Message{
				Action: "join",
				Value:  user.Name,
			}
			for u := range n.Users {
				if err := u.Socket.WriteJSON(msg); err != nil {
					u.Game = ""
					delete(n.Users, u)
				}
			}
			user.Game = n.DisplayId
			n.Users[user] = true

		/* Leaveチャネルに動きがあった場合(ユーザの退室) */
		case user := <-n.Leave:
			user.Game = ""
			delete(n.Users, user)
			msg := Message{
				Action: "leave",
				Value:  user.Name,
			}
			for u := range n.Users {
				if err := u.Socket.WriteJSON(msg); err != nil {
					u.Game = ""
					delete(n.Users, u)
					// ここでもleave?
				}
			}
		}
	}
}
