package commands

import (
	"strings"
	"yume/game"
	ses "yume/session"
	"log"
)

func toggleFlag(session *ses.Session, command string, value bool) {
	if session.Player.HasFlag("admin") { // TODO: something like a sudo mode, once buffs are implemented
		tokens := strings.Split(command, " ")

		playerName := tokens[1]
		flagName := tokens[2]

		// Step 1: if player is signed on, let's find them and do them up
		doneOnline := false

		for e := ses.Sessions.Front(); e != nil; e = e.Next() {
			c := *e.Value.(*ses.Session)

			if c.Player.NameEquals(playerName) {
				c.Player.SetFlag(flagName, value)
				c.Player.SaveToFile()
				c.Tell("You were granted flag: %s", flagName)
				session.Tell("Granted flag %s to %s (online)", flagName, playerName)
				log.Printf("%s granted flag %s to %s", session.Player.Name, flagName, playerName)
				doneOnline = true
				break
			}
		}

		if !doneOnline {
			player, err := game.LoadPlayerFromFile(playerName)

			if err != nil {
				session.Tell("Something went wrong while granting flag")
				return
			}

			player.SetFlag(flagName, value)
			player.SaveToFile()
			session.Tell("Granted flag %s to %s (offline)", flagName, playerName)
			log.Printf("%s granted flag %s to %s", session.Player.Name, flagName, playerName)
		}
	} else {
		session.Tell("You're not allowed to do that.")
	}
}

func handleSetFlag(session *ses.Session, command string) {
	toggleFlag(session, command, true)
}

func handleUnsetFlag(session *ses.Session, command string) {
	toggleFlag(session, command, false)
}
