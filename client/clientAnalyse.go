package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func welcome(serverName string) {
	fmt.Println("Bienvenue sur " + serverName + ".")
}

func userin(userName string) {
	fmt.Println(userName + " a rejoint le chat. Bienvenue à lui!")
}

func messageReceived(userName string, msg string) {
	fmt.Println(userName + " dit: " + msg)
}

func userout(userName string) {
	fmt.Println(userName + " nous a quitté.")
}

func analyseMessage(text string) {
	message := strings.Split(text, "\t")
	message[1] = strings.TrimRight(message[1], "\r\n")
	switch message[0] {
	case "TCCHAT_WELCOME":
		welcome(message[1])
	case "TCCHAT_USERIN":
		userin(message[1])
	case "TCCHAT_USEROUT":
		userout(message[1])
	case "TCCHAT_BCAST":
		if len(message) > 3 {
			for i := 3; i < len(message); i++ {
				message[2] += "\t" + message[i]
			}
			message[2] = strings.TrimRight(message[1], "\r\t")
		}
		messageReceived(message[1], message[2])
	}
}

func main() {
	f, err := os.Open("log.txt")
	check(err)
	reader := bufio.NewReader(f)
	for k := 0; k < 10; k++ {
		message, err := reader.ReadString('\n')
		check(err)
		analyseMessage(message)
	}
}
