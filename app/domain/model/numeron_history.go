package model

import "github.com/jinzhu/gorm"

type NumeronHistory struct {
	gorm.Model
	NumeronId     uint `json:numeron_id`
	Numeron       *Numeron
	PlayerId      uint `json:player_id`
	NumeronPlayer *NumeronPlayer
	Code          string `json:code`
	Result        string `json:result`
	Turn          int    `json:turn`
}
