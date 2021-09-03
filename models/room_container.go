package models

type RoomContainer struct {
	Id          int64
	ContainerId int64
	Container   Container
	RoomId      int64
	Room        Room
}
