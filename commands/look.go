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

func LookAtRoom(session *ses.Session, roomId uint) {
	room := models.Rooms[roomId]

	session.Tell("\n%s", room.Description)
	session.Tell("\nExits: %s", room.Exits())

	rcis := []models.RoomCurrentInventory{}

	models.Db.Where("room_id = ? AND visible_to_id = ?", roomId, session.Player.ID).Preload("Room").Preload("Container").Preload("Item").Order("container_id desc").Find(&rcis)

	var lastContainer models.Container

	// TODO: a/an detection
	for _, rci := range rcis {
		if rci.Container.IsFloor {
			session.Tell("There is a/an %s on the floor.", rci.Item.Name)
		} else {
			if lastContainer.ID != rci.ContainerId {
				session.Tell("There is a/an %s here.", rci.Container.Name)
				lastContainer = rci.Container
			}
		}
	}
}
