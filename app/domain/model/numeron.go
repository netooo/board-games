package model

import (
	_ "github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

type Numeron struct {
	gorm.Model
	Status  int `json:status`
	Join    chan *NumeronPlayer
	Leave   chan *NumeronPlayer
	Players map[*NumeronPlayer]bool
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
func (n *Numeron) Run() {
	for {
		// チャネルの動きを監視し、処理を決定する
		select {

		/* joinチャネルに動きがあった場合(プレイヤーの入室) */
		case client := <-n.Join:
			// プレイヤーmapのbool値を真にする
			n.Players[client] = true
		}
	}
}
