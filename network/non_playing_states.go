package network

import (
	cfg "yume/config"
	"yume/game"
	"strings"
)

func handleNewConnection(conn *Connection) {
	conn.prompt(cfg.GetMessage("login_prompt"))

	command, _ := conn.chomp()

	if command == "new" {
		conn.State = NewCharacter
	} else { // existing character
		if game.PlayerFileExists(command) {
			conn.prompt(cfg.GetMessage("provide_password"), command)
			player, _ := game.LoadPlayerFromFile(command)
			command, _ = conn.chomp()
			if player.ComparePassword(command) {
				conn.Player = player
				conn.State = Playing
				conn.tell(cfg.GetMessage("welcome_back"), conn.Player.Name)
			} else {
				conn.State = NewConnection
			}

		} else {
			conn.tell(cfg.GetMessage("character_not_found"))
			conn.prompt(cfg.GetMessage("login_prompt"))
			conn.State = NewConnection
		}
	}
}

func handleNewCharacter(conn *Connection) {
	conn.prompt(cfg.GetMessage("name_prompt"))

	name, _ := conn.chomp()

	if cfg.IsBadName(name) || game.PlayerFileExists(name) {
		conn.tell(cfg.GetMessage("disallowed_name"))
	} else {
		conn.Player.Name = name
		conn.State = NewPassword
		conn.prompt(cfg.GetMessage("accepted_name"), conn.Player.Name)
	}
}

func handleNewPassword(conn *Connection) {
	pass, _ := conn.chomp()
	conn.Player.SetPassword(pass)
	conn.State = RepeatPassword
}

func handleRepeatPassword(conn *Connection) {
	conn.prompt(cfg.GetMessage("repeat_password_prompt"))

	pass, _ := conn.chomp()

	if conn.Player.ComparePassword(pass) {
		conn.State = SelectRace
	} else {
		conn.prompt(cfg.GetMessage("repeat_password_different"))
		conn.State = NewPassword
	}
}

func handleSelectRace(conn *Connection) {
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
		return
	}

	err := conn.Player.SaveToFile()

	if err != nil {
		conn.State = NewConnection
		return
	}

	conn.tell(cfg.GetMessage("race_selected"))
	conn.State = Playing
}

var nonPlayingStates = map[ConnectionState]func(*Connection){
	NewConnection: handleNewConnection,
	NewCharacter: handleNewCharacter,
	NewPassword: handleNewPassword,
	RepeatPassword: handleRepeatPassword,
	SelectRace: handleSelectRace,
}
