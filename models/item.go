package models

type ItemCategory uint

var Items map[uint]Item

const (
	MiscItem = iota
)

type Item struct {
	ID          uint
	Name        string
	Description string
	Value       int64
	Category    ItemCategory
}

func LoadAllItems() {
	items := []Item{}
	Db.Find(&items)

	Items = make(map[uint]Item, len(items))

	for _, i := range items {
		Items[i.ID] = i
	}
}
