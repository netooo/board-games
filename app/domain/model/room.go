package model

import (
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
