package commands

import (
	"container/list"
	"log"
	"strings"
	"yume/models"
	ses "yume/session"
)

func toggleFlag(session *ses.Session, command string, value bool) {
	if session.Player.HasFlag("admin") { // TODO: something like a sudo mode, once buffs are implemented
		tokens := strings.Split(command, " ")

		playerName := tokens[1]
		flagName := tokens[2]

		// Step 1: if player is signed on, let's find them and do them up
		doneOnline := false

		ses.Sessions.ConcurrentIterate(func(e *list.Element) bool {
			c := *e.Value.(*ses.Session)

			if strings.EqualFold(c.Player.Name, playerName) {
				c.Player.SetFlag(flagName, value)
				c.Tell("You were granted flag: %s", flagName)
				session.Tell("Granted flag %s to %s (online)", flagName, playerName)
				log.Printf("%s granted flag %s to %s", session.Player.Name, flagName, playerName)
				doneOnline = true
				return true
			}

			return false
		})

		if !doneOnline {
			player, found := models.FindPlayerByName(playerName)

			if !found {
				session.Tell("Something went wrong while granting flag")
				return
			}

			player.SetFlag(flagName, value)
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
