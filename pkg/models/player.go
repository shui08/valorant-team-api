package models

import (
	"github.com/shui08/valorant-team-api/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Player struct {
	IRLName        string  `json:"irlname"`
	RiotID         string  `json:"riotid"`
	Rank           string  `json:"rank"`
	Role           string  `json:"role"`
	Main           string  `json:"main"`
	ACS            float64 `json:"acs"`
	KDR            float64 `json:"kdr"`
	DamagePerRound float64 `json:"dmg/round"`
	HS             float64 `json:"hs%"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Player{})
	if err != nil {
		panic(err)
	}
}
