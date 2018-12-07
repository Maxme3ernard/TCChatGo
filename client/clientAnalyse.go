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

//Fonction appelée quand l'utilisateur est connecté au chat, qui lui souhaite la bienvenue.
func welcome(serverName string) {
	fmt.Println("Bienvenue sur " + serverName + ".")
}

//Fonction appelée quand un utilisateur rejoint le chat, affiche l'information.
func userin(userName string) {
	fmt.Println(userName + " a rejoint le chat. Bienvenue à lui!")
}

//Fonction appelée quand un utilisateur envoie un message, affiche le message.
func messageReceived(userName string, msg string) {
	fmt.Println(userName + " dit: " + msg)
}

//Fonction appelée quand un utilisateur se déconnecte, affiche l'information.
func userout(userName string) {
	fmt.Println(userName + " nous a quitté.")
}

//Fonction eprmettant de décoder le string reçu en TCP et lancer les instructions correspondantes.
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

//Fonction principale du client, qui va tourner en boucle jusqu'à ce que l'utilisateur se déconnecte.
func clientRun(conn net.Conn) {
	text := ""
	stop := false
	reader := bufio.NewReader(conn) //Création du reader chargé de lire les string échangés à travers la connexion
	for !stop {
		//Lecture de l'entrée par l'utilisateur
		if text != "" {
			analyseText
			text = ""
		}
		status, err := reader.ReadString('\n')
		check(err)
		analyseMessage(status)
	}

}

func main() {
	conn, err := net.Dial("tcp", "10.3.141.1:8080") //Etablissement de la connexion
	if err != nil {
		// handle error
	}
	clientRun(conn)
}
