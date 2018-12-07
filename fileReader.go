package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	//todo
}

func analyseMessage(text string) {
	//text="TCCHAT_REGISTER\tGautier\nTCCHAT_MESSAGE\ttrace ta route mamen\nTCCHAT_REGISTER\tRobin\nTCCHAT_REGISTER\tMaxime\n"
	line := strings.Split(text, "\n")
	for i := 0; i < len(line); i++ {
		message := strings.Split(line[i], "\t")
		switch message[0] {
		case "TCCHAT_REGISTER":
			fmt.Println(message[1] + " a rejoint le chat!")

		}
	}
}

func main() {
	f, err := os.Open("bonjour.txt")
	check(err)
	text, _ := bufio.NewReader(f).ReadString('\n')
	fmt.Println(text)
	fmt.Println("TCCHAT_REGISTER\tgautier\nTCCHAT_MESSAGE\ttrace ta route mamen\nTCCHAT_REGISTER\tRobin\nTCCHAT_REGISTER\tMaxime\n")
	//analyseMessage(text)

}
