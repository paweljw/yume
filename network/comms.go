package network

import (
	"yume/color"
)

func TellEveryone(s string) {
	for e := Connections.Front(); e != nil; e = e.Next() {
		e.Value.(Connection).Connection.Write([]byte(color.Colorize(s, color.Cyan)))
	}
}
