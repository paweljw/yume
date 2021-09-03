package commands

import (
	"log"
	"strings"
	cfg "yume/config"
	"yume/models"
	"yume/session"
)

func handleNewSession(conn *session.Session) {
	conn.Tell(cfg.GetMessage("login_prompt"))

	command, _ := conn.Chomp()

	if command == "new" {
		conn.State = session.NewCharacter
	} else if command == "quit" {
		conn.Tell("So be it.")
		conn.Finishing = true
		return
	} else {
		player, found := models.FindPlayerByName(command)

		if found {
			conn.Tell(cfg.GetMessage("provide_password"), command)
			command, _ = conn.Chomp()

			if player.ComparePassword(command) {
				conn.Player = &player
				conn.State = session.Playing

				conn.Tell(cfg.GetMessage("welcome_back"), conn.Player.Name)
				MovePlayer(conn, conn.Player.SavedRoomId)

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

	player := models.Player{}
	models.Db.Where("name ILIKE ?", name).First(&player)

	if cfg.IsBadName(name) || player.ID != 0 {
		conn.Tell(cfg.GetMessage("disallowed_name"))
	} else {
		conn.Player = &models.Player{}
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
		conn.Player.Race = models.Human
	case "elf":
		conn.Player.Race = models.Elf
	case "dwarf":
		conn.Player.Race = models.Dwarf
	default:
		conn.State = session.SelectRace
		return
	}

	// Finalizing here
	conn.Player.CurrentRoomId = 1
	conn.Player.SavedRoomId = 1
	err := models.Db.Create(&conn.Player).Error

	if err != nil {
		conn.State = session.NewSession
		conn.Tell("Something went very wrong. Please try again.")
		return
	}

	conn.Tell(cfg.GetMessage("race_selected"))
	log.Printf("New character registered: %s ((%s))", conn.Player.Name, conn.Player.Race)
	conn.State = session.Playing
	MovePlayer(conn, conn.Player.CurrentRoomId)
}

var NonPlayingStates = map[session.SessionState]func(*session.Session){
	session.NewSession:     handleNewSession,
	session.NewCharacter:   handleNewCharacter,
	session.NewPassword:    handleNewPassword,
	session.RepeatPassword: handleRepeatPassword,
	session.SelectRace:     handleSelectRace,
}
