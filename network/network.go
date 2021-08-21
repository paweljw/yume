package network

import (
	"log"
	"strings"
	cfg "yume/config"
	"yume/game"
	ses "yume/session"
	"yume/commands"
)


func HandleSession(session *ses.Session) {
	log.Printf("Serving %s\n", session.Connection.RemoteAddr().String())
	session.Player = game.DefaultPlayer()
	session.Finishing = false

	session.Tell(cfg.GetMessage("motd"))
	session.State = ses.NewSession

	for {
		if session.State == ses.Playing {
			// TODO: Proper command tokenizer
			command, _ := session.Chomp()

			actionWord := strings.ToLower(strings.Split(command, " ")[0])

			handler := commands.Map[actionWord]

			if handler != nil {
				handler(session, command)
			} else {
				session.Tell("What?")
			}

		} else { // this covers non playing states
			handler := commands.NonPlayingStates[session.State]

			if handler != nil {
				handler(session)
				continue
			} else {
				session.Tell("How'd you even get here, bro? Bye.")
				break
			}
		}

		if session.Finishing {
			break
		}
	}

	if session.Player.IsSaveable() {
		session.Tell("Saving character...")
		session.Player.SaveToFile()
		log.Printf("Saved %s to disk", session.Player.Name)
	}

	log.Printf("Finished serving %s (%s)\n", session.Connection.RemoteAddr().String(), session.Player.Name)
	session.Connection.Close()
	ses.RemoveSessionFromList(session)
}

