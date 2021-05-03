package model

import "github.com/jinzhu/gorm"

type Numeron struct {
	gorm.Model
	Status int `json:status`
}

type Status int

const (
	Unknown Status = iota // Unknown  == 0
	Ready                 // Ready    == 1
	Play                  // Play     == 2
	Finish                // Finish   == 3
)

func (s Status) String() string {
	switch s {
	case Unknown:
		return "不明"
	case Ready:
		return "準備中"
	case Play:
		return "実行中"
	case Finish:
		return "終了"
	}
	return "未定義"
}
