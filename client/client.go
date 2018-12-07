package client

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
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
	var msg = "bonjour"

	conn.Write([]byte(msg))
}
