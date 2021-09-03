package commands

import (
	"yume/models"
	ses "yume/session"
)

// TODO: allow looking at things, once things exist
func handleLook(session *ses.Session, command string) {
	session.Tell("You look around.")

	LookAtRoom(session, session.Player.CurrentRoomId)
}

func LookAtRoom(session *ses.Session, roomId int64) {
	room := models.Rooms[roomId]

	session.Tell("\n%s", room.Description)
	session.Tell("\nExits: %s", room.Exits())
}
