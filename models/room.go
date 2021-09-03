package models

type Room struct {
	ID          int64
	Description string
	ZoneId      int64
	Zone        Zone
}
