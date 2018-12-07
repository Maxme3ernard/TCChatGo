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
		envoyerMessage(message[1])
	case "TCCHAT_DISCONNECT":
		deconnecter()
	}
}

func main() {
	f, err := os.Open("bonjour.txt")
	check(err)
	text, err := bufio.NewReader(f).ReadString('\n')
	check(err)
	fmt.Println(text)
	//fmt.Println("TCCHAT_REGISTER\tgautier\nTCCHAT_MESSAGE\ttrace ta route mamen\nTCCHAT_REGISTER\tRobin\nTCCHAT_REGISTER\tMaxime\n")
	analyseMessage(text)

}
