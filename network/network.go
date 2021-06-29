package network

import (
	"log"
	"net"
	"bufio"
	"strings"
	"container/list"
	cfg "yume/config"
)

var Connections = list.New()

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
	Connection net.Conn
	State      ConnectionState
}

func (conn *Connection) HandleConnection() {
	log.Printf("Serving %s\n", conn.Connection.RemoteAddr().String())
	conn.Connection.Write([]byte(cfg.GetMessage("motd")))

	TellEveryone(cfg.GetMessage("player_online"))

	for {
		netData, err := bufio.NewReader(conn.Connection).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		log.Println(temp)
		if temp == "quit" {
			break
		}

		conn.Connection.Write([]byte("Generic response\n"))
	}

	log.Printf("Finished serving %s\n", conn.Connection.RemoteAddr().String())
	conn.Connection.Close()
	conn.removeFromList()
}

func (conn *Connection) removeFromList() {
	for e := Connections.Front(); e != nil; e = e.Next() {
		if e.Value == conn {
			Connections.Remove(e)
			return
		}
	}
}
