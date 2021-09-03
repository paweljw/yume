package models

type ItemCategory uint

const (
	MiscItem = iota
)

type Item struct {
	ID          int64
	Name        string
	Description string
	Value       int64
	Category    ItemCategory
}
