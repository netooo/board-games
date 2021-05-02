package model

import "github.com/jinzhu/gorm"

type Numeron struct {
	gorm.Model
	Status int `json:status`
}
