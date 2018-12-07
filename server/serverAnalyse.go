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

func register(nom string) {
	fmt.Println(nom + " a rejoint le chat!")
}

func envoyerMessage(Message_Payload string) {
	fmt.Println("...USER... dit: " + Message_Payload)
}

func deconnecter() {
	fmt.Println("...USER... s'est déconnecté")
}

func analyseMessage(text string) {
	message := strings.Split(text, "\t")
	message[1] = strings.TrimRight(message[1], "\r\n")
	switch message[0] {
	case "TCCHAT_REGISTER":
		register(message[1])
	case "TCCHAT_MESSAGE":
		if len(message) > 2 {
			for i := 2; i < len(message); i++ {
				message[1] += "\t" + message[i]
			}
			message[1] = strings.TrimRight(message[1], "\r\t")
		}

		envoyerMessage(message[1])
	case "TCCHAT_DISCONNECT":
		deconnecter()
	}
}

func handleConnection(c net.Conn) {
	reader := bufio.NewReader(c)
	for k := 0; k < 10; k++ {
		message, err := reader.ReadString('\n')
		check(err)
		analyseMessage(message)
	}
	//c.Write([]byte(conn))
}

func main() {
	ln, err := net.Listen("tcp", ":12345")
	fmt.Println(ln)
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}

	/*f, err := os.Open("bonjour.txt")
	check(err)*/
}
