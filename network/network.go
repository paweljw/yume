package network

import (
	"log"
	"net"
	"bufio"
	"strings"
	"container/list"
	cfg "yume/config"
	"yume/game"
	"time"
)

var Connections = list.New()

type Connection struct {
	Connection net.Conn
	State      ConnectionState
	Player	   *game.Player
}

func (conn *Connection) HandleConnection() {
	log.Printf("Serving %s\n", conn.Connection.RemoteAddr().String())
	conn.Player = new(game.Player)

	conn.tell(cfg.GetMessage("motd"))
	conn.State = NewConnection

	for {
		if conn.State == Playing {
			conn.tell("We'd be applying the main command map here.")
			command, _ := conn.chomp()

			if command == "quit" {
				break
			}
		} else { // this covers non playing states
			switch conn.State {
			case NewConnection:
				conn.prompt(cfg.GetMessage("login_prompt"))

				command, _ := conn.chomp()

				if command == "new" {
					conn.State = NewCharacter
					continue
				} else { // existing character
					if game.PlayerFileExists(command) {
						conn.prompt(cfg.GetMessage("provide_password"), command)
						player, _ := game.LoadPlayerFromFile(command)
						command, _ = conn.chomp()
						if player.ComparePassword(command) {
							conn.tell(cfg.GetMessage("welcome_back"), conn.Player.Name)
							conn.Player = player
							conn.State = Playing
							continue
						} else {
							conn.State = NewConnection
						}

					} else {
						conn.tell(cfg.GetMessage("character_not_found"))
						conn.prompt(cfg.GetMessage("login_prompt"))
						conn.State = NewConnection
					}
				}

			case NewCharacter:
				conn.prompt(cfg.GetMessage("name_prompt"))

				name, _ := conn.chomp()

				if cfg.IsBadName(name) || game.PlayerFileExists(name) {
					conn.tell(cfg.GetMessage("disallowed_name"))
					continue
				} else {
					conn.Player.Name = name
					conn.State = NewPassword
					conn.prompt(cfg.GetMessage("accepted_name"), conn.Player.Name)
					continue
				}
			case NewPassword:
				pass, _ := conn.chomp()
				conn.Player.SetPassword(pass)
				conn.State = RepeatPassword
				continue
			case RepeatPassword:
				conn.prompt(cfg.GetMessage("repeat_password_prompt"))
				pass, _ := conn.chomp()
				if conn.Player.ComparePassword(pass) {
					conn.State = SelectRace
					continue
				} else {
					conn.prompt(cfg.GetMessage("repeat_password_different"))
					conn.State = NewPassword
					continue
				}
			case SelectRace:
				conn.prompt(cfg.GetMessage("select_race_prompt"))
				race, _ := conn.chomp()
				race = strings.ToLower(race)

				switch race {
				case "human":
					conn.Player.Race = game.Human
				case "elf":
					conn.Player.Race = game.Elf
				case "dwarf":
					conn.Player.Race = game.Dwarf
				default:
					conn.State = SelectRace
					continue
				}

				err := conn.Player.SaveToFile()

				if err != nil {
					conn.State = NewConnection
					continue
				}

				conn.tell(cfg.GetMessage("race_selected"))
				conn.State = Playing
				continue
			default:
				conn.tell("How'd you even get here, bro? Bye.")
				break
			}
		}
	}


	conn.tell("Abayo.")
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
