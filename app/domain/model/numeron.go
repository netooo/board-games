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
	Start     chan *User             `json:"-"`
	SetCode   chan *User             `json:"-"`
	Attack    chan *User             `json:"-"`
	Finish    chan *User             `json:"-"`
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

		case user := <-n.Start:
			msg := Message{
				Action: "start",
				Value:  user.Name,
			}

			for _, p := range n.Players {
				if err := p.User.Socket.WriteJSON(msg); err != nil {
					p.User.Game = ""
					delete(n.Players, p.Order)
					// ここでもleave?
				}
			}

		/* SetCodeチャネルに動きがあった場合(コードの設定) */
		case user := <-n.SetCode:
			all := true
			for _, p := range n.Players {
				if p.Code == "" {
					all = false
					break
				}
			}

			if all {
				msg := Message{
					Action: "completed_code",
					Value:  user.UserId,
				}

				for _, p := range n.Players {
					if err := p.User.Socket.WriteJSON(msg); err != nil {
						// ここでもleave?
					}
				}
			} else {
				msg := Message{
					Action: "set_code",
					Value:  "",
				}

				if err := user.Socket.WriteJSON(msg); err != nil {
					// ここでもleave?
				}
			}

		/* Attackチャネルに動きがあった場合(コードの宣言) */
		case user := <-n.Attack:
			var result string
			for _, p := range n.Players {
				if p.UserId == user.ID {
					result = p.Result
					break
				}
			}

			msg := Message{
				Action: "attack_code",
				Value:  result,
			}

			for _, p := range n.Players {
				if err := p.User.Socket.WriteJSON(msg); err != nil {
					// ここでもleave?
				}
			}

		/* Finishチャネルに動きがあった場合(ゲーム終了) */
		case user := <-n.Finish:
			msg := Message{
				Action: "finish",
				Value:  user.Name,
			}

			for _, p := range n.Players {
				if err := p.User.Socket.WriteJSON(msg); err != nil {
					// ここでもleave?
				}
			}
		}
	}
}
