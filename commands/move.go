package commands

import (
	"log"
	"strings"
	"yume/game"
	ses "yume/session"
)

func handleMovement(session *ses.Session, command string) {
	// TODO: check if player can move (debuffs, fight, jail, etc)
	currentRoom := game.Rooms[uint64(session.Player.CurrentRoomId)]
	direction := normalizeDirection(command)

	nextRoom, exists := currentRoom.Connections[direction]

	if !exists {
		session.Tell("There is no exit %s", direction)
		return
	}

	MovePlayer(session, nextRoom)
}

func MovePlayer(session *ses.Session, toRoom uint64) {
	if room, ok := game.Rooms[toRoom]; ok {
		log.Printf("%s moved from %d to %d", session.Player.Name, session.Player.CurrentRoomId, room.Id)
		session.Player.CurrentRoomId = room.ID
		LookAtRoom(session, room.Id)
	} else {
		session.Tell("That's disallowed - nonexistent room.")
		log.Printf("%s attempted to access nonexistent room %d", session.Player.Name, toRoom)
	}

}

func normalizeDirection(direction string) string {
	split := strings.Split(direction, " ")

	switch split[0] {
	case "w":
		return "west"
	case "west":
		return "west"
	case "e":
		return "east"
	case "east":
		return "east"
	case "n":
		return "north"
	case "north":
		return "north"
	case "s":
		return "south"
	case "south":
		return "south"
	case "ne":
		return "northeast"
	case "northeast":
		return "northeast"
	case "se":
		return "southeast"
	case "southeast":
		return "southeast"
	case "sw":
		return "southwest"
	case "southwest":
		return "southwest"
	case "nw":
		return "northwest"
	case "northwest":
		return "northwest"
	case "u":
		return "up"
	case "up":
		return "up"
	case "d":
		return "down"
	case "down":
		return "down"
	default:
		if len(split) > 1 {
			return normalizeDirection(strings.Join(split[1:], " "))
		} else {
			return "invalid"
		}
	}
}
