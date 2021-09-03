package models

import (
	"strconv"

	"gorm.io/gorm"
)

type RoomContainer struct {
	gorm.Model
	ContainerId uint
	Container   Container
	RoomId      uint
	Room        Room
}

func (rc *RoomContainer) CanSpawnFor(pl Player) bool {
	res, found, _ := Redis.HGet(
		strconv.FormatUint(uint64(rc.ID), 10)+"-spawns",
		strconv.FormatUint(uint64(pl.ID), 10),
	)

	if found {
		return false
	}

	return res != "t"
}

func (rc *RoomContainer) MarkSpawnedFor(pl Player) {
	Redis.HSet(
		strconv.FormatUint(uint64(rc.ID), 10)+"-spawns",
		strconv.FormatUint(uint64(pl.ID), 10),
		"t",
	)
}
