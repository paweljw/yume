package network

import (
	"log"
	"strings"
	"yume/commands"
	cfg "yume/config"
	"yume/models"
	ses "yume/session"
)

func HandleSession(session *ses.Session) {
	log.Printf("Serving %s\n", session.Connection.RemoteAddr().String())
	session.Player = &models.Player{}
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

				// EEUGH. Still needs a rework of the ordering here.
				if session.Finishing {
					break
				}

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

	if session.Player.ID != 0 {
		session.Tell("Saving character...")
		models.Db.Save(session.Player)
		log.Printf("Saved %s to db", session.Player.Name)
	}

	log.Printf("Finished serving %s (%s)\n", session.Connection.RemoteAddr().String(), session.Player.Name)
	session.Connection.Close()
	ses.RemoveSessionFromList(session)
}
