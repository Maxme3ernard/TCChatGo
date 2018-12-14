package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var connexions = make(map[net.Conn]string)
var NomServeur = "Chat TC du groupe 3"
var mux = sync.Mutex{}

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
		_, err := conn.Write([]byte("TCCHAT_WELCOME" + "\t" + NomServeur + "\n"))
		check(err)
		for clients := range connexions {
			_, err := clients.Write([]byte("TCCHAT_USERIN" + "\t" + connexions[conn] + "\n"))
			check(err)
		}
		fmt.Println(connexions[conn] + " a rejoint le chat!")
	}
}

//Fonction appelée lorsqu'un client envoie un message sur le chat.
func envoyerMessage(MessagePayload string, conn net.Conn) {
	MessagePayload = strings.TrimRight(MessagePayload, "\r\n")
	for clients := range connexions {
		_, err := clients.Write([]byte("TCCHAT_BCAST\t" + connexions[conn] + "\t" + MessagePayload + "\n"))
		check(err)
	}
	fmt.Println(connexions[conn] + " dit: " + MessagePayload)
}

//Fonction appelée quand un client se déconnecte du chat.
func deconnecter(conn net.Conn) {
	for clients := range connexions {
		_, err := clients.Write([]byte("TCCHAT_USEROUT" + "\t" + connexions[conn] + "\n"))
		check(err)
	}
	fmt.Println(connexions[conn] + " s'est déconnecté")
	delete(connexions, conn)
}

//Fonction appelée pour décoder le string reçu par le serveur.
func analyseMessage(text string, conn net.Conn) {
	text = strings.TrimRight(text, "\r\n")
	message := strings.Split(text, "\t")
	switch message[0] {
	case "TCCHAT_REGISTER":
		register(message[1], conn)
	case "TCCHAT_MESSAGE":
		if len(message) > 2 {
			for i := 2; i < len(message); i++ {
				message[1] += "\t" + message[i]
			}
			message[1] = strings.TrimRight(message[1], "\t")
		}
		envoyerMessage(message[1], conn)
	case "TCCHAT_DISCONNECT":
		deconnecter(conn)
	}
}

//Fonction appelée lorsqu'un nouveau client cherche à se connecter.
func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			err = conn.Close()
			check(err)
			break
		}
		if message != "" {
			mux.Lock()
			analyseMessage(message, conn)
			mux.Unlock()
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080") //Le serveur écoute en permanence sur le port 8080.
	fmt.Println(ln)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept() //Lorsqu'une nouvelle connexion est demandée, le serveur l'accepte...
		if err != nil {
			panic(err)
		}
		go handleConnection(conn) //... et lance un thread qui va s'occuper de cette connexion, tout en écoutant le port 8080 pour d'autres connexions.
	}
}
