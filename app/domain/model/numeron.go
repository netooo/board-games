package model

import (
	"github.com/jinzhu/gorm"
)

type Numeron struct {
	gorm.Model
	Status  int `json:status`
	OwnerId int `json:owner_id`
	Owner   *NumeronPlayer
	Players *[]NumeronPlayer
}

var players = make(map[*NumeronPlayer]bool)
var join = make(chan *NumeronPlayer)

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
func (n *Numeron) Run(player *NumeronPlayer) {
	players[player] = true
	for {
		//TODO: 部屋の常時起動
	}
}

func (n *Numeron) Join(player *NumeronPlayer) {
	players[player] = true
	// 現在接続しているクライアント全てに入室を通知する
	for p := range players {
		//TODO: socketを利用して、playersにplayerが入室したことを通知する
	}
}
