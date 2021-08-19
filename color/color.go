package color

import (
	"strings"
)

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Purple = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[37m"

const Bold = "\u001b[1m"
const Underline = "\u001b[4m"
const Reversed = "\u001b[7m"

var ColorMap = map[string]string {
	"{crst}": Reset,
	"{cred}": Red,
	"{cgreen}": Green,
	"{cyellow}": Yellow,
	"{cblue}": Blue,
	"{cpurple}": Purple,
	"{ccyan}": Cyan,
	"{cwhite}": White,
	"{frst}": Reset,
	"{fbold}": Bold,
	"{funderline}": Underline,
	"{freversed}": Reversed,
}

func Colorize(s string, c string) string {
	return c + s + Reset
}

func ColorByTags(s string) string {
	var ret = s

	for tag, color := range ColorMap {
		ret = strings.ReplaceAll(ret, tag, color)
	}

	return ret
}
