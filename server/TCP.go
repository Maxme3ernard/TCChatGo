package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(c net.Conn) {
	text, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Println(text)
	c.Write([]byte(text))
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
}
