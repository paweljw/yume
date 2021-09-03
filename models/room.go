package models

import (
	"log"
	"math/rand"
	"strings"
)

var Rooms map[uint]Room

type Room struct {
	ID              uint
	Description     string
	ZoneId          uint
	Zone            Zone
	RoomConnections []RoomConnection `gorm:"foreignKey:from_id"`
	RoomContainers  []RoomContainer
}

func LoadAllRooms() {
	rooms := []Room{}
	Db.Preload("RoomConnections").Preload("RoomContainers").Find(&rooms)

	Rooms = make(map[uint]Room, len(rooms))

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

func (room *Room) Connections() map[Direction]uint {
	rc := room.RoomConnections
	mp := make(map[Direction]uint, len(rc))

	for _, c := range rc {
		mp[c.Direction] = c.ToId
	}

	return mp
}

func (room *Room) SpawnContainersFor(pl Player) {
	for _, rc := range room.RoomContainers {
		if rc.CanSpawnFor(pl) {
			container := Container{}
			Db.Preload("ContainerInventories").Find(&container, rc.ContainerId)

			for _, item := range container.ContainerInventories {
				rci := RoomCurrentInventory{
					RoomId:          room.ID,
					ContainerId:     container.ID,
					RoomContainerId: rc.ID,
					ItemId:          item.ItemId,
					VisibleToId:     pl.ID,
				}

				Db.Find(&rci.Item, rci.ItemId)

				log.Printf("Maybe spawning %s (%d) for %s in room %d", rci.Item.Name, rci.Item.ID, pl.Name, room.ID)

				random := rand.Float64()

				if random <= item.Rate {
					Db.Save(&rci)
					log.Printf("Definitely spawning %s (%d) for %s in room %d (%.2f/%.2f)", rci.Item.Name, rci.Item.ID, pl.Name, room.ID, random, item.Rate)
				}
			}

			rc.MarkSpawnedFor(pl)
		} else {
			log.Printf("Skipping container %d for player %s", rc.ID, pl.Name)
		}
	}

}
