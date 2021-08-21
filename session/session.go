package session

import (
	"net"
	"log"
	"bufio"
	"strings"
	"container/list"
	"yume/game"
)

// TODO: Lists in go are not thread-safe, oh joy.
var Sessions = list.New()

type Session struct {
	Connection net.Conn
	State      SessionState
	Player	   *game.Player
	Finishing  bool
}

func (session *Session) Chomp() (string, error) {
	netData, err := bufio.NewReader(session.Connection).ReadString('\n')
	if err != nil {
		log.Println(err)
		session.Tell("Something went very wrong. Please sign in again.")
		return "", err
	}
	session.Tell("\n") // Send a single newline as a keepalive

	command := strings.TrimSpace(netData)

	return command, nil
}

func (session *Session) Tell(msg string, a...interface{}) {
	TellPlayer(*session, msg, a...)
}

func RemoveSessionFromList(session *Session) {
	for e := Sessions.Front(); e != nil; e = e.Next() {
		if e.Value == session {
			Sessions.Remove(e)
			return
		}
	}
}

