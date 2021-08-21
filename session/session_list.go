package session

import (
	"container/list"
	"yume/commons"
)

type SessionList struct {
	commons.SafeList
}

func (list *SessionList) ConcurrentPushBack(value interface{}) {
	list.Mutex.Lock()
	list.List.PushBack(value)
	list.Mutex.Unlock()
}

type iterationCallback func(*list.Element) bool

func (list *SessionList) ConcurrentIterate(callback iterationCallback) {
	list.Mutex.Lock()

	for e:= list.List.Front(); e != nil; e = e.Next() {
		if callback(e) {
			break
		}
	}

	list.Mutex.Unlock()

}

func CreateSessionList() SessionList {
	return SessionList{SafeList: commons.SafeList{List: list.New()}}
}
