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

	TellPlayer(*conn, cfg.GetMessage("motd"))
	TellEveryone(cfg.GetMessage("player_online"), 13)

	for {
		netData, err := bufio.NewReader(conn.Connection).ReadString('\n')
		if err != nil {
			log.Println(err)
			return // TODO: Better (or actual) handling
		}

		temp := strings.TrimSpace(string(netData))
		log.Println(temp)
		if temp == "quit" {
			break
		}

		TellPlayer(*conn, "Generic response")
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
