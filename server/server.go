package main

import (
	"log"
	"math/rand"
	"net"
	"time"
	ntw "yume/network"
	ses "yume/session"
	cfg "yume/config"
	"yume/models"
	"yume/game"
)

func main() {
	log.Println("Yume is initializing")

	models.EstablishConnection()

	zone := models.Zone{ID: 1, Name: "hey", Description: "oh god no"}

	err := models.DB.Create(&zone)

	if err != nil {
		panic(err)
	}

	cfg.LoadConfiguration()
	log.Println("Loaded configuration file")

	log.Println("Begin loading rooms")
	game.LoadAllRooms()
	log.Println("Loaded all rooms")

	log.Println("Begin loading items")
	game.LoadAllItems()
	log.Println("Loaded all items")

	PORT := ":19000"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	log.Printf("Accepting connections on %s", PORT)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		// TODO: handle banned IPs
		connection := ses.Session{Connection: c, State: ses.NewSession}

		ses.Sessions.ConcurrentPushBack(&connection)

		go ntw.HandleSession(&connection)
	}
}
