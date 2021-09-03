package models

type ContainerInventory struct {
	ID          int64
	ContainerId int64
	Container   Container
	ItemId      int64
	Item        Item
	Rate        float64 `gorm:"default:1.0"`
}
