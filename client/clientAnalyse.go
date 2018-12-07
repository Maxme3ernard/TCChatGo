package main

import (
	"bufio"
	"fmt"
	"net"
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
	conn, err := net.Dial("tcp", "10.3.141.1:8080")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "Ca roule ma poule?"+"\n")
	reader := bufio.NewReader(conn)
	for k := 0; k < 10; k++ {
		status, err := reader.ReadString('\n')
		check(err)
		fmt.Println(status)
	}
}
