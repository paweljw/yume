package game

import (
	"os"
	"encoding/json"
	"path/filepath"
	"log"
	"strings"
)

var Rooms = make(map[uint64]Room)

type Room struct {
	Id uint64
	Zone string
	CarpenterTag string
	Description string
	Connections map[string]uint64
}

func (room *Room) Exits() string {
	keys := make([]string, len(room.Connections))

	i := 0
	for k := range room.Connections {
		keys[i] = k
		i++
	}

	return strings.Join(keys, ", ")
}

func LoadRoomFromFile(filename string) (Room, error) {
	room := Room{}

	jsonRepr, err := os.ReadFile(filename)

	if err != nil {
		return room, err
	}

	err = json.Unmarshal(jsonRepr, &room)

	return room, err
}

func LoadAllRooms() {
	var files []string

	err := filepath.Walk("./resources/rooms", func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		log.Printf("Loading room from %s", file)
		room, err := LoadRoomFromFile(file)

		if err != nil {
			panic(err)
		}

		Rooms[room.Id] = room
	}
}
