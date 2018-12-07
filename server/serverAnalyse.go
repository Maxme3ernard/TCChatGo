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

func main() {
	f, err := os.Open("bonjour.txt")
	check(err)
	reader := bufio.NewReader(f)
	for k := 0; k < 10; k++ {
		message, err := reader.ReadString('\n')
		check(err)
		analyseMessage(message)
	}
}
