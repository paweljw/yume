package models

import "gorm.io/gorm"

type PlayerCurrentInventory struct {
	gorm.Model
	PlayerId   uint
	Player     Player
	ItemId     uint
	Item       Item
	IsBound    bool `gorm:"default:false"`
	IsEquipped bool `gorm:"default:false"`
}
