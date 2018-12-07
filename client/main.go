package main

import (
	"fmt"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func welcome(serverName string) {
	fmt.Println("Bienvenue sur " + serverName)
}

func userin(userName string) {
	fmt.Println(userName + " a rejoint le chat. Bienvenue à lui!")
}

func userout(userName string) {
	fmt.Println(userName + " nous a quitté.")
}

func analyseMessage(text string) {
	message := strings.Split(text, "\t")
	message[1] = strings.TrimRight(message[1], "\r\n")
	switch message[0] {

	}
}
