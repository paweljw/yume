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
	Id           uint
	FromId       uint
	From         Room
	ToId         uint
	To           Room
	Direction    Direction
	LockedById   sql.NullInt64
	LockedBy     Item
	LockedByFlag string
}

func (rc *RoomConnection) ShortDirection() string {
	switch rc.Direction {
	case North:
		return "n"
	case East:
		return "e"
	case South:
		return "s"
	case West:
		return "w"
	case NorthEast:
		return "ne"
	case SouthEast:
		return "se"
	case SouthWest:
		return "sw"
	case NorthWest:
		return "nw"
	case Up:
		return "u"
	case Down:
		return "d"
	default:
		return "x"
	}
}

func StringToDirection(s string) Direction {
	switch s {
	case "w":
		return West
	case "west":
		return West
	case "e":
		return East
	case "east":
		return East
	case "n":
		return North
	case "north":
		return North
	case "s":
		return South
	case "south":
		return South
	case "ne":
		return NorthEast
	case "northeast":
		return NorthEast
	case "se":
		return SouthEast
	case "southeast":
		return SouthEast
	case "sw":
		return SouthWest
	case "southwest":
		return SouthWest
	case "nw":
		return NorthWest
	case "northwest":
		return NorthWest
	case "u":
		return Up
	case "up":
		return Up
	case "d":
		return Down
	case "down":
		return Down
	}

	return West
}
