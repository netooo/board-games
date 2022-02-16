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

type JoinMessage struct {
	Action string
	Name   string
}

type LeaveMessage struct {
	Action string
	Name   string
}

type StartMessage struct {
	Action string
	Name   string
	UserId string
}

type SetMessage struct {
	Action string
	UserId string
}

type FinishMessage struct {
	Action string
	Name   string
}

type AttackMessage struct {
	Action     string
	AttackUser string
	UserId     string
	Code       string
	Result     string
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
			msg := JoinMessage{
				Action: "join",
				Name:   user.Name,
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
			msg := LeaveMessage{
				Action: "leave",
				Name:   user.Name,
			}
			for u := range n.Users {
				if err := u.Socket.WriteJSON(msg); err != nil {
					u.Game = ""
					delete(n.Users, u)
					// ここでもleave?
				}
			}

		case user := <-n.Start:
			msg := StartMessage{
				Action: "start",
				Name:   user.Name,
				UserId: user.DisplayId,
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
				for _, p := range n.Players {
					msg := SetMessage{
						Action: "completed_code",
						UserId: p.User.DisplayId,
					}

					if err := p.User.Socket.WriteJSON(msg); err != nil {
						// ここでもleave?
					}
				}
			} else {
				msg := SetMessage{
					Action: "set_code",
					UserId: "",
				}

				if err := user.Socket.WriteJSON(msg); err != nil {
					// ここでもleave?
				}
			}

		/* Attackチャネルに動きがあった場合(コードの宣言) */
		case user := <-n.Attack:
			var attackUser, code, result string
			for _, p := range n.Players {
				if p.UserId == user.ID {
					attackUser = p.User.DisplayId
					code = p.Attack
					result = p.Result
					break
				}
			}

			for _, p := range n.Players {
				msg := AttackMessage{
					Action:     "attack",
					AttackUser: attackUser,
					UserId:     p.User.DisplayId,
					Code:       code,
					Result:     result,
				}

				if err := p.User.Socket.WriteJSON(msg); err != nil {
					// ここでもleave?
				}
			}

		/* Finishチャネルに動きがあった場合(ゲーム終了) */
		case user := <-n.Finish:
			msg := FinishMessage{
				Action: "finish",
				Name:   user.Name,
			}

			for _, p := range n.Players {
				if err := p.User.Socket.WriteJSON(msg); err != nil {
					// ここでもleave?
				}
			}
		}
	}
}
