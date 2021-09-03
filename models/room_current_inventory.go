package models

import "gorm.io/gorm"

type RoomCurrentInventory struct {
	gorm.Model
	RoomId          uint
	Room            Room
	ContainerId     uint
	Container       Container
	RoomContainerId uint
	RoomContainer   RoomContainer
	ItemId          uint
	Item            Item
	VisibleToId     uint
	VisibleTo       Player
}
