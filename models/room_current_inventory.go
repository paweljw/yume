package models

import "gorm.io/gorm"

type RoomCurrentInventory struct {
	gorm.Model
	RoomContainerId int64
	RoomContainer   RoomContainer
	ItemId          int64
	Item            Item
	VisibleToId     int64
	VisibleTo       Player
}
