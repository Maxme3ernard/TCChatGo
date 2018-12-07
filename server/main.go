package main

import (
	"fmt"
	"strings"
)

var text = "TCCHAT_REGISTER\tBasile\n"
var clients = make(map[string]Client)

// Definition du client
type Client struct {
	nickname string
}

// Fonction permettant de connecter un client au chat.
func createClient(nick string) {
	clients[nick] = Client{nick}
}

// Fonction permettant de déconnecter le client du chat.
func (c Client) disconnect() {
	delete(clients, c.nickname)
}

func analysemessage(text string) {
	line := strings.Split(text, "\n")
	for i := 0; i < len(line); i++ {
		message := strings.Split(line[i], "\t")
		switch message[0] {
		case "TCChat_REGISTER":

		}
	}
}

func main() {

	message := strings.Split(text, "\t")
	client1 := Client{nickname: message[1]}
	fmt.Println(client1)

	if message[0] == "TCCHAT_REGISTER" {
	}
}
