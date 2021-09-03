package models

type ContainerInventory struct {
	ID          uint
	ContainerId uint
	Container   Container
	ItemId      uint
	Item        Item
	Rate        float64 `gorm:"default:1.0"`
}
