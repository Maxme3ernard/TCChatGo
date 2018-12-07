package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func analyseMessage(text string) {
	message := strings.Split(text, "\t")
	message[1] = strings.TrimRight(message[1], "\r\n")
	switch message[0] {
	case "TCCHAT_REGISTER":
		fmt.Println(message[1] + " a rejoint le chat!")
	}
}

func main() {
	f, err := os.Open("bonjour.txt")
	check(err)
	text, err := bufio.NewReader(f).ReadString('\n')
	check(err)
	fmt.Println(text)
	//fmt.Println("TCCHAT_REGISTER\tgautier\nTCCHAT_MESSAGE\ttrace ta route mamen\nTCCHAT_REGISTER\tRobin\nTCCHAT_REGISTER\tMaxime\n")
	analyseMessage(text)

}
