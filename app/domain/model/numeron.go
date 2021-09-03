package model

import (
	"github.com/jinzhu/gorm"
)

type Numeron struct {
	gorm.Model
	RoomId int   `json:"room_id"`
	Room   *Room `json:"-"`
}
