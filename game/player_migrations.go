package game

import (
	"log"
)

const TARGET_VERSION = 2

var Migrations = map[uint64]func(*Player) {
	0: migrateToVersion1,
	1: migrateToVersion2,
}

func MigratePlayer(player *Player) *Player {
	version := player.Version

	for version < TARGET_VERSION {
		Migrations[version](player)

		version = player.Version
		log.Printf("Migrated %s to version %d", player.Name, version)
	}

	player.SaveToFile()

	return player
}

func migrateToVersion1(player *Player) {
	player.Version = 1
}

func migrateToVersion2(player *Player) {
	player.RoomId = 100001
	player.Version = 2
}

func DefaultPlayer() *Player {
	player := Player{
		Name: "",
		Password: "",
		Race: Human,
		Flags: make(map[string]bool),
		Version: TARGET_VERSION,
		RoomId: 100001,
	}

	return &player
}
