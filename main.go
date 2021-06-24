package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	ntw "yume/network"
)

func main() {
	PORT := ":19000"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		// TODO: handle banned IPs
		connection := ntw.Connection{Connection: c, State: ntw.NewConnection}
		ntw.Connections.PushBack(connection)
		go connection.HandleConnection()
	}
}
