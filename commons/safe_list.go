package commons

import (
	"sync"
	"container/list"
)

type SafeList struct {
	Mutex	sync.Mutex
	List 	*list.List
}
