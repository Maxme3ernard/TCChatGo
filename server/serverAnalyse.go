package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var connexions = make(map[net.Conn]string)
var NOM_SERVEUR = "Chat TC du groupe 3"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Fonction appelée lors de l'arrivée d'un client sur le chat.
func register(nom string, conn net.Conn) {
	_, ok := connexions[conn]
	if ok {
		deconnecter(conn)
	} else {
		connexions[conn] = nom
		conn.Write([]byte("TCCHAT_WELCOME" + "\t" + NOM_SERVEUR + "\n"))
		for clients := range connexions {
			clients.Write([]byte("TCCHAT_USERIN" + "\t" + connexions[conn] + "\n"))
		}
		//fmt.Println(connexions[conn] + " a rejoint le chat!")
	}
}

//Fonction appelée lorsqu'un client envoie un message sur le chat.
func envoyerMessage(Message_Payload string, conn net.Conn) {
	for clients := range connexions {
		clients.Write([]byte("TCCHAT_BCAST\t" + connexions[conn] + "\t" + Message_Payload + "\n"))
	}
	fmt.Println(connexions[conn] + " dit: " + Message_Payload)
}

//Fonction appelée quand un client se déconnecte du chat.
func deconnecter(conn net.Conn) {
	for clients := range connexions {
		clients.Write([]byte("TCCHAT_USEROUT" + "\t" + connexions[conn] + "\n"))
	}
	delete(connexions, conn)
	conn.Close()
	//fmt.Println(connexions[conn] + " s'est déconnecté")
}

//Fonction appelée pour décoder le string reçu par le serveur.
func analyseMessage(text string, conn net.Conn) {
	message := strings.Split(text, "\t")
	message[1] = strings.TrimRight(message[1], "\r\n")
	switch message[0] {
	case "TCCHAT_REGISTER":
		register(message[1], conn)
	case "TCCHAT_MESSAGE":
		if len(message) > 2 {
			for i := 2; i < len(message); i++ {
				message[1] += "\t" + message[i]
			}
			message[1] = strings.TrimRight(message[1], "\r\t")
		}

		envoyerMessage(message[1])
	case "TCCHAT_DISCONNECT":
		deconnecter(conn)
	}
}

//Fonction appelée lorsqu'un nouveau client cherche à se connecter.
func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for k := 0; k < 10; k++ {
		message, err := reader.ReadString('\n')
		check(err)
		analyseMessage(message, conn)
	}
	//conn.Write([]byte(message))
}

func main() {
	ln, err := net.Listen("tcp", ":8080") //Le serveur écoute en permanence sur le port 8080.
	fmt.Println(ln)
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept() //Lorsqu'une nouvelle connexion est demandée, le serveur l'accepte...
		if err != nil {
			// handle error
		}
		go handleConnection(conn) //... et lance un thread qui va s'occuper de cette connexion, tout en écoutant le port 8080 pour d'autres connexions.
	}

	/*f, err := os.Open("bonjour.txt")
	check(err)*/
}
