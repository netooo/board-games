package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId   string `json:user_id`
	Name     string `json:name`
	Email    string `json:email`
	Password string `json:password`
}