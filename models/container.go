package models

type Container struct {
	ID                   uint
	Name                 string
	Description          string
	IsFloor              bool
	ContainerInventories []ContainerInventory
}
