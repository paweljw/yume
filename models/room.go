package models

var Rooms = make(map[uint64]Room)

type Room struct {
	ID          int64
	Description string
	ZoneId      int64
	Zone        Zone
}
