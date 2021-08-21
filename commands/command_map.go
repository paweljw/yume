package commands

import (
	"yume/session"
)

func handleQuit(conn *session.Session, command string) {
	conn.Tell("So be it.")
	conn.Finishing = true
}

var Map = map[string]func(conn *session.Session, command string) {
	"quit": handleQuit,
	"set_flag": handleSetFlag,
	"unset_flag": handleUnsetFlag,
}
