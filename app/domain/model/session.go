package model

import (
	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	SessionId string `json:session_id`
	Data      string `json:data`
	UserId    int    `json:user_id`
	User      *User
}
