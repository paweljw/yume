package game

import (
	"log"
)

const TARGET_VERSION = 1

var Migrations = map[uint64]func(*Player) {
	0: migrateToVersion1,
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
