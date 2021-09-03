package models

import "strings"

var Rooms map[int64]Room

type Room struct {
	ID              int64
	Description     string
	ZoneId          int64
	Zone            Zone
	RoomConnections []RoomConnection `gorm:"foreignKey:from_id"`
}

func LoadAllRooms() {
	rooms := []Room{}
	Db.Preload("RoomConnections").Find(&rooms)

	Rooms = make(map[int64]Room, len(rooms))

	for _, r := range rooms {
		r.Description = strings.ReplaceAll(r.Description, `\n`, "\n")
		Rooms[r.ID] = r
	}
}

func (room *Room) Exits() string {
	exits := []string{}

	for _, con := range room.RoomConnections {
		exits = append(exits, con.ShortDirection())
	}

	return strings.Join(exits, " ")
}

func (room *Room) Connections() map[Direction]int64 {
	rc := room.RoomConnections
	mp := make(map[Direction]int64, len(rc))

	for _, c := range rc {
		mp[c.Direction] = c.ToId
	}

	return mp
}
