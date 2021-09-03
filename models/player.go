package models

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"gorm.io/gorm"
)

type Race int

const (
	Human = iota
	Elf
	Dwarf
)

type Pronouns int

const (
	HeHim = iota
	SheHer
	TheyThem
)

type Player struct {
	gorm.Model
	Name          string
	Password      string
	Race          Race
	Pronouns      Pronouns
	SavedRoomId   int64
	SavedRoom     Room
	CurrentRoomId int64
	CurrentRoom   Room
}

func FindPlayerByName(name string) (Player, bool) {
	player := Player{}

	Db.Where("name ILIKE ?", name).First(&player)

	return player, player.ID != 0
}

func (player *Player) ComparePassword(unsecurePassword string) bool {
	hasher := sha256.New()
	hasher.Write([]byte(unsecurePassword))
	return player.Password == hex.EncodeToString(hasher.Sum(nil))
}

func (player *Player) SetPassword(unsecurePassword string) {
	hasher := sha256.New()
	hasher.Write([]byte(unsecurePassword))
	player.Password = hex.EncodeToString(hasher.Sum(nil))
}

func (player *Player) HasFlag(flag string) bool {
	res, found, _ := Redis.HGet(
		strconv.FormatUint(
			uint64(player.ID), 10)+"-flags",
		flag,
	)

	if !found {
		return false
	}

	return res == "t"
}

func (player *Player) SetFlag(flag string, value bool) {
	var strValue string

	if value {
		strValue = "t"
	} else {
		strValue = "f"
	}

	_, _ = Redis.HSet(strconv.FormatUint(uint64(player.ID), 10)+"-flags", flag, strValue)
}
