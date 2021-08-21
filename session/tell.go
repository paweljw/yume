package session

import (
	"yume/color"
	"strings"
	"fmt"
)

func TellEveryone(s string, a ...interface{}) {
	for e := Sessions.Front(); e != nil; e = e.Next() {
		tellSession(*e.Value.(*Session), s, a...)
	}
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
