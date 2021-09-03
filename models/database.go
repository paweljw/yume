package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func EstablishConnection() error {
	// TODO: Configurable DSN
	dsn := "host=localhost user=postgres password=postgres dbname=yume_development port=5439 sslmode=disable"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	Db.AutoMigrate(
		&Zone{},
		&Room{},
		&Item{},
		&Container{},
		&ContainerInventory{},
		&RoomConnection{},
		&RoomContainer{},
		&Player{},
		&PlayerCurrentInventory{},
		&RoomCurrentInventory{},
	)

	return err
}
