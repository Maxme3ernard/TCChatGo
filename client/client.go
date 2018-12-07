package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":12345")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "Connection baby"+"\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

}
