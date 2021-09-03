package models

import (
	"github.com/garyburd/redigo/redis"
	"github.com/shomali11/xredis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Redis *xredis.Client

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

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	Redis = xredis.NewClient(pool)

	return nil
}
