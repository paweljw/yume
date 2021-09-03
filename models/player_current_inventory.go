package models

import "gorm.io/gorm"

type PlayerCurrentInventory struct {
	gorm.Model
	PlayerId   int64
	Player     Player
	ItemId     int64
	Item       Item
	IsBound    bool `gorm:"default:false"`
	IsEquipped bool `gorm:"default:false"`
}
