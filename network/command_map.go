package network

func handleQuit(conn *Connection, command string) {
	conn.tell("So be it.")
	conn.Finishing = true
}

var commandMap = map[string]func(conn *Connection, command string) {
	"quit": handleQuit,
	"set_flag": handleSetFlag,
	"unset_flag": handleUnsetFlag,
}
