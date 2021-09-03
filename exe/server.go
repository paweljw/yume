package main

import (
	"log"
	"math/rand"
	"net"
	"time"
	cfg "yume/config"
	"yume/game"
	"yume/models"
	ntw "yume/network"
	ses "yume/session"

	cli "github.com/urfave/cli/v2"
)

func server(_ *cli.Context) error {
	log.Println("Yume is initializing")

	err := models.EstablishConnection()

	if err != nil {
		log.Fatal(err)
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
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		return err
	}

	defer l.Close()
	rand.Seed(time.Now().Unix())

	log.Printf("Accepting connections on %s", PORT)

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}
		// TODO: handle banned IPs
		connection := ses.Session{Connection: c, State: ses.NewSession}

		ses.Sessions.ConcurrentPushBack(&connection)

		go ntw.HandleSession(&connection)
	}
}
