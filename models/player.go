package models

import "gorm.io/gorm"

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
