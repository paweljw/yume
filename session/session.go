package session

import (
	"net"
	"log"
	"bufio"
	"strings"
	"container/list"
	"yume/game"
)

var Sessions = CreateSessionList()

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
	Sessions.ConcurrentIterate(func(e *list.Element) bool {
		if e.Value == session {
			Sessions.List.Remove(e)
			return true
		}

		return false
	})
}
