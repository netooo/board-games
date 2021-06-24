package model

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

type NumeronPlayer struct {
	gorm.Model
	NumeronId int `json:numeron_id`
	Numeron   Numeron
	UserId    int `json:user_id`
	User      User
	Order     int    `json:order`
	Code      string `json:code`
	Rank      int    `json:rank`
	Socket    *websocket.Conn
}

type Order int
type Rank int

const (
	First  Order = iota // First  == 0
	Second              // Second == 1
)
const (
	Draw Rank = iota // Draw == 0
	Win              // Win  == 1
	Lose             // Lose == 2
)

func (o Order) String() string {
	switch o {
	case First:
		return "先攻"
	case Second:
		return "後攻"
	}
	return "未定義"
}
func (r Rank) String() string {
	switch r {
	case Draw:
		return "引き分け"
	case Win:
		return "勝利"
	case Lose:
		return "敗北"
	}
	return "未定義"
}
