package model

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	DisplayId string          `json:"display_id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	Game      string          `json:"-"`
	Socket    *websocket.Conn `json:"-"`
}
