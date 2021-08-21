package commands

import (
	"yume/session"
)

func handleQuit(conn *session.Session, command string) {
	conn.Tell("So be it.")
	conn.Finishing = true
}

// TODO: some help
var Map = map[string]func(conn *session.Session, command string) {
	"quit": handleQuit,
	"set_flag": handleSetFlag,
	"unset_flag": handleUnsetFlag,
	"go": handleMovement,
	"walk": handleMovement,
	"w": handleMovement,
	"west": handleMovement,
	"e": handleMovement,
	"east": handleMovement,
	"n": handleMovement,
	"north": handleMovement,
	"s": handleMovement,
	"south": handleMovement,
	"ne": handleMovement,
	"northeast": handleMovement,
	"se": handleMovement,
	"southeast": handleMovement,
	"sw": handleMovement,
	"southwest": handleMovement,
	"nw": handleMovement,
	"northwest": handleMovement,
	"u": handleMovement,
	"up": handleMovement,
	"d": handleMovement,
	"down": handleMovement,
	"look": handleLook,
}
