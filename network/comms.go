package network

import (
	"yume/color"
	"strings"
	"fmt"
)

func TellEveryone(s string, a ...interface{}) {
	for e := Connections.Front(); e != nil; e = e.Next() {
		tellConnection(e.Value.(Connection), s, a...)
	}
}

func TellPlayer(c Connection, s string, a ...interface{}) {
	tellConnection(c, s, a...)
}

func tellConnection(c Connection, s string, a ...interface{}) {
	var str = fmt.Sprintf(s, a...)
	if !strings.HasSuffix(str, "\n") {
		str = str + "\n"
	}
	str = color.ColorByTags(str)

	c.Connection.Write([]byte(str))
}
