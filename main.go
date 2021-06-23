package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const MIN = 1
const MAX = 100

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func handleConnection(c net.Conn) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	c.Write([]byte("Hello, client!\n"))
	c.Write([]byte("#$#mcp version: 2.1 to: 2.1\n"))

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)

		result := strconv.Itoa(random()) + "\n"
		c.Write([]byte(string(result)))
	}

	log.Printf("Finished serving %s\n", c.RemoteAddr().String())
	c.Close()
}

func main() {
	PORT := ":19000"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go handleConnection(c)
	}
}
