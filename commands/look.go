package commands

import (
	"yume/game"
	ses "yume/session"
)

// TODO: allow looking at things, once things exist
func handleLook(session *ses.Session, command string) {
	session.Tell("You look around.")

	LookAtRoom(session, session.Player.RoomId)
}

func LookAtRoom(session *ses.Session, roomId uint64) {
	room := game.Rooms[roomId]

	session.Tell(room.Description)
	session.Tell("\nExits: %s", room.Exits())
}
