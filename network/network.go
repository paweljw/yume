package network

import (
	"log"
	"net"
	"bufio"
	"strings"
	"container/list"
	cfg "yume/config"
	"yume/game"
)

var Connections = list.New()

type Connection struct {
	Connection net.Conn
	State      ConnectionState
	Player	   *game.Player
	Finishing  bool
}

func (conn *Connection) HandleConnection() {
	log.Printf("Serving %s\n", conn.Connection.RemoteAddr().String())
	conn.Player = new(game.Player)
	conn.Finishing = false

	conn.tell(cfg.GetMessage("motd"))
	conn.State = NewConnection

	for {
		if conn.State == Playing {
			conn.tell("We'd be applying the main command map here.")
			command, _ := conn.chomp()

			actionWord := strings.ToLower(strings.Split(command, " ")[0])

			handler := commandMap[actionWord]

			if handler != nil {
				handler(conn, command)
			} else {
				conn.tell("What?")
			}

		} else { // this covers non playing states
			handler := nonPlayingStates[conn.State]

			if handler != nil {
				handler(conn)
				continue
			} else {
				conn.tell("How'd you even get here, bro? Bye.")
				break
			}
		}

		if conn.Finishing {
			break
		}
	}

	if conn.Player.IsSaveable() {
		conn.tell("Saving character...")
		conn.Player.SaveToFile()
	}

	log.Printf("Finished serving %s\n", conn.Connection.RemoteAddr().String())
	conn.Connection.Close()
	conn.removeFromList()
}

func (conn *Connection) chomp() (string, error) {
	netData, err := bufio.NewReader(conn.Connection).ReadString('\n')
	if err != nil {
		log.Println(err)
		TellPlayer(*conn, "Something went very wrong. Please sign in again.")
		return "", err
	}
	conn.tell("\n") // Send a single newline as a keepalive

	command := strings.TrimSpace(netData)

	return command, nil
}

func (conn *Connection) removeFromList() {
	for e := Connections.Front(); e != nil; e = e.Next() {
		if e.Value == conn {
			Connections.Remove(e)
			return
		}
	}
}


func (conn *Connection) tell(msg string, a...interface{}) {
	TellPlayer(*conn, msg, a...)
}

func (conn *Connection) prompt(msg string, a...interface{}) {
	PromptPlayer(*conn, msg, a...)
}
