package game

import (
	"yume/commons"
	"github.com/jinzhu/copier"
	"os"
	"encoding/json"
	"path/filepath"
	"log"
)

var Items = make(map[uint64]Item)

type Inventory struct {
	Contents	commons.SafeList
}

// Omitting the actual inventory - this will get created "manually" on load
// Omitting VisibleFor, as this will be decided on item creation (usually)

type Item struct {
	Id					uint64
	Name				string
	Value				uint64
	Description			string
	Container			bool
	Inventory			Inventory `json:"-"`
	InventoryListing 	[]uint64
	VisibleFor			[]string `json:"-"`
}

func NewItemFromId(id uint64) Item {
	item := Item{}
	copier.Copy(&item, Items[id])

	item.Inventory = Inventory{Contents: commons.NewSafeList()}

	// TODO: Use InventoryListing to create the actual inventory
	item.VisibleFor = make([]string, 1)

	return item
}

func NewItemFromIdFor(id uint64, players []string) Item {
	item := NewItemFromId(id)

	item.VisibleFor = players

	return item
}

func LoadItemFromFile(filename string) (Item, error) {
	item := Item{}
	jsonRepr, err := os.ReadFile(filename)

	if err != nil {
		return item, err
	}

	err = json.Unmarshal(jsonRepr, &item)

	return item, err
}

func LoadAllItems() {
	var files []string

	err := filepath.Walk("./resources/items", func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		log.Printf("Loading item from %s", file)
		item, err := LoadItemFromFile(file)

		if err != nil {
			panic(err)
		}

		Items[item.Id] = item
	}
}
