package network

import (
	"strings"
	"yume/game"
	"log"
)

func toggleFlag(conn *Connection, command string, value bool) {
	if conn.Player.HasFlag("admin") { // TODO: something like a sudo mode, once buffs are implemented
		tokens := strings.Split(command, " ")

		playerName := tokens[1]
		flagName := tokens[2]

		// Step 1: if player is signed on, let's find them and do them up
		doneOnline := false

		for e := Connections.Front(); e != nil; e = e.Next() {
			c := *e.Value.(*Connection)

			if c.Player.NameEquals(playerName) {
				c.Player.SetFlag(flagName, value)
				c.Player.SaveToFile()
				c.tell("You were granted flag: %s", flagName)
				conn.tell("Granted flag %s to %s (online)", flagName, playerName)
				log.Printf("%s granted flag %s to %s", conn.Player.Name, flagName, playerName)
				doneOnline = true
				break
			}
		}

		if !doneOnline {
			player, err := game.LoadPlayerFromFile(playerName)

			if err != nil {
				conn.tell("Something went wrong while granting flag")
				return
			}

			player.SetFlag(flagName, value)
			player.SaveToFile()
			conn.tell("Granted flag %s to %s (offline)", flagName, playerName)
			log.Printf("%s granted flag %s to %s", conn.Player.Name, flagName, playerName)
		}
	} else {
		conn.tell("You're not allowed to do that.")
	}
}

func handleSetFlag(conn *Connection, command string) {
	toggleFlag(conn, command, true)
}

func handleUnsetFlag(conn *Connection, command string) {
	toggleFlag(conn, command, false)
}
