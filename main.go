package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	conn "yume/connection"
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
		connection := conn.Connection{Connection: c, State: conn.NewConnection}
		conn.Connections.PushBack(connection)
		go connection.HandleConnection()
	}
}
