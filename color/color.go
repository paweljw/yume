package color

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Purple = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[37m"

func Colorize(s string, c string) string {
	return c + s + Reset
}
