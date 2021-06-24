package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	"yume/color"
)

const MIN = 1
const MAX = 100

var connections = list.New()

type ConnectionState int

const (
	NewConnection ConnectionState = iota
	NewCharacter
	NewPassword
	RepeatPassword
	SelectRace
	ExistingCharacter
	ExistingCharacterPassword
	Playing
)

type Connection struct {
	connection net.Conn
	state      ConnectionState
}

func (conn *Connection) handleConnection() {
	log.Printf("Serving %s\n", conn.connection.RemoteAddr().String())
	conn.connection.Write([]byte("Hello, client!\n"))
	tellEveryone("A new challenger appears!\n")

	for {
		netData, err := bufio.NewReader(conn.connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "quit" {
			break
		}
		fmt.Println(temp)

		conn.connection.Write([]byte("Generic response\n"))
	}

	log.Printf("Finished serving %s\n", conn.connection.RemoteAddr().String())
	conn.connection.Close()
	conn.removeFromList()
}

func (conn *Connection) removeFromList() {
	for e := connections.Front(); e != nil; e = e.Next() {
		if e.Value == conn {
			connections.Remove(e)
			return
		}
	}
}

func tellEveryone(s string) {
	for e := connections.Front(); e != nil; e = e.Next() {
		e.Value.(Connection).connection.Write([]byte(color.Colorize(s, color.Cyan)))
	}
}

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
		connection := Connection{connection: c, state: NewConnection}
		connections.PushBack(connection)
		go connection.handleConnection()
	}
}
