package models

type ItemCategory uint

var Items map[int64]Item

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

func LoadAllItems() {
	items := []Item{}
	Db.Find(&items)

	Items = make(map[int64]Item, len(items))

	for _, i := range items {
		Items[i.ID] = i
	}
}
