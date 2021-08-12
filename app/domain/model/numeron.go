package model

import (
	"github.com/jinzhu/gorm"
)

type Numeron struct {
	gorm.Model
	Status  int            `json:"status"`
	OwnerId int            `json:"owner_id"`
	Owner   *User          `json:"-"`
	Join    chan *User     `json:"-"`
	Leave   chan *User     `json:"-"`
	Players map[*User]bool `json:"-"`
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

/*
ヌメロンルームを起動する
*/
func (n *Numeron) Run(user *User) {
	// 作成者を入室させる
	n.Players[user] = true

	for {
		// チャネルの動きを監視し、処理を決定する
		select {
		/* Joinチャネルに動きがあった場合(クライアントの入室) */
		case player := <-n.Join:
			// クライアントmapのbool値を真にする
			n.Players[player] = true

		/* Leaveチャネルに動きがあった場合(クライアントの退室) */
		case player := <-n.Leave:
			// クライアントmapから対象クライアントを削除する
			delete(n.Players, player)
		}
	}
}
