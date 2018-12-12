package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func analyseText(text string, conn net.Conn) {
	if strings.TrimRight(text, "\n") == "!stop" {
		_, err := fmt.Fprintf(conn, "TCCHAT_DISCONNECT"+"\n")
		check(err)
	} else {
		_, err := fmt.Fprintf(conn, "TCCHAT_MESSAGE\t"+text+"\n")
		check(err)

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
	fmt.Println(userName + ": " + msg)
}

//Fonction appelée quand un utilisateur se déconnecte, affiche l'information.
func userout(userName string) {
	fmt.Println(userName + " nous a quitté.")
}

//Fonction eprmettant de décoder le string reçu en TCP et lancer les instructions correspondantes.
func analyseMessage(text string) {
	text = strings.TrimRight(text, "\r\n")
	message := strings.Split(text, "\t")
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
			message[2] = strings.TrimRight(message[2], "\r\t")
		}
		messageReceived(message[1], message[2])
	}
}

//Fonction principale du client, qui va tourner en boucle jusqu'à ce que l'utilisateur se déconnecte.
func clientRun(conn net.Conn) {
	text := ""
	stop := false
	readerConsole := bufio.NewReader(os.Stdin)
	fmt.Println("Quel est votre pseudo?")
	name, _ := readerConsole.ReadString('\n')
	_, err := fmt.Fprintf(conn, "TCCHAT_REGISTER\t"+name+"\n")
	check(err)
	for !stop {
		text, _ = readerConsole.ReadString('\n')
		print(text)

		if text != "" {
			analyseText(text+"\n", conn)
			text = ""
		}
	}

}

func handleConnection(conn net.Conn) {
	readerConnexion := bufio.NewReader(conn) //Création du reader chargé de lire les string échangés à travers la connexion
	for {
		status, err := readerConnexion.ReadString('\n')
		check(err)
		analyseMessage(status)
	}
}

func setupCon() {
	conn, err := net.Dial("tcp", "localhost:8080") //Etablissement de la connexion
	if err != nil {
		// handle error
	}
	go handleConnection(conn)
	clientRun(conn)
}

func main() {
	setupCon()
}
