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

func handleNewCharacterState(conn *Connection, command string) {
	switch conn.State {

	case NewCharacter:
		if cfg.IsBadName(command) {
			conn.prompt(cfg.GetMessage("disallowed_name"))
		} else { // TODO: Handle existing name
			conn.Player.Name = command
			conn.State = NewPassword
			conn.prompt(cfg.GetMessage("accepted_name"), command)
		}

	case NewPassword:
		conn.Player.SetPassword(command)
		conn.State = RepeatPassword
		conn.prompt(cfg.GetMessage("repeat_password_prompt"))
	case RepeatPassword:
		if conn.Player.ComparePassword(command) {
			conn.State = SelectRace
			conn.prompt(cfg.GetMessage("select_race_prompt"))
		} else {
			conn.State = NewPassword
			conn.prompt(cfg.GetMessage("repeat_password_different"))
		}
	case SelectRace:
		race := strings.ToLower(command)

		switch race {
		case "human":
			conn.Player.Race = game.Human
		case "elf":
			conn.Player.Race = game.Elf
		case "dwarf":
			conn.Player.Race = game.Dwarf
		default:
			conn.State = SelectRace
			conn.prompt(cfg.GetMessage("select_race_prompt"))

			return
		}

		err := conn.Player.SaveToFile()

		if err != nil {
			conn.State = NewConnection
			conn.prompt(cfg.GetMessage("login_prompt"))
			return
		}

		conn.State = Playing
		conn.tell(cfg.GetMessage("race_selected"))
		conn.prompt("")
	}
}

func handleExistingCharacter(conn *Connection, command string) {
	switch conn.State {
	case ExistingCharacter:
		if game.PlayerFileExists(command) {
			conn.prompt(cfg.GetMessage("provide_password"), command)
			conn.State = ExistingCharacterPassword
			player, _ := game.LoadPlayerFromFile(command)
			conn.Player = player
		} else {
			conn.tell(cfg.GetMessage("character_not_found"))
			conn.prompt(cfg.GetMessage("login_prompt"))
			conn.State = NewConnection
		}
	case ExistingCharacterPassword:
		if conn.Player.ComparePassword(command) {
			conn.State = Playing
			conn.tell(cfg.GetMessage("welcome_back"), conn.Player.Name)
		} else {
			conn.Player = new(game.Player)
			conn.State = NewConnection
			conn.prompt(cfg.GetMessage("login_prompt"))
		}
	}
}

type Connection struct {
	Connection net.Conn
	State      ConnectionState
	Player	   *game.Player
}

func (conn *Connection) HandleConnection() {
	log.Printf("Serving %s\n", conn.Connection.RemoteAddr().String())
	conn.Player = new(game.Player)

	conn.tell(cfg.GetMessage("motd"))
	conn.prompt(cfg.GetMessage("login_prompt"))

	for {
		netData, err := bufio.NewReader(conn.Connection).ReadString('\n')
		if err != nil {
			log.Println(err)
			TellPlayer(*conn, "Something went very wrong. Please sign in again.")
			break
		}
		conn.tell("\n") // Send a single newline as a keepalive

		command := strings.TrimSpace(netData)

		// `quit` is a top-level command, available _everywhere_.
		if command == "quit" {
			TellPlayer(*conn, "Well, bye for now.")
			break
		}

		// refactor with option to loop through state updates here
		if isNewCharacterState(conn.State) {
			handleNewCharacterState(conn, command)
		} else if conn.State == ExistingCharacterPassword {
			handleExistingCharacter(conn, command)
		} else if conn.State == NewConnection {
			switch command {
			case "new":
				conn.State = NewCharacter
				conn.prompt(cfg.GetMessage("name_prompt"))
			default:
				conn.State = ExistingCharacter
				handleExistingCharacter(conn, command)
			}
		} else if conn.State == Playing {
			// TODO: Eventually take over with the main command map here
			conn.prompt("Well, nice of you to drop by. You can go ahead and ``quit``.")
		} else {
			conn.State = NewConnection
			TellPlayer(*conn, "Damn, how'd you even GET here? Bootin' ya.")
			break
		}
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


func (conn *Connection) tell(msg string, a...interface{}) {
	TellPlayer(*conn, msg, a...)
}

func (conn *Connection) prompt(msg string, a...interface{}) {
	PromptPlayer(*conn, msg, a...)
}
