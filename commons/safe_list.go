package commons

import (
	"sync"
	"container/list"
)

type SafeList struct {
	Mutex	sync.Mutex
	List 	*list.List
}

func NewSafeList() SafeList {
	return SafeList{List: list.New()}
}

func (list *SafeList) ConcurrentPushBack(value interface{}) {
	list.Mutex.Lock()
	list.List.PushBack(value)
	list.Mutex.Unlock()
}

type iterationCallback func(*list.Element) bool

func (list *SafeList) ConcurrentIterate(callback iterationCallback) {
	list.Mutex.Lock()

	for e:= list.List.Front(); e != nil; e = e.Next() {
		if callback(e) {
			break
		}
	}

	list.Mutex.Unlock()

}
