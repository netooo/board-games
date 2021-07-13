package model

import "github.com/jinzhu/gorm"

type NumeronHistory struct {
	gorm.Model
	NumeronId     int `json:numeron_id`
	Numeron       *Numeron
	PlayerId      int `json:player_id`
	NumeronPlayer *NumeronPlayer
	Action        string `json:action`
	Result        string `json:result`
	Turn          int    `json:turn`
}
