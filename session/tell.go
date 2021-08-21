package session

import (
	"yume/color"
	"strings"
	"fmt"
	"container/list"
)

func TellEveryone(s string, a ...interface{}) {
	Sessions.ConcurrentIterate(func (e *list.Element) bool {
		tellSession(*e.Value.(*Session), s, a...)
		return false
	})
}

func TellPlayer(c Session, s string, a ...interface{}) {
	tellSession(c, s, a...)
}

func tellSession(c Session, s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	if !strings.HasSuffix(str, "\n") {
		str = str + "\n"
	}
	str = color.ColorByTags(str)

	c.Connection.Write([]byte(str))
}
