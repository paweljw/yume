package session

import (
	"yume/commons"
)

type SessionList struct {
	commons.SafeList
}

func CreateSessionList() SessionList {
	return SessionList{SafeList: commons.NewSafeList()}
}
