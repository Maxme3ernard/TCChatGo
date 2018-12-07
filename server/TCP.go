package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(c net.Conn) {
	conn, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Println(conn)
	conn = "Yes et toi baby?" + "\n"
	c.Write([]byte(conn))
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
