package model

import "github.com/jinzhu/gorm"

type NumeronPlayer struct {
	gorm.Model
	NumeronId int `json:numeron_id`
	Numeron   Numeron
	UserId    int `json:user_id`
	User      User
	Result    int `json:result`
	Order     int `json:order`
}

type Result int
type Order int

const (
	Draw Result = iota // Draw == 0
	Win                // Win  == 1
	Lose               // Lose == 2
)
const (
	First  Order = iota // First  == 0
	Second              // Second == 1
)

func (r Result) String() string {
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
func (o Order) String() string {
	switch o {
	case First:
		return "先攻"
	case Second:
		return "後攻"
	}
	return "未定義"
}
