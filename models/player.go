package models

import (
	"crypto/sha256"
	"encoding/hex"

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
