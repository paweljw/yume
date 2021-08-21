package commands

import (
	cfg "yume/config"
	"yume/session"
	"yume/game"
	"strings"
	"log"
)

func handleNewSession(conn *session.Session) {
	conn.Tell(cfg.GetMessage("login_prompt"))

	command, _ := conn.Chomp()

	if command == "new" {
		conn.State = session.NewCharacter
	} else { // existing character
		if game.PlayerFileExists(command) {
			conn.Tell(cfg.GetMessage("provide_password"), command)
			player, _ := game.LoadPlayerFromFile(command)
			command, _ = conn.Chomp()
			if player.ComparePassword(command) {
				conn.Player = game.MigratePlayer(player)
				conn.State = session.Playing
				conn.Tell(cfg.GetMessage("welcome_back"), conn.Player.Name)
				log.Printf("Successful sign-in from %s", conn.Player.Name)
			} else {
				conn.State = session.NewSession
				log.Printf("Unsuccessful sign-in for %s", player.Name)
			}

		} else {
			conn.Tell(cfg.GetMessage("character_not_found"))
			conn.Tell(cfg.GetMessage("login_prompt"))
			log.Printf("Non-existing character login attempt: %s", command)
			conn.State = session.NewSession
		}
	}
}

func handleNewCharacter(conn *session.Session) {
	conn.Tell(cfg.GetMessage("name_prompt"))

	name, _ := conn.Chomp()

	if cfg.IsBadName(name) || game.PlayerFileExists(name) {
		conn.Tell(cfg.GetMessage("disallowed_name"))
	} else {
		conn.Player.Name = name
		conn.State = session.NewPassword
		conn.Tell(cfg.GetMessage("accepted_name"), conn.Player.Name)
	}
}

func handleNewPassword(conn *session.Session) {
	pass, _ := conn.Chomp()
	conn.Player.SetPassword(pass)
	conn.State = session.RepeatPassword
}

func handleRepeatPassword(conn *session.Session) {
	conn.Tell(cfg.GetMessage("repeat_password_prompt"))

	pass, _ := conn.Chomp()

	if conn.Player.ComparePassword(pass) {
		conn.State = session.SelectRace
	} else {
		conn.Tell(cfg.GetMessage("repeat_password_different"))
		conn.State = session.NewPassword
	}
}

func handleSelectRace(conn *session.Session) {
	conn.Tell(cfg.GetMessage("select_race_prompt"))
	race, _ := conn.Chomp()
	race = strings.ToLower(race)

	switch race {
	case "human":
		conn.Player.Race = game.Human
	case "elf":
		conn.Player.Race = game.Elf
	case "dwarf":
		conn.Player.Race = game.Dwarf
	default:
		conn.State = session.SelectRace
		return
	}

	err := conn.Player.SaveToFile()

	if err != nil {
		conn.State = session.NewSession
		conn.Tell("Something went very wrong. Please try again.")
		return
	}

	conn.Tell(cfg.GetMessage("race_selected"))
	log.Printf("New character registered: %s ((%s))", conn.Player.Name, conn.Player.Race)
	conn.State = session.Playing
}

var NonPlayingStates = map[session.SessionState]func(*session.Session){
	session.NewSession: handleNewSession,
	session.NewCharacter: handleNewCharacter,
	session.NewPassword: handleNewPassword,
	session.RepeatPassword: handleRepeatPassword,
	session.SelectRace: handleSelectRace,
}
