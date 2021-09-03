package models

import "database/sql"

type Direction int

const (
	North = iota
	East
	South
	West
	NorthEast
	SouthEast
	SouthWest
	NorthWest
	Up
	Down
)

type RoomConnection struct {
	Id           int64
	FromId       int64
	From         Room
	ToId         int64
	To           Room
	Direction    Direction
	LockedById   sql.NullInt64
	LockedBy     Item
	LockedByFlag string
}
