package main

import (
	"fmt"
	"strings"

	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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
	dat, err := ioutil.ReadFile("bonjour.txt")
	check(err)
	fmt.Print(string(dat))
	var text = string(dat)
	analyseMessage(text)

}
