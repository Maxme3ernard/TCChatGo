package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "10.3.141.1:8080")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "Ca roule ma poule?"+"\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

}
